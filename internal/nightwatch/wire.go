// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

//go:build wireinject
// +build wireinject

package nightwatch

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	gwstore "github.com/onexstack/onex/internal/gateway/store"
	"github.com/onexstack/onex/internal/nightwatch/biz"
	"github.com/onexstack/onex/internal/nightwatch/handler"
	"github.com/onexstack/onex/internal/nightwatch/pkg/validation"
	nwstore "github.com/onexstack/onex/internal/nightwatch/store"
	"github.com/onexstack/onex/internal/pkg/client/store"
	ucstore "github.com/onexstack/onex/internal/usercenter/store"
)

func InitializeAggregateStore(*gorm.DB) (store.Interface, error) {
	wire.Build(
		store.ProviderSet,
		gwstore.ProviderSet,
		ucstore.ProviderSet,
	)

	return nil, nil
}

func InitializeHandler(*gorm.DB) *handler.Handler {
	wire.Build(
		validation.New,
		biz.ProviderSet,
		nwstore.ProviderSet,
		handler.NewHandler,
	)

	return nil
}

func InitializeStore(*gorm.DB) (nwstore.IStore, error) {
	wire.Build(nwstore.ProviderSet)

	return nil, nil
}
