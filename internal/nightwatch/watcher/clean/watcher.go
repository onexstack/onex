// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package clean is a watcher implement.
package clean

import (
	"context"

	"github.com/onexstack/onexstack/pkg/log"
	"github.com/onexstack/onexstack/pkg/watch/registry"

	"github.com/onexstack/onex/internal/nightwatch/watcher"
	"github.com/onexstack/onex/internal/pkg/client/store"
)

var _ registry.Watcher = (*Watcher)(nil)

// watcher implement.
type Watcher struct {
	store store.Interface
}

// Run runs the watcher.
func (w *Watcher) Run() {
	_, miners, err := w.store.Gateway().Miner().List(context.Background(), nil)
	if err != nil {
		log.Errorw(err, "Failed to list miners")
		return
	}

	for _, m := range miners {
		log.Infow("Retrieve a miner", "miner", m.Name)
	}
}

// SetAggregateConfig initializes the watcher for later execution.
func (w *Watcher) SetAggregateConfig(config *watcher.AggregateConfig) {
	w.store = config.AggregateStore
}

func init() {
	registry.Register("clean", &Watcher{})
}
