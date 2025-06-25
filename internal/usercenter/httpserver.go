// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package usercenter

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/onexstack/onex/internal/pkg/pprof"
	v1 "github.com/onexstack/onex/pkg/api/usercenter/v1"
)

// NewHTTPServer creates and configures a new HTTP server instance.
func (c *ServerConfig) NewHTTPServer() *http.Server {
	opts := []http.ServerOption{
		// http.WithDiscovery(nil),
		// http.WithEndpoint("discovery:///matrix.creation.service.grpc"),
		// Define the middleware chain with variable options.
		http.Middleware(c.middlewares...),
		// Add filter options to the middleware chain.
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization",
				"X-Idempotent-ID",
			}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.cfg.HTTPOptions.Network != "" {
		opts = append(opts, http.Network(c.cfg.HTTPOptions.Network))
	}
	if c.cfg.HTTPOptions.Timeout != 0 {
		opts = append(opts, http.Timeout(c.cfg.HTTPOptions.Timeout))
	}
	if c.cfg.HTTPOptions.Addr != "" {
		opts = append(opts, http.Address(c.cfg.HTTPOptions.Addr))
	}
	if c.cfg.TLSOptions.UseTLS {
		opts = append(opts, http.TLSConfig(c.cfg.TLSOptions.MustTLSConfig()))
	}

	// Create and return the server instance.
	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/openapi/", h)
	srv.Handle("/metrics", promhttp.Handler())
	srv.Handle("", pprof.NewHandler())

	// Register the Gateway HTTP handler.
	v1.RegisterUserCenterHTTPServer(srv, c.handler)

	return srv
}
