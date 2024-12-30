// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package order

//go:generate mockgen -destination mock_order.go -package order github.com/onexstack/onex/internal/fakeserver/biz/order OrderBiz

import (
	"context"
	"errors"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/jinzhu/copier"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"

	"github.com/onexstack/onex/internal/fakeserver/model"
	"github.com/onexstack/onex/internal/fakeserver/store"
	fsv1 "github.com/onexstack/onex/pkg/api/fakeserver/v1"
	"github.com/onexstack/onex/pkg/log"
	"github.com/onexstack/onex/pkg/store/where"
)

type OrderBiz interface {
	Create(ctx context.Context, rq *fsv1.CreateOrderRequest) (*fsv1.CreateOrderResponse, error)
	Update(ctx context.Context, rq *fsv1.UpdateOrderRequest) (*fsv1.UpdateOrderResponse, error)
	Delete(ctx context.Context, rq *fsv1.DeleteOrderRequest) (*fsv1.DeleteOrderResponse, error)
	Get(ctx context.Context, rq *fsv1.GetOrderRequest) (*fsv1.GetOrderResponse, error)
	List(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error)

	OrderExpansion
}

// OrderExpansion defines additional methods for order operations.
type OrderExpansion interface {
	ListWithWorkerPool(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error)
	ListWithAnts(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error)
}

type orderBiz struct {
	ds store.IStore
}

var _ OrderBiz = (*orderBiz)(nil)

func New(ds store.IStore) *orderBiz {
	return &orderBiz{ds: ds}
}

func (b *orderBiz) Create(ctx context.Context, rq *fsv1.CreateOrderRequest) (*fsv1.CreateOrderResponse, error) {
	var orderM model.OrderM
	_ = copier.Copy(&orderM, rq)

	if err := b.ds.Orders().Create(ctx, &orderM); err != nil {
		return nil, fsv1.ErrorOrderCreateFailed("create order failed: %v", err)
	}

	return &fsv1.CreateOrderResponse{OrderID: orderM.OrderID}, nil
}

func (b *orderBiz) Update(ctx context.Context, rq *fsv1.UpdateOrderRequest) (*fsv1.UpdateOrderResponse, error) {
	orderM, err := b.ds.Orders().Get(ctx, where.T(ctx).F("order_id", rq.OrderID))
	if err != nil {
		return nil, err
	}

	if rq.Customer != nil {
		orderM.Customer = *rq.Customer
	}

	if rq.Product != nil {
		orderM.Product = *rq.Product
	}

	if rq.Quantity != nil {
		orderM.Quantity = *rq.Quantity
	}

	if err := b.ds.Orders().Update(ctx, orderM); err != nil {
		return nil, err
	}

	return &fsv1.UpdateOrderResponse{}, nil
}

// Delete 是 OrderBiz 接口中 `Delete` 方法的实现.
func (b *orderBiz) Delete(ctx context.Context, rq *fsv1.DeleteOrderRequest) (*fsv1.DeleteOrderResponse, error) {
	if err := b.ds.Orders().Delete(ctx, where.T(ctx).F("order_id", rq.OrderID)); err != nil {
		return nil, err
	}

	return &fsv1.DeleteOrderResponse{}, nil
}

func (b *orderBiz) Get(ctx context.Context, rq *fsv1.GetOrderRequest) (*fsv1.GetOrderResponse, error) {
	orderM, err := b.ds.Orders().Get(ctx, where.T(ctx).F("order_id", rq.OrderID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fsv1.ErrorOrderNotFound(err.Error())
		}

		return nil, err
	}

	var order fsv1.Order
	_ = copier.Copy(&order, orderM)
	order.CreatedAt = timestamppb.New(orderM.CreatedAt)
	order.UpdatedAt = timestamppb.New(orderM.UpdatedAt)

	return &fsv1.GetOrderResponse{Order: &order}, nil
}

func (b *orderBiz) List(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error) {
	count, orderList, err := b.ds.Orders().List(ctx, where.T(ctx).P(int(rq.Offset), int(rq.Limit)))
	if err != nil {
		return nil, err
	}

	var m sync.Map
	eg, ctx := errgroup.WithContext(ctx)
	// 使用 goroutine 提高接口性能
	for _, order := range orderList {
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				var o fsv1.Order
				_ = copier.Copy(&o, order)
				m.Store(order.ID, &fsv1.Order{
					OrderID:   order.OrderID,
					Customer:  order.Customer,
					Product:   order.Product,
					Quantity:  order.Quantity,
					CreatedAt: timestamppb.New(order.CreatedAt),
					UpdatedAt: timestamppb.New(order.UpdatedAt),
				})

				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		log.C(ctx).Errorw(err, "Failed to wait all function calls returned")
		return nil, err
	}

	// The following code block is used to maintain the consistency of query order.
	orders := make([]*fsv1.Order, 0, len(orderList))
	for _, item := range orderList {
		order, _ := m.Load(item.ID)
		orders = append(orders, order.(*fsv1.Order))
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(orders))

	return &fsv1.ListOrderResponse{TotalCount: count, Orders: orders}, nil
}

// ListWithWorkerPool retrieves a list of all orders from the database use workerpool package.
// Concurrency limits can effectively protect downstream services and control the resource
// consumption of components.
func (b *orderBiz) ListWithWorkerPool(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error) {
	count, orderList, err := b.ds.Orders().List(ctx, where.T(ctx).P(int(rq.Offset), int(rq.Limit)))
	if err != nil {
		return nil, err
	}

	var m sync.Map
	wp := workerpool.New(100)

	// Use goroutine to improve interface performance
	for _, order := range orderList {
		wp.Submit(func() {
			var o fsv1.Order
			// Here simulates a time-consuming concurrent logic.
			_ = copier.Copy(&o, order)
			m.Store(order.ID, &fsv1.Order{
				OrderID:   order.OrderID,
				Customer:  order.Customer,
				Product:   order.Product,
				Quantity:  order.Quantity,
				CreatedAt: timestamppb.New(order.CreatedAt),
				UpdatedAt: timestamppb.New(order.UpdatedAt),
			})

			return
		})
	}

	wp.StopWait()

	// The following code block is used to maintain the consistency of query order.
	orders := make([]*fsv1.Order, 0, len(orderList))
	for _, item := range orderList {
		order, _ := m.Load(item.ID)
		orders = append(orders, order.(*fsv1.Order))
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(orders))

	return &fsv1.ListOrderResponse{TotalCount: count, Orders: orders}, nil
}

// ListWithAnts retrieves a list of all orders from the database use ants package.
// Concurrency limits can effectively protect downstream services and control the
// resource consumption of components.
func (b *orderBiz) ListWithAnts(ctx context.Context, rq *fsv1.ListOrderRequest) (*fsv1.ListOrderResponse, error) {
	count, orderList, err := b.ds.Orders().List(ctx, where.T(ctx).P(int(rq.Offset), int(rq.Limit)))
	if err != nil {
		return nil, err
	}

	var m sync.Map
	var wg sync.WaitGroup
	pool, _ := ants.NewPool(100)
	defer pool.Release()

	// Use goroutine to improve interface performance
	for _, order := range orderList {
		wg.Add(1)
		_ = pool.Submit(func() {
			defer wg.Done()

			var o fsv1.Order
			// Here simulates a time-consuming concurrent logic.
			_ = copier.Copy(&o, order)
			m.Store(order.ID, &fsv1.Order{
				OrderID:   order.OrderID,
				Customer:  order.Customer,
				Product:   order.Product,
				Quantity:  order.Quantity,
				CreatedAt: timestamppb.New(order.CreatedAt),
				UpdatedAt: timestamppb.New(order.UpdatedAt),
			})

			return
		})
	}

	wg.Wait()

	// The following code block is used to maintain the consistency of query order.
	orders := make([]*fsv1.Order, 0, len(orderList))
	for _, item := range orderList {
		order, _ := m.Load(item.ID)
		orders = append(orders, order.(*fsv1.Order))
	}

	log.C(ctx).Debugw("Get orders from backend storage", "count", len(orders))

	return &fsv1.ListOrderResponse{TotalCount: count, Orders: orders}, nil
}
