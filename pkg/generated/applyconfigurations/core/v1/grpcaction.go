// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GRPCActionApplyConfiguration represents a declarative configuration of the GRPCAction type for use
// with apply.
type GRPCActionApplyConfiguration struct {
	Port    *int32  `json:"port,omitempty"`
	Service *string `json:"service,omitempty"`
}

// GRPCActionApplyConfiguration constructs a declarative configuration of the GRPCAction type for use with
// apply.
func GRPCAction() *GRPCActionApplyConfiguration {
	return &GRPCActionApplyConfiguration{}
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *GRPCActionApplyConfiguration) WithPort(value int32) *GRPCActionApplyConfiguration {
	b.Port = &value
	return b
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *GRPCActionApplyConfiguration) WithService(value string) *GRPCActionApplyConfiguration {
	b.Service = &value
	return b
}
