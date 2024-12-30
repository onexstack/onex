// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// ResourceStatusApplyConfiguration represents a declarative configuration of the ResourceStatus type for use
// with apply.
type ResourceStatusApplyConfiguration struct {
	Name      *v1.ResourceName                   `json:"name,omitempty"`
	Resources []ResourceHealthApplyConfiguration `json:"resources,omitempty"`
}

// ResourceStatusApplyConfiguration constructs a declarative configuration of the ResourceStatus type for use with
// apply.
func ResourceStatus() *ResourceStatusApplyConfiguration {
	return &ResourceStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ResourceStatusApplyConfiguration) WithName(value v1.ResourceName) *ResourceStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *ResourceStatusApplyConfiguration) WithResources(values ...*ResourceHealthApplyConfiguration) *ResourceStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResources")
		}
		b.Resources = append(b.Resources, *values[i])
	}
	return b
}
