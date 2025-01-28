// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/jinzhu/copier"
	"github.com/onexstack/onexstack/pkg/db"
	"github.com/onexstack/onexstack/pkg/i18n"
	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"
	"github.com/onexstack/onexstack/pkg/store/where"
	"github.com/onexstack/onexstack/pkg/version"
	"go.opentelemetry.io/otel"
	"golang.org/x/text/language"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	"github.com/onexstack/onex/internal/gateway/pkg/locales"
	authzmw "github.com/onexstack/onex/internal/gateway/pkg/middleware/authz"
	"github.com/onexstack/onex/internal/pkg/client/usercenter"
	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/internal/pkg/idempotent"
	"github.com/onexstack/onex/internal/pkg/middleware/authz"
	i18nmw "github.com/onexstack/onex/internal/pkg/middleware/i18n"
	idempotentmw "github.com/onexstack/onex/internal/pkg/middleware/idempotent"
	"github.com/onexstack/onex/internal/pkg/middleware/logging"
	"github.com/onexstack/onex/internal/pkg/middleware/tracing"
	"github.com/onexstack/onex/internal/pkg/middleware/validate"
	"github.com/onexstack/onex/pkg/api/gateway/v1"
	clientset "github.com/onexstack/onex/pkg/generated/clientset/versioned"
	"github.com/onexstack/onex/pkg/generated/informers"
)

var (
	// Name is the name of the compiled software.
	Name = "onex-gateway"

	ID, _ = os.Hostname()

	Version = version.Get().String()
)

// Config contains application-related configurations.
type Config struct {
	GRPCOptions       *genericoptions.GRPCOptions
	HTTPOptions       *genericoptions.HTTPOptions
	TLSOptions        *genericoptions.TLSOptions
	UserCenterOptions *usercenter.UserCenterOptions
	MySQLOptions      *genericoptions.MySQLOptions
	RedisOptions      *genericoptions.RedisOptions
	EtcdOptions       *genericoptions.EtcdOptions
	JaegerOptions     *genericoptions.JaegerOptions
	ConsulOptions     *genericoptions.ConsulOptions

	// the rest config for the onex-apiserver
	Kubeconfig *rest.Config
}

// Server represents the web server.
type Server struct {
	srv server.Server
}

// ServerConfig contains the core dependencies and configurations of the server.
type ServerConfig struct {
	cfg         *Config
	appConfig   server.KratosAppConfig
	handler     v1.GatewayServer
	middlewares []middleware.Middleware
}

// NewServer initializes and returns a new Server instance.
func (cfg *Config) NewServer(ctx context.Context) (*Server, error) {
	where.RegisterTenant("userID", func(ctx context.Context) string {
		return contextx.UserID(ctx)
	})

	if err := cfg.JaegerOptions.SetTracerProvider(); err != nil {
		return nil, err
	}

	// You can use `sigs.k8s.io/controller-runtime/pkg/client`.New to created a client
	// which can support unstructured types also.
	// cl, err := client.New(c.Kubeconfig, client.Options{})
	client, err := clientset.NewForConfig(cfg.Kubeconfig)
	if err != nil {
		log.Errorw(err, "Unable to establish connection with onex-apiserver")
		return nil, err
	}

	var mysqlOptions db.MySQLOptions
	var redisOptions db.RedisOptions
	_ = copier.Copy(&mysqlOptions, cfg.MySQLOptions)
	_ = copier.Copy(&redisOptions, cfg.RedisOptions)

	// Create the core server instance.
	srv, err := InitializeWebServer(ctx.Done(), cfg, client, &mysqlOptions, &redisOptions) //, cfg.UserCenterOptions, cfg.RedisOptions, cfg.EtcdOptions)
	if err != nil {
		return nil, err
	}

	return &Server{srv: srv}, nil
}

// Run starts the server and listens for termination signals.
// It gracefully shuts down the server upon receiving a termination signal.
func (s *Server) Run(ctx context.Context) error {
	return server.Serve(ctx, s.srv)
}

func createInformers(stopCh <-chan struct{}, client clientset.Interface) (informers.SharedInformerFactory, error) {
	f := informers.NewSharedInformerFactory(client, time.Minute)
	msinfor := f.Apps().V1beta1().MinerSets().Informer()
	minfor := f.Apps().V1beta1().Miners().Informer()

	f.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, msinfor.HasSynced, minfor.HasSynced) {
		log.Errorf("Failed to wait for caches to populate")
		return nil, fmt.Errorf("failed to wait caches to populate")
	}

	return f, nil
}

func NewWhiteListMatcher() selector.MatchFunc {
	whitelist := make(map[string]struct{})
	// Placeholder
	// whitelist[v1.Gateway_GetMiner_FullMethodName] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}

func NewMiddlewares(logger krtlog.Logger, idt *idempotent.Idempotent, authz authz.Authorizer, val validate.RequestValidator) []middleware.Middleware {
	meter := otel.Meter("metrics")
	seconds, _ := metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	counter, _ := metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	return []middleware.Middleware{
		recovery.Recovery(
			recovery.WithHandler(func(ctx context.Context, rq, err any) error {
				data, _ := json.Marshal(rq)
				log.W(ctx).Errorw(fmt.Errorf("%v", err), "Catching a panic", "rq", string(data))
				return nil
			}),
		),
		metrics.Server(
			metrics.WithSeconds(seconds),
			metrics.WithRequests(counter),
		),
		i18nmw.Translator(i18n.WithLanguage(language.English), i18n.WithFS(locales.Locales)),
		// circuitbreaker.Client(),
		idempotentmw.Idempotent(idt),
		ratelimit.Server(),
		tracing.Server(),
		selector.Server(authzmw.Authz(authz)).Match(NewWhiteListMatcher()).Build(),
		validate.Validator(val),
		logging.Server(logger),
	}
}

// NewWebServer creates and configures a new core web server.
func NewWebServer(serverConfig *ServerConfig) (server.Server, error) {
	grpcsrv := serverConfig.NewGRPCServer()
	httpsrv := serverConfig.NewHTTPServer()
	return server.NewKratosServer(serverConfig.appConfig, grpcsrv, httpsrv)
}

func ProvideKratosAppConfig(registrar registry.Registrar) server.KratosAppConfig {
	return server.KratosAppConfig{
		ID:        ID,
		Name:      Name,
		Version:   Version,
		Metadata:  map[string]string{},
		Registrar: registrar,
	}
}

func ProvideKratosLogger() krtlog.Logger {
	return server.NewKratosLogger(ID, Name, Version)
}
