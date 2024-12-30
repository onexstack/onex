// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// CapabilitiesApplyConfiguration represents a declarative configuration of the Capabilities type for use
// with apply.
type CapabilitiesApplyConfiguration struct {
	Add  []v1.Capability `json:"add,omitempty"`
	Drop []v1.Capability `json:"drop,omitempty"`
}

// CapabilitiesApplyConfiguration constructs a declarative configuration of the Capabilities type for use with
// apply.
func Capabilities() *CapabilitiesApplyConfiguration {
	return &CapabilitiesApplyConfiguration{}
}

// WithAdd adds the given value to the Add field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Add field.
func (b *CapabilitiesApplyConfiguration) WithAdd(values ...v1.Capability) *CapabilitiesApplyConfiguration {
	for i := range values {
		b.Add = append(b.Add, values[i])
	}
	return b
}

// WithDrop adds the given value to the Drop field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Drop field.
func (b *CapabilitiesApplyConfiguration) WithDrop(values ...v1.Capability) *CapabilitiesApplyConfiguration {
	for i := range values {
		b.Drop = append(b.Drop, values[i])
	}
	return b
}
