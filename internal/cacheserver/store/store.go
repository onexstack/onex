// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package store

//go:generate mockgen -destination mock_store.go -package store github.com/onexstack/onex/internal/cacheserver/store IStore

import (
	"sync"

	"github.com/dgraph-io/ristretto"
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/onexstack/onex/internal/cacheserver/store/secret"
	"github.com/onexstack/onex/pkg/cache"
	ristrettostore "github.com/onexstack/onex/pkg/cache/store/ristretto"
)

// ProviderSet is a Wire provider set that declares dependency injection rules.
// It includes the NewStore constructor function to generate datastore instances.
// wire.Bind is used to bind the IStore interface to the concrete implementation *datastore,
// allowing automatic injection of *datastore instances wherever IStore is required.
var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

var (
	once sync.Once
	// S is a global variable for convenient access to the initialized datastore
	// instance from other packages.
	S *datastore
)

// IStore defines the methods that the Store layer needs to implement.
type IStore interface {
	Secret() *cache.ChainCache[any]
}

// datastore is the concrete implementation of the IStore.
type datastore struct {
	db     *gorm.DB
	local  cache.Cache[any]
	secret *cache.ChainCache[any]
}

// Ensure datastore implements the IStore.
var _ IStore = (*datastore)(nil)

// NewStore initializes a singleton instance of type IStore.
// It ensures that the datastore is only created once using sync.Once.
func NewStore(db *gorm.DB, disable bool) *datastore {
	// Initialize the singleton datastore instance only once.
	once.Do(func() {
		caches := make([]cache.Cache[any], 0)

		// ristretto configuration has been verified in the application, so this is a legal
		// configuration and no error will be returned.
		riscache, _ := ristretto.NewCache(&ristretto.Config{
			NumCounters: 1000,
			MaxCost:     100,
			BufferItems: 64,
		})

		risstore := ristrettostore.NewRistretto(riscache)
		local := cache.New[any](risstore)
		if !disable {
			caches = append(caches, local)
		}

		mysqlStore := secret.New(db)
		caches = append(caches, cache.New[any](mysqlStore))

		S = &datastore{
			db:     db,
			local:  local,
			secret: cache.NewChain[any](caches...),
		}
	})

	return S
}

// Secret returns a ChainCache for managing secrets.
// Chain returns an instance that implements the ChainStore interface.
func (ds *datastore) Secret() *cache.ChainCache[any] {
	return ds.secret
}
