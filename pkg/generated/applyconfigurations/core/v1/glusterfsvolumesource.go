// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GlusterfsVolumeSourceApplyConfiguration represents a declarative configuration of the GlusterfsVolumeSource type for use
// with apply.
type GlusterfsVolumeSourceApplyConfiguration struct {
	EndpointsName *string `json:"endpoints,omitempty"`
	Path          *string `json:"path,omitempty"`
	ReadOnly      *bool   `json:"readOnly,omitempty"`
}

// GlusterfsVolumeSourceApplyConfiguration constructs a declarative configuration of the GlusterfsVolumeSource type for use with
// apply.
func GlusterfsVolumeSource() *GlusterfsVolumeSourceApplyConfiguration {
	return &GlusterfsVolumeSourceApplyConfiguration{}
}

// WithEndpointsName sets the EndpointsName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndpointsName field is set to the value of the last call.
func (b *GlusterfsVolumeSourceApplyConfiguration) WithEndpointsName(value string) *GlusterfsVolumeSourceApplyConfiguration {
	b.EndpointsName = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *GlusterfsVolumeSourceApplyConfiguration) WithPath(value string) *GlusterfsVolumeSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *GlusterfsVolumeSourceApplyConfiguration) WithReadOnly(value bool) *GlusterfsVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
