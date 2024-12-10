// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

package handler

import (
	"context"

	fsv1 "github.com/superproj/onex/pkg/api/fakeserver/v1"
)

// CreateOrder handles the creation of a new order.
func (h *FakeServerHandler) CreateOrder(ctx context.Context, rq *fsv1.CreateOrderRequest) (*fsv1.CreateOrderResponse, error) {
	return h.biz.Orders().Create(ctx, rq)
}

// UpdateOrder handles the update of an existing order.
func (h *FakeServerHandler) UpdateOrder(ctx context.Context, rq *fsv1.UpdateOrderRequest) (*fsv1.UpdateOrderResponse, error) {
	return h.biz.Orders().Update(ctx, rq)
}

// DeleteOrder handles the deletion of an order.
func (h *FakeServerHandler) DeleteOrder(ctx context.Context, rq *fsv1.DeleteOrderRequest) (*fsv1.DeleteOrderResponse, error) {
	return h.biz.Orders().Delete(ctx, rq)
}

// GetOrder handles fetching the details of a specific order.
func (h *FakeServerHandler) GetOrder(ctx context.Context, rq *fsv1.GetOrderRequest) (*fsv1.GetOrderResponse, error) {
	return h.biz.Orders().Get(ctx, rq)
}

// ListOrder handles fetching a list of orders.
func (h *FakeServerHandler) ListOrder(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error) {
	return h.biz.Orders().List(ctx, rq)
}
