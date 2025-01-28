// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package gateway

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

// NewGRPCServer creates and configures a new gRPC server instance.
func (c *ServerConfig) NewGRPCServer() *grpc.Server {
	opts := []grpc.ServerOption{
		// grpc.WithDiscovery(nil),
		// grpc.WithEndpoint("discovery:///matrix.creation.service.grpc"),
		// Define the middleware chain with variable options.
		grpc.Middleware(c.middlewares...),
	}

	if c.cfg.GRPCOptions.Network != "" {
		opts = append(opts, grpc.Network(c.cfg.GRPCOptions.Network))
	}
	if c.cfg.GRPCOptions.Timeout != 0 {
		opts = append(opts, grpc.Timeout(c.cfg.GRPCOptions.Timeout))
	}
	if c.cfg.GRPCOptions.Addr != "" {
		opts = append(opts, grpc.Address(c.cfg.GRPCOptions.Addr))
	}
	if c.cfg.TLSOptions.UseTLS {
		opts = append(opts, grpc.TLSConfig(c.cfg.TLSOptions.MustTLSConfig()))
	}

	// Create a new gRPC server with the configured options.
	srv := grpc.NewServer(opts...)

	// Register the Gateway service handler with the gRPC server.
	v1.RegisterGatewayServer(srv, c.handler)

	return srv
}
