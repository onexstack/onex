// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package fakeserver

import (
	"context"
	"os"

	"github.com/jinzhu/copier"

	"github.com/onexstack/onex/internal/fakeserver/biz"
	"github.com/onexstack/onex/internal/fakeserver/handler"
	"github.com/onexstack/onex/internal/fakeserver/model"
	"github.com/onexstack/onex/internal/fakeserver/store"
	"github.com/onexstack/onex/internal/fakeserver/store/fake"
	"github.com/onexstack/onex/internal/fakeserver/store/mysql"
	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/pkg/db"
	genericoptions "github.com/onexstack/onex/pkg/options"
	"github.com/onexstack/onex/pkg/store/where"
)

var (
	// Name is the name of the compiled software.
	Name = "onex-fakeserver"

	// ID contains the host name and any error encountered during the retrieval.
	ID, _ = os.Hostname()
)

// Config represents the configuration of the service.
type Config struct {
	FakeStore     bool
	GRPCOptions   *genericoptions.GRPCOptions
	HTTPOptions   *genericoptions.HTTPOptions
	TLSOptions    *genericoptions.TLSOptions
	MySQLOptions  *genericoptions.MySQLOptions
	JaegerOptions *genericoptions.JaegerOptions
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() completedConfig {
	return completedConfig{cfg}
}

type completedConfig struct {
	*Config
}

// FakeServer represents the fake server.
type FakeServer struct {
	httpsrv    Server
	grpcServer Server
	config     completedConfig
}

// New returns a new instance of Server from the given config.
func (c completedConfig) New(stopCh <-chan struct{}) (*FakeServer, error) {
	where.RegisterTenant("user_id", func(ctx context.Context) string {
		return contextx.FromUserID(ctx)
	})

	if err := c.JaegerOptions.SetTracerProvider(); err != nil {
		return nil, err
	}

	var ds store.IStore
	if c.FakeStore {
		ds = fake.NewStore(500)
	} else {
		var dbOptions db.MySQLOptions
		_ = copier.Copy(&dbOptions, c.MySQLOptions)

		ins, err := db.NewMySQL(&dbOptions)
		if err != nil {
			return nil, err
		}
		ins.AutoMigrate(&model.OrderM{})
		ds = mysql.NewStore(ins)
	}

	biz := biz.NewBiz(ds)
	h := handler.NewFakeServerHandler(biz)

	grpcServer, err := NewGRPCServer(c.GRPCOptions, c.TLSOptions, h)
	if err != nil {
		return nil, err
	}

	// Need start grpc server first. http server depends on grpc sever.
	go grpcServer.RunOrDie()

	httpsrv, err := NewHTTPServer(c.HTTPOptions, c.TLSOptions, c.GRPCOptions)
	if err != nil {
		return nil, err
	}

	return &FakeServer{grpcServer: grpcServer, httpsrv: httpsrv, config: c}, nil
}

func (s *FakeServer) Run(stopCh <-chan struct{}) error {
	go s.httpsrv.RunOrDie()

	<-stopCh

	// The most gracefully way is to shutdown the dependent service first,
	// and then shutdown the depended service.
	s.httpsrv.GracefulStop()
	s.grpcServer.GracefulStop()

	return nil
}
