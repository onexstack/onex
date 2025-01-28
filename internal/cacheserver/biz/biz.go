// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/onexstack/onex/internal/cacheserver/biz IBiz

import (
	"github.com/golang/protobuf/ptypes/any"
	"github.com/google/wire"

	namespacedv1 "github.com/onexstack/onex/internal/cacheserver/biz/v1/namespaced"
	secretv1 "github.com/onexstack/onex/internal/cacheserver/biz/v1/secret"
	"github.com/onexstack/onex/internal/cacheserver/store"
	"github.com/onexstack/onex/pkg/cache"
)

// ProviderSet is a Wire provider set used to declare dependency injection rules.
// Includes the NewBiz constructor to create a biz instance.
// wire.Bind binds the IBiz interface to the concrete implementation *biz,
// so places that depend on IBiz will automatically inject a *biz instance.
var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

// IBiz defines the methods that must be implemented by the business layer.
type IBiz interface {
	// NamespacedV1 returns the NamespacedBiz business interface.
	NamespacedV1(namespace string) namespacedv1.NamespacedBiz
	// SecretV1 returns the SecretBiz business interface.
	SecretV1() secretv1.SecretBiz
}

// biz is a concrete implementation of IBiz.
type biz struct {
	cache cache.Cache[*any.Any]
	store store.IStore
}

// Ensure that biz implements the IBiz.
var _ IBiz = (*biz)(nil)

// NewBiz creates an instance of IBiz.
func NewBiz(cache cache.Cache[*any.Any], store store.IStore) *biz {
	return &biz{cache: cache, store: store}
}

// NamespacedV1 returns an instance that implements the NamespacedBiz.
func (b *biz) NamespacedV1(namespace string) namespacedv1.NamespacedBiz {
	return namespacedv1.New(b.cache, namespace)
}

// SecretV1 returns an instance that implements the SecretBiz.
func (b *biz) SecretV1() secretv1.SecretBiz {
	return secretv1.New(b.store.Secret())
}
