// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package known

const (
	AdminUsername = "admin"
	AdminUserID   = "user-admin"
)

const (
	// XTraceID 用来定义上下文中的键，代表请求 ID.
	XTraceID = "x-trace-id"

	// XUserID 用来定义上下文的键，代表请求用户 ID. UserID 整个用户生命周期唯一.
	XUserID = "x-user-id"

	// XUsername 用来定义上下文的键，代表请求用户名.
	XUsername = "x-username"
)

const (
	// MaxErrGroupConcurrency defines the maximum concurrency level
	// for error group operations.
	MaxErrGroupConcurrency = 100
)
