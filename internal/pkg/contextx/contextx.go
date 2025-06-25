// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

//nolint:unused
package contextx

import (
	"context"

	"github.com/golang-jwt/jwt/v4"

	"github.com/onexstack/onex/internal/usercenter/model"
)

// 定义全局上下文中的键.
type (
	transCtx     struct{}
	noTransCtx   struct{}
	transLockCtx struct{}
	userIDCtx    struct{}
	traceIDCtx   struct{}
)

type (
	claimsKey      struct{}
	userKey        struct{}
	userMKey       struct{}
	accessTokenKey struct{}
	traceIDKey     struct{}
)

// WithClaims put claims info into context.
func WithClaims(ctx context.Context, claims *jwt.RegisteredClaims) context.Context {
	return context.WithValue(ctx, claimsKey{}, claims)
}

// Claims extract claims info from context.
func Claims(ctx context.Context) *jwt.RegisteredClaims {
	claims, _ := ctx.Value(claimsKey{}).(*jwt.RegisteredClaims)
	return claims
}

// WithUserID put userID into context.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userKey{}, userID)
}

// UserID extract userID from context.
func UserID(ctx context.Context) string {
	userID, _ := ctx.Value(userKey{}).(string)
	return userID
}

// Namespace is an alias for UserID.
func Namespace(ctx context.Context) string {
	userID, _ := ctx.Value(userKey{}).(string)
	return userID
}

// WithAccessToken put accessToken into context.
func WithAccessToken(ctx context.Context, accessToken string) context.Context {
	return context.WithValue(ctx, accessTokenKey{}, accessToken)
}

// AccessToken extract accessToken from context.
func AccessToken(ctx context.Context) string {
	accessToken, _ := ctx.Value(accessTokenKey{}).(string)
	return accessToken
}

// WithUserM put *UserM into context.
func WithUserM(ctx context.Context, user *model.UserM) context.Context {
	return context.WithValue(ctx, userMKey{}, user)
}

// UserM extract *UserM from extract.
func UserM(ctx context.Context) *model.UserM {
	user, _ := ctx.Value(userMKey{}).(*model.UserM)
	return user
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func TraceID(ctx context.Context) string {
	traceID, _ := ctx.Value(traceIDKey{}).(string)
	return traceID
}
