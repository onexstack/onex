// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ContainerImageApplyConfiguration represents a declarative configuration of the ContainerImage type for use
// with apply.
type ContainerImageApplyConfiguration struct {
	Names     []string `json:"names,omitempty"`
	SizeBytes *int64   `json:"sizeBytes,omitempty"`
}

// ContainerImageApplyConfiguration constructs a declarative configuration of the ContainerImage type for use with
// apply.
func ContainerImage() *ContainerImageApplyConfiguration {
	return &ContainerImageApplyConfiguration{}
}

// WithNames adds the given value to the Names field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Names field.
func (b *ContainerImageApplyConfiguration) WithNames(values ...string) *ContainerImageApplyConfiguration {
	for i := range values {
		b.Names = append(b.Names, values[i])
	}
	return b
}

// WithSizeBytes sets the SizeBytes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SizeBytes field is set to the value of the last call.
func (b *ContainerImageApplyConfiguration) WithSizeBytes(value int64) *ContainerImageApplyConfiguration {
	b.SizeBytes = &value
	return b
}
