// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/onexstack/onex/internal/gateway/locales"
	"github.com/onexstack/onex/internal/pkg/middleware/auth"
	"github.com/onexstack/onex/internal/pkg/contextx"
	jwtutil "github.com/onexstack/onex/internal/pkg/util/jwt"
	"github.com/onexstack/onex/pkg/api/zerrors"
	"github.com/onexstack/onex/pkg/i18n"
	"github.com/onexstack/onex/pkg/log"
)

// Auth is a authentication and authorization middleware.
func Auth(a auth.AuthProvider) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, rq any) (reply any, err error) {
			accessToken := jwtutil.TokenFromServerContext(ctx)
			if tr, ok := transport.FromServerContext(ctx); ok {
				userID, allowed, err := a.Auth(ctx, accessToken, "*", tr.Operation())
				if err != nil {
					log.Errorw(err, "Authorization failure occurs", "operation", tr.Operation())
					return nil, err
				}
				if !allowed {
					return nil, zerrors.ErrorForbidden(i18n.FromContext(ctx).T(locales.NoPermission))
				}
				ctx = contextx.NewUserID(ctx, userID)
				ctx = contextx.NewAccessToken(ctx, accessToken)
				ctx = log.WithContext(ctx, "user.id", userID)
			}

			return handler(ctx, rq)
		}
	}
}
