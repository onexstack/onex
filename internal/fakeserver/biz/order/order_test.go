// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package order

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/onexstack/onex/internal/fakeserver/model"
	"github.com/onexstack/onex/internal/fakeserver/store"
	"github.com/onexstack/onex/internal/pkg/zid"
	v1 "github.com/onexstack/onex/pkg/api/fakeserver/v1"
)

func fakeOrder(id int64) *model.OrderM {
	return &model.OrderM{
		ID:        id,
		OrderID:   zid.Order.New(uint64(1)),
		Customer:  "colin",
		Product:   "iphone15",
		Quantity:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Test_orderBiz_Create(t *testing.T) {
	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		rq  *v1.CreateOrderRequest
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *v1.CreateOrderResponse
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &orderBiz{
				ds: tt.fields.ds,
			}
			got, err := b.Create(tt.args.ctx, tt.args.rq)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_orderBiz_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 构造期望的返回结果
	fakeOrders := []*model.OrderM{fakeOrder(1), fakeOrder(2), fakeOrder(3)}
	wantOrders := make([]*v1.Order, 0, len(fakeOrders))
	for _, o := range fakeOrders {
		wantOrders = append(wantOrders, &v1.Order{
			OrderID:   o.OrderID,
			Customer:  o.Customer,
			Product:   o.Product,
			Quantity:  o.Quantity,
			CreatedAt: timestamppb.New(o.CreatedAt),
			UpdatedAt: timestamppb.New(o.UpdatedAt),
		})
	}

	mockOrderStore := store.NewMockOrderStore(ctrl)
	mockOrderStore.EXPECT().List(gomock.Any(), gomock.Any()).Return(int64(3), fakeOrders, nil).Times(1)

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Orders().Return(mockOrderStore).Times(1)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		rq  *v1.ListOrderRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.ListOrderResponse
	}{
		{
			name: "default",
			fields: fields{
				ds: mockStore,
			},
			args: args{
				ctx: context.Background(),
				rq:  &v1.ListOrderRequest{},
			},
			want: &v1.ListOrderResponse{TotalCount: 3, Orders: wantOrders},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &orderBiz{
				ds: tt.fields.ds,
			}
			got, err := b.List(tt.args.ctx, tt.args.rq)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
