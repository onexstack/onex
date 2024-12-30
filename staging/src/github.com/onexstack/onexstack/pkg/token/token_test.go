// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package token

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

// TestInit 测试 Init 函数
func TestInit(t *testing.T) {
	// 测试默认配置
	assert.Equal(t, "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", config.key)
	assert.Equal(t, "identityKey", config.identityKey)
	assert.Equal(t, 2*time.Hour, config.expiration)

	// 测试自定义配置
	Init("newKey", "newIdentityKey", 3*time.Hour)

	assert.Equal(t, "newKey", config.key)
	assert.Equal(t, "newIdentityKey", config.identityKey)
	assert.Equal(t, 3*time.Hour, config.expiration)

	// 再次调用 Init，确保配置不会被覆盖
	Init("anotherKey", "anotherIdentityKey", 1*time.Hour)

	assert.Equal(t, "newKey", config.key)                 // 仍然是 "newKey"
	assert.Equal(t, "newIdentityKey", config.identityKey) // 仍然是 "newIdentityKey"
	assert.Equal(t, 3*time.Hour, config.expiration)       // 仍然是 3小时
}

// TestSign 测试 Sign 函数
func TestSign(t *testing.T) {
	identityKey := "testUser"
	tokenString, _, err := Sign(identityKey)

	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	// 解析 token
	parsedIdentityKey, err := Parse(tokenString, config.key)
	assert.NoError(t, err)
	assert.Equal(t, identityKey, parsedIdentityKey)
}

// TestParseInvalidToken 测试解析无效的 token
func TestParseInvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"
	identityKey, err := Parse(invalidToken, config.key)

	assert.Error(t, err)
	assert.Empty(t, identityKey)
}

// TestParseRequestWithGin 测试从 Gin 上下文解析 token
func TestParseRequestWithGin(t *testing.T) {
	// 设置 Gin 上下文
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer ")

	// 创建 Gin 上下文
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// 测试解析请求
	identityKey, err := ParseRequest(c)

	assert.Error(t, err)
	assert.Empty(t, identityKey)

	// 测试有效的 token
	validToken, _, _ := Sign("testUser")
	req.Header.Set("Authorization", "Bearer "+validToken)
	c.Request = req

	identityKey, err = ParseRequest(c)
	assert.NoError(t, err)
	assert.Equal(t, "testUser", identityKey)
}

// TestParseRequestWithGRPC 测试从 gRPC 上下文解析 token
func TestParseRequestWithGRPC(t *testing.T) {
	// 创建 gRPC 上下文
	md := metadata.New(map[string]string{"Authorization": "Bearer "})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	// 测试解析请求
	identityKey, err := ParseRequest(ctx)

	assert.Error(t, err)
	assert.Empty(t, identityKey)

	// 测试有效的 token
	validToken, _, _ := Sign("testUser")
	md = metadata.New(map[string]string{"Authorization": "Bearer " + validToken})
	ctx = metadata.NewIncomingContext(context.Background(), md)

	identityKey, err = ParseRequest(ctx)
	assert.NoError(t, err)
	assert.Equal(t, "testUser", identityKey)
}
