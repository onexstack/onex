// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
)

// GRPCServer 代表一个 GRPC 服务器.
type GRPCServer struct {
	srv *grpc.Server
	lis net.Listener
}

// NewGRPCServer 创建一个新的 GRPC 服务器实例.
func NewGRPCServer(
	grpcOptions *genericoptions.GRPCOptions,
	tlsOptions *genericoptions.TLSOptions,
	serverOptions []grpc.ServerOption,
	registerServer func(grpc.ServiceRegistrar),
) (*GRPCServer, error) {
	lis, err := net.Listen("tcp", grpcOptions.Addr)
	if err != nil {
		log.Errorw(err, "Failed to listen")
		return nil, err
	}

	if tlsOptions != nil && tlsOptions.UseTLS {
		tlsConfig := tlsOptions.MustTLSConfig()
		serverOptions = append(serverOptions, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	grpcsrv := grpc.NewServer(serverOptions...)

	registerServer(grpcsrv)
	registerHealthServer(grpcsrv)
	reflection.Register(grpcsrv)

	return &GRPCServer{
		srv: grpcsrv,
		lis: lis,
	}, nil
}

// RunOrDie 启动 GRPC 服务器并在出错时记录致命错误.
func (s *GRPCServer) RunOrDie() {
	log.Infow("Start to listening the incoming requests", "protocol", "grpc", "addr", s.lis.Addr().String())
	if err := s.srv.Serve(s.lis); err != nil {
		log.Fatalw("Failed to serve grpc server", "err", err)
	}
}

// GracefulStop 优雅地关闭 GRPC 服务器.
func (s *GRPCServer) GracefulStop(ctx context.Context) {
	log.Infow("Gracefully stop grpc server")
	s.srv.GracefulStop()
}

// registerHealthServer 注册健康检查服务.
func registerHealthServer(grpcsrv *grpc.Server) {
	// 创建健康检查服务实例
	healthServer := health.NewServer()

	// 设定服务的健康状态
	healthServer.SetServingStatus("MiniBlog", grpc_health_v1.HealthCheckResponse_SERVING)

	// 注册健康检查服务
	grpc_health_v1.RegisterHealthServer(grpcsrv, healthServer)
}
