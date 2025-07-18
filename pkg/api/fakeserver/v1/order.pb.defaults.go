// This file defines the Protobuf messages and services for managing Orders.
//

// Code generated by protoc-gen-defaults. DO NOT EDIT.

package v1

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	_ *timestamppb.Timestamp
	_ *durationpb.Duration
	_ *wrapperspb.BoolValue
)

func (x *Order) Default() {
}

func (x *CreateOrderRequest) Default() {
}

func (x *CreateOrderResponse) Default() {
}

func (x *UpdateOrderRequest) Default() {
}

func (x *UpdateOrderResponse) Default() {
}

func (x *DeleteOrderRequest) Default() {
}

func (x *DeleteOrderResponse) Default() {
}

func (x *GetOrderRequest) Default() {
}

func (x *GetOrderResponse) Default() {
}

func (x *ListOrderRequest) Default() {
}

func (x *ListOrderResponse) Default() {
}
