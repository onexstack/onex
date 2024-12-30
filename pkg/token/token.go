// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package token

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Config 包括 token 包的配置选项.
type Config struct {
	// key 用于签发和解析 token 的密钥.
	key string
	// identityKey 是 token 中用户身份的键.
	identityKey string
	// expiration 是签发的 token 过期时间
	expiration time.Duration
}

var (
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", "identityKey", 2 * time.Hour}
	once   sync.Once // 确保配置只被初始化一次
)

// Init 设置包级别的配置 config, config 会用于本包后面的 token 签发和解析.
func Init(key string, identityKey string, expiration time.Duration) {
	once.Do(func() {
		if key != "" {
			config.key = key // 设置密钥
		}
		if identityKey != "" {
			config.identityKey = identityKey // 设置身份键
		}
		if expiration != 0 {
			config.expiration = expiration
		}
	})
}

// Parse 使用指定的密钥 key 解析 token，解析成功返回 token 上下文，否则报错.
func Parse(tokenString string, key string) (string, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保 token 加密算法是预期的加密算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(key), nil // 返回密钥
	})
	// 解析失败
	if err != nil {
		return "", err
	}

	var identityKey string
	// 如果解析成功，从 token 中取出 token 的主题
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if key, exists := claims[config.identityKey]; exists {
			if identity, valid := key.(string); valid {
				identityKey = identity // 获取身份键
			}
		}
	}
	if identityKey == "" {
		return "", jwt.ErrSignatureInvalid
	}

	return identityKey, nil
}

// ParseRequest 从请求头中获取令牌，并将其传递给 Parse 函数以解析令牌.
func ParseRequest(ctx context.Context) (string, error) {
	var (
		token string
		err   error
	)

	switch typed := ctx.(type) {
	// 使用 Gin 框架开发的 HTTP 服务
	case *gin.Context:
		header := typed.Request.Header.Get("Authorization")
		if len(header) == 0 {
			//nolint: err113
			return "", errors.New("the length of the `Authorization` header is zero") // 返回错误
		}

		// 从请求头中取出 token
		_, _ = fmt.Sscanf(header, "Bearer %s", &token) // 解析 Bearer token
	// 使用 google.golang.org/grpc 框架开发的 gRPC 服务
	default:
		token, err = auth.AuthFromMD(typed, "Bearer")
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "invalid auth token")
		}
	}

	return Parse(token, config.key) // 解析 token
}

// Sign 使用 jwtSecret 签发 token，token 的 claims 中会存放传入的 subject.
func Sign(identityKey string) (string, time.Time, error) {
	// 计算过期时间
	expireAt := time.Now().Add(config.expiration)

	// Token 的内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,       // 存放用户身份
		"nbf":              time.Now().Unix(), // token 生效时间
		"iat":              time.Now().Unix(), // token 签发时间
		"exp":              expireAt.Unix(),   // token 过期时间
	})

	// 签发 token
	tokenString, err := token.SignedString([]byte(config.key))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expireAt, nil // 返回 token 字符串、过期时间和错误
}
