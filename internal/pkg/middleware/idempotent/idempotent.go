// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package idempotent

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/onexstack/onex/internal/pkg/idempotent"
	"github.com/onexstack/onex/pkg/api/errno"
	v1 "github.com/onexstack/onex/pkg/api/gateway/v1"
)

func idempotentBlacklist() selector.MatchFunc {
	blacklist := make(map[string]struct{})
	blacklist[v1.Gateway_CreateMiner_FullMethodName] = struct{}{}
	blacklist[v1.Gateway_CreateMinerSet_FullMethodName] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := blacklist[operation]; ok {
			return true
		}
		return false
	}
}

func Idempotent(idt *idempotent.Idempotent) middleware.Middleware {
	return selector.Server(
		func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, rq any) (rp any, err error) {
				if tr, ok := transport.FromServerContext(ctx); ok {
					token := tr.RequestHeader().Get("X-Idempotent-ID")
					if token != "" {
						if idt.Check(ctx, token) {
							return handler(ctx, rq)
						}
						return nil, errno.ErrorIdempotentTokenExpired("idempotent token is invalid")
					}
				}

				return nil, errno.ErrorIdempotentMissingToken("idempotent token is missing")
			}
		},
	).Match(idempotentBlacklist()).Build()
}
