// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package server

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	krtlog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	consulapi "github.com/hashicorp/consul/api"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/onexstack/onexstack/pkg/log"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
)

// The purpose of defining the AppConfig is to demonstrate the usage of wire.Struct.
type KratosAppConfig struct {
	ID        string
	Name      string
	Version   string
	Metadata  map[string]string
	Registrar registry.Registrar
}

// The purpose of defining the AppConfig is to demonstrate the usage of wire.Struct.
type KratosServer struct {
	kapp *kratos.App
}

func NewKratosServer(cfg KratosAppConfig, servers ...transport.Server) (*KratosServer, error) {
	kapp := kratos.New(
		kratos.ID(cfg.ID+"."+cfg.Name),
		kratos.Name(cfg.Name),
		kratos.Version(cfg.Version),
		kratos.Metadata(cfg.Metadata),
		kratos.Logger(NewKratosLogger(cfg.ID, cfg.Name, cfg.Version)),
		kratos.Registrar(cfg.Registrar),
		kratos.Server(servers...),
	)

	return &KratosServer{kapp: kapp}, nil
}

func (s *KratosServer) RunOrDie() {
	log.Infow("Start to listening the incoming requests", "protocol", "kratos")
	if err := s.kapp.Run(); err != nil {
		log.Fatalw("Failed to serve kratos application", "err", err)
	}
}

func (s *KratosServer) GracefulStop(ctx context.Context) {
	log.Infow("Gracefully stop kratos application")
	if err := s.kapp.Stop(); err != nil {
		log.Errorw(err, "Failed to gracefully shutdown kratos application")
	}
}

func NewKratosLogger(id, name, version string) krtlog.Logger {
	return krtlog.With(log.Default(),
		"ts", krtlog.DefaultTimestamp,
		"caller", krtlog.DefaultCaller,
		"service.id", id,
		"service.name", name,
		"service.version", version,
	)
}

func NewEtcdRegistrar(opts *genericoptions.EtcdOptions) registry.Registrar {
	if opts == nil {
		panic("etcd registrar options must be set.")
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   opts.Endpoints,
		DialTimeout: opts.DialTimeout,
		TLS:         opts.TLSOptions.MustTLSConfig(),
		Username:    opts.Username,
		Password:    opts.Password,
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

func NewConsulRegistrar(opts *genericoptions.ConsulOptions) registry.Registrar {
	if opts == nil {
		panic("consul registrar options must be set.")
	}

	c := consulapi.DefaultConfig()
	c.Address = opts.Addr
	c.Scheme = opts.Scheme
	cli, err := consulapi.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
