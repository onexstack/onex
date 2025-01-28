package usercenter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/golang-jwt/jwt/v4"
	"github.com/onexstack/onexstack/pkg/authn"
	jwtauthn "github.com/onexstack/onexstack/pkg/authn/jwt"
	"github.com/onexstack/onexstack/pkg/authn/jwt/store/redis"
	"github.com/onexstack/onexstack/pkg/core"
	"github.com/onexstack/onexstack/pkg/db"
	"github.com/onexstack/onexstack/pkg/i18n"
	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"
	"github.com/onexstack/onexstack/pkg/store/where"
	"github.com/onexstack/onexstack/pkg/version"
	"go.opentelemetry.io/otel"
	"golang.org/x/text/language"

	"github.com/onexstack/onex/internal/pkg/contextx"
	mwjwt "github.com/onexstack/onex/internal/pkg/middleware/authn/jwt"
	i18nmw "github.com/onexstack/onex/internal/pkg/middleware/i18n"
	"github.com/onexstack/onex/internal/pkg/middleware/logging"
	"github.com/onexstack/onex/internal/pkg/middleware/tracing"
	"github.com/onexstack/onex/internal/pkg/middleware/validate"
	"github.com/onexstack/onex/internal/usercenter/pkg/locales"
	"github.com/onexstack/onex/pkg/api/usercenter/v1"
)

var (
	// Name is the name of the compiled software.
	Name = "onex-usercenter"

	ID, _ = os.Hostname()

	Version = version.Get().String()
)

// Config contains application-related configurations.
type Config struct {
	GRPCOptions   *genericoptions.GRPCOptions
	HTTPOptions   *genericoptions.HTTPOptions
	TLSOptions    *genericoptions.TLSOptions
	JWTOptions    *genericoptions.JWTOptions
	MySQLOptions  *genericoptions.MySQLOptions
	RedisOptions  *genericoptions.RedisOptions
	EtcdOptions   *genericoptions.EtcdOptions
	KafkaOptions  *genericoptions.KafkaOptions
	JaegerOptions *genericoptions.JaegerOptions
	ConsulOptions *genericoptions.ConsulOptions
}

// Server represents the web server.
type Server struct {
	srv server.Server
}

// ServerConfig contains the core dependencies and configurations of the server.
type ServerConfig struct {
	cfg         *Config
	appConfig   server.KratosAppConfig
	handler     v1.UserCenterServer
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

	var mysqlOptions db.MySQLOptions
	_ = core.Copy(&mysqlOptions, cfg.MySQLOptions)

	// Create the core server instance.
	srv, err := InitializeWebServer(ctx.Done(), cfg, &mysqlOptions, cfg.JWTOptions, cfg.RedisOptions, cfg.KafkaOptions)
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

func NewWhiteListMatcher() selector.MatchFunc {
	whitelist := make(map[string]struct{})
	whitelist[v1.OperationUserCenterLogin] = struct{}{}
	whitelist[v1.OperationUserCenterCreateUser] = struct{}{}
	whitelist[v1.OperationUserCenterAuth] = struct{}{}
	whitelist[v1.OperationUserCenterAuthorize] = struct{}{}
	whitelist[v1.OperationUserCenterAuthenticate] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}

func NewMiddlewares(logger krtlog.Logger, authn authn.Authenticator, val validate.RequestValidator) []middleware.Middleware {
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
		// idempotentmw.Idempotent(idt),
		ratelimit.Server(),
		tracing.Server(),
		metadata.Server(),
		selector.Server(mwjwt.Server(authn)).Match(NewWhiteListMatcher()).Build(),
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

// NewAuthenticator creates a new JWT-based Authenticator using the provided JWT and Redis options.
func NewAuthenticator(jwtOpts *genericoptions.JWTOptions, redisOpts *genericoptions.RedisOptions) (authn.Authenticator, error) {
	// Create a list of options for jwtauthn.
	opts := []jwtauthn.Option{
		// Specify the issuer of the token
		jwtauthn.WithIssuer("onex-usercenter"),
		// Specify the default expiration time for the token to be issued
		jwtauthn.WithExpired(jwtOpts.Expired),
		// Specify the key to be used when issuing the token
		jwtauthn.WithSigningKey([]byte(jwtOpts.Key)),
		// WithKeyfunc will be used by the Parse methods as a callback function to supply
		// the key for verification.  The function receives the parsed,
		// but unverified Token.  This allows you to use properties in the
		// Header of the token (such as `kid`) to identify which key to use.
		jwtauthn.WithKeyfunc(func(t *jwt.Token) (any, error) {
			// Verify that the signing method is HMAC.
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwtauthn.ErrTokenInvalid
			}
			// Return the signing key.
			return []byte(jwtOpts.Key), nil
		}),
	}

	// Set the signing method based on the provided option.
	var method jwt.SigningMethod
	switch jwtOpts.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}

	opts = append(opts, jwtauthn.WithSigningMethod(method))

	// Create a Redis store for jwtauthn.
	store := redis.NewStore(&redis.Config{
		Addr:      redisOpts.Addr,
		Username:  redisOpts.Username,
		Password:  redisOpts.Password,
		Database:  redisOpts.Database,
		KeyPrefix: "authn_",
	})

	// Create a new jwtauthn instance using the Redis store and options.
	authn := jwtauthn.New(store, opts...)

	return authn, nil
}
