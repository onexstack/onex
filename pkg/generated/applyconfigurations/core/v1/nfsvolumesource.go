// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NFSVolumeSourceApplyConfiguration represents a declarative configuration of the NFSVolumeSource type for use
// with apply.
type NFSVolumeSourceApplyConfiguration struct {
	Server   *string `json:"server,omitempty"`
	Path     *string `json:"path,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
}

// NFSVolumeSourceApplyConfiguration constructs a declarative configuration of the NFSVolumeSource type for use with
// apply.
func NFSVolumeSource() *NFSVolumeSourceApplyConfiguration {
	return &NFSVolumeSourceApplyConfiguration{}
}

// WithServer sets the Server field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Server field is set to the value of the last call.
func (b *NFSVolumeSourceApplyConfiguration) WithServer(value string) *NFSVolumeSourceApplyConfiguration {
	b.Server = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *NFSVolumeSourceApplyConfiguration) WithPath(value string) *NFSVolumeSourceApplyConfiguration {
	b.Path = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *NFSVolumeSourceApplyConfiguration) WithReadOnly(value bool) *NFSVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
