// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/onexstack/onex/pkg/api/cacheserver/v1"
)

// Set stores a key-value pair in the cache with an optional expiration time.
func (h *Handler) Set(ctx context.Context, rq *v1.SetRequest) (*emptypb.Empty, error) {
	return h.biz.NamespacedV1(rq.Namespace).Set(ctx, rq.Key, rq.Value, rq.Expire)
}

// Del removes a key from the cache by namespace.
func (h *Handler) Del(ctx context.Context, rq *v1.DelRequest) (*emptypb.Empty, error) {
	return h.biz.NamespacedV1(rq.Namespace).Del(ctx, rq.Key)
}

// Get retrieves a key's value from the cache by namespace.
func (h *Handler) Get(ctx context.Context, rq *v1.GetRequest) (*v1.GetResponse, error) {
	return h.biz.NamespacedV1(rq.Namespace).Get(ctx, rq.Key)
}
