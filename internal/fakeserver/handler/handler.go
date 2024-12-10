// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

package handler

import (
	"github.com/superproj/onex/internal/fakeserver/biz"
	v1 "github.com/superproj/onex/pkg/api/fakeserver/v1"
)

// FakeServerHandler implements the gRPC server for fake server operations.
type FakeServerHandler struct {
	v1.UnimplementedFakeServerServer

	biz biz.IBiz // Business logic interface for handling operations
}

// NewFakeServerHandler creates a new instance of FakeServerHandler.
func NewFakeServerHandler(biz biz.IBiz) *FakeServerHandler {
	return &FakeServerHandler{biz: biz}
}
