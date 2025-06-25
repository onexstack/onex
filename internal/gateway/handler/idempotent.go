// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package handler

import (
	"context"

	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

func (s *Handler) GetIdempotentToken(ctx context.Context, rq *v1.IdempotentRequest) (*v1.IdempotentResponse, error) {
	return &v1.IdempotentResponse{Token: s.idt.Token(ctx)}, nil
}
