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
	"github.com/onexstack/onexstack/pkg/i18n"
	"github.com/onexstack/onexstack/pkg/log"

	"github.com/onexstack/onex/internal/gateway/pkg/locales"
	"github.com/onexstack/onex/internal/pkg/contextx"
	"github.com/onexstack/onex/internal/pkg/middleware/authz"
	jwtutil "github.com/onexstack/onex/internal/pkg/util/jwt"
	"github.com/onexstack/onex/pkg/api/errno"
)

// Authz is a authentication and authorization middleware.
func Authz(authz authz.Authorizer) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, rq any) (reply any, err error) {
			accessToken := jwtutil.TokenFromServerContext(ctx)
			if tr, ok := transport.FromServerContext(ctx); ok {
				userID, allowed, err := authz.Authorize(ctx, accessToken, "*", tr.Operation())
				if err != nil {
					log.Errorw(err, "Authorization failure occurs", "operation", tr.Operation())
					return nil, err
				}
				if !allowed {
					return nil, errno.ErrorForbidden(i18n.FromContext(ctx).T(locales.NoPermission))
				}
				ctx = contextx.WithUserID(ctx, userID)
				ctx = contextx.WithAccessToken(ctx, accessToken)
				ctx = contextx.WithUserID(ctx, userID)
			}

			return handler(ctx, rq)
		}
	}
}
