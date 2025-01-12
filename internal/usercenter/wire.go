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
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/onexstack/onex/internal/pkg/bootstrap"
	"github.com/onexstack/onex/internal/pkg/validation"
	"github.com/onexstack/onex/internal/usercenter/auth"
	"github.com/onexstack/onex/internal/usercenter/biz"
	"github.com/onexstack/onex/internal/usercenter/server"
	"github.com/onexstack/onex/internal/usercenter/service"
	"github.com/onexstack/onex/internal/usercenter/store"
	customvalidation "github.com/onexstack/onex/internal/usercenter/validation"
	"github.com/onexstack/onex/pkg/db"
	genericoptions "github.com/onexstack/onex/pkg/options"
)

// wireApp builds and returns a Kratos app with the given options.
// It uses the Wire library to automatically generate the dependency injection code.
func wireApp(
	bootstrap.AppInfo,
	*server.Config,
	*db.MySQLOptions,
	*genericoptions.JWTOptions,
	*genericoptions.RedisOptions,
	*genericoptions.EtcdOptions,
	*genericoptions.KafkaOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		bootstrap.NewEtcdRegistrar,
		server.ProviderSet,
		store.ProviderSet,
		db.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		auth.ProviderSet,
		store.SetterProviderSet,
		NewAuthenticator,
		validation.ProviderSet,
		customvalidation.ProviderSet,
	)

	return nil, nil, nil
}
