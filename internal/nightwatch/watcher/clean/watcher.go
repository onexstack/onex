// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package clean is a watcher implement.
package clean

import (
	"context"

	"github.com/onexstack/onex/internal/nightwatch/watcher"
	"github.com/onexstack/onex/internal/pkg/client/store"
	"github.com/onexstack/onex/pkg/log"
	"github.com/onexstack/onex/pkg/watch"
)

var _ watch.Watcher = (*cleanWatcher)(nil)

// watcher implement.
type cleanWatcher struct {
	store store.Interface
}

// Run runs the watcher.
func (w *cleanWatcher) Run() {
	_, miners, err := w.store.Gateway().Miners().List(context.Background(), "")
	if err != nil {
		log.Errorw(err, "Failed to list miners")
		return
	}

	for _, m := range miners {
		log.Infow("Retrieve a miner", "miner", m.Name)
	}
}

// SetAggregateConfig initializes the watcher for later execution.
func (w *cleanWatcher) SetAggregateConfig(config *watcher.AggregateConfig) {
	w.store = config.Store
}

func init() {
	watch.Register("clean", &cleanWatcher{})
}
