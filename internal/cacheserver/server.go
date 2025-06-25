// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package cacheserver

import (
	"context"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/jinzhu/copier"
	"github.com/onexstack/onexstack/pkg/db"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"
	"k8s.io/apimachinery/pkg/util/wait"

	v1 "github.com/onexstack/onex/pkg/api/cacheserver/v1"
	"github.com/onexstack/onex/pkg/cache"
	redisstore "github.com/onexstack/onex/pkg/cache/store/redis"
)

// Config contains application-related configurations.
type Config struct {
	DisableCache  bool
	GRPCOptions   *genericoptions.GRPCOptions
	TLSOptions    *genericoptions.TLSOptions
	RedisOptions  *genericoptions.RedisOptions
	MySQLOptions  *genericoptions.MySQLOptions
	JaegerOptions *genericoptions.JaegerOptions
}

// Server represents the web server.
type Server struct {
	srv server.Server
}

// ServerConfig contains the core dependencies and configurations of the server.
type ServerConfig struct {
	cfg     *Config
	handler v1.CacheServerServer
}

// NewServer initializes and returns a new Server instance.
func (cfg *Config) NewServer(ctx context.Context) (*Server, error) {
	// Initialize Jaeger tracing.
	if err := cfg.JaegerOptions.SetTracerProvider(); err != nil {
		return nil, err
	}

	rds, err := cfg.RedisOptions.NewClient()
	if err != nil {
		return nil, err
	}

	redisStore := redisstore.NewRedis(rds)
	l2cache := cache.New[*any.Any](redisStore)
	l2mgr := cache.NewL2[*any.Any](l2cache, cache.L2WithDisableCache(cfg.DisableCache))
	l2mgr.Wait(wait.ContextForChannel(ctx.Done()))

	// Copy MySQL options to avoid modifying the original configuration.
	var dbOptions db.MySQLOptions
	_ = copier.Copy(&dbOptions, cfg.MySQLOptions)

	// Create the core server instance.
	srv, err := InitializeWebServer(cfg, &dbOptions, l2mgr, cfg.DisableCache)
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

// NewWebServer creates and configures a new core web server.
func NewWebServer(serverConfig *ServerConfig) (server.Server, error) {
	return serverConfig.NewGRPCServer()
}
