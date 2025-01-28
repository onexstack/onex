// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package cacheserver

import (
	"github.com/onexstack/onexstack/pkg/server"
	"google.golang.org/grpc"

	v1 "github.com/onexstack/onex/pkg/api/cacheserver/v1"
)

func (c *ServerConfig) NewGRPCServer() (server.Server, error) {
	// Initialize the gRPC server.
	return server.NewGRPCServer(
		c.cfg.GRPCOptions,
		c.cfg.TLSOptions,
		[]grpc.ServerOption{},
		func(s grpc.ServiceRegistrar) { v1.RegisterCacheServerServer(s, c.handler) })
}
