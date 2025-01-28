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

// SetSecret stores a secret in the system or updates an existing one.
func (h *Handler) SetSecret(ctx context.Context, rq *v1.SetSecretRequest) (*emptypb.Empty, error) {
	return h.biz.SecretV1().Set(ctx, rq)
}

// DelSecret removes a secret from the system.
func (h *Handler) DelSecret(ctx context.Context, rq *v1.DelSecretRequest) (*emptypb.Empty, error) {
	return h.biz.SecretV1().Del(ctx, rq)
}

// GetSecret retrieves a secret from the system.
func (h *Handler) GetSecret(ctx context.Context, rq *v1.GetSecretRequest) (*v1.GetSecretResponse, error) {
	return h.biz.SecretV1().Get(ctx, rq)
}
