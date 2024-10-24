// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PodDNSConfigOptionApplyConfiguration represents a declarative configuration of the PodDNSConfigOption type for use
// with apply.
type PodDNSConfigOptionApplyConfiguration struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// PodDNSConfigOptionApplyConfiguration constructs a declarative configuration of the PodDNSConfigOption type for use with
// apply.
func PodDNSConfigOption() *PodDNSConfigOptionApplyConfiguration {
	return &PodDNSConfigOptionApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PodDNSConfigOptionApplyConfiguration) WithName(value string) *PodDNSConfigOptionApplyConfiguration {
	b.Name = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *PodDNSConfigOptionApplyConfiguration) WithValue(value string) *PodDNSConfigOptionApplyConfiguration {
	b.Value = &value
	return b
}
