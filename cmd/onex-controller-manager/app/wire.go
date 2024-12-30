// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

//go:build wireinject
// +build wireinject

package app

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/google/wire"

	"github.com/onexstack/onex/internal/gateway/store"
	"github.com/onexstack/onex/pkg/db"
)

func wireStoreClient(*db.MySQLOptions) (store.IStore, error) {
	wire.Build(
		db.ProviderSet,
		store.ProviderSet,
	)

	return nil, nil
}
