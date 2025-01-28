// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

//go:build wireinject
// +build wireinject

package toyblc

import (
	"github.com/google/wire"
	"github.com/onexstack/onexstack/pkg/server"

	"github.com/onexstack/onex/internal/toyblc/pkg/blc"
	"github.com/onexstack/onex/internal/toyblc/pkg/ws"
)

// InitializeWebServer sets up and initializes the web server with all necessary dependencies.
func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		NewAggregatorServer,
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
		wire.FieldsOf(new(*Config), "Miner", "MinMineInterval", "Accounts", "Peers"),
		ws.ProviderSet,
		blc.NewBlockSet,
		wire.FieldsOf(new(*Config), "Address"),
	)
	return nil, nil
}
