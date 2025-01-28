// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package usercenter

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/google/wire"
	"github.com/onexstack/onexstack/pkg/db"
	genericoptions "github.com/onexstack/onexstack/pkg/options"
	"github.com/onexstack/onexstack/pkg/server"
	genericvalidation "github.com/onexstack/onexstack/pkg/validation"

	"github.com/onexstack/onex/internal/pkg/middleware/validate"
	"github.com/onexstack/onex/internal/usercenter/biz"
	"github.com/onexstack/onex/internal/usercenter/handler"
	"github.com/onexstack/onex/internal/usercenter/pkg/auth"
	"github.com/onexstack/onex/internal/usercenter/pkg/validation"
	"github.com/onexstack/onex/internal/usercenter/store"
)

func InitializeWebServer(
	<-chan struct{},
	*Config,
	*db.MySQLOptions,
	*genericoptions.JWTOptions,
	*genericoptions.RedisOptions,
	*genericoptions.KafkaOptions,
) (server.Server, error) {
	wire.Build(
		wire.NewSet(server.NewEtcdRegistrar, wire.FieldsOf(new(*Config), "EtcdOptions")), // dep by AppConfig
		ProvideKratosAppConfig, // server.KratosAppConfig, dep by NewWebServer
		ProvideKratosLogger,    // dep by NewMiddlewares
		// func NewMiddlewares(logger krtlog.Logger,authn authn.Authenticator, val validate.RequestValidator) []middleware.Middleware {
		NewAuthenticator,
		NewWebServer,
		NewMiddlewares,
		store.SetterProviderSet,
		auth.ProviderSet,
		handler.ProviderSet,
		store.ProviderSet,
		biz.ProviderSet,
		db.ProviderSet,
		wire.NewSet(
			validation.ProviderSet,
			genericvalidation.NewValidator,
			wire.Bind(new(validate.RequestValidator), new(*genericvalidation.Validator)),
		),
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
	)
	return nil, nil
}
