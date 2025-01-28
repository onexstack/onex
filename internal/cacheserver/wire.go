// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

//go:build wireinject
// +build wireinject

package cacheserver

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/golang/protobuf/ptypes/any"
	"github.com/google/wire"
	"github.com/onexstack/onexstack/pkg/db"
	"github.com/onexstack/onexstack/pkg/server"

	"github.com/onexstack/onex/internal/cacheserver/biz"
	"github.com/onexstack/onex/internal/cacheserver/handler"
	"github.com/onexstack/onex/internal/cacheserver/store"
	"github.com/onexstack/onex/pkg/cache"
)

func InitializeWebServer(*Config, *db.MySQLOptions, cache.Cache[*any.Any], bool) (server.Server, error) {
	wire.Build(
		NewWebServer,
		db.ProviderSet,
		store.ProviderSet,
		biz.ProviderSet,
		handler.ProviderSet,
		wire.Struct(new(ServerConfig), "*"),
	)
	return nil, nil
}
