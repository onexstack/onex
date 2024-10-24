// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// AttachedVolumeApplyConfiguration represents a declarative configuration of the AttachedVolume type for use
// with apply.
type AttachedVolumeApplyConfiguration struct {
	Name       *v1.UniqueVolumeName `json:"name,omitempty"`
	DevicePath *string              `json:"devicePath,omitempty"`
}

// AttachedVolumeApplyConfiguration constructs a declarative configuration of the AttachedVolume type for use with
// apply.
func AttachedVolume() *AttachedVolumeApplyConfiguration {
	return &AttachedVolumeApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *AttachedVolumeApplyConfiguration) WithName(value v1.UniqueVolumeName) *AttachedVolumeApplyConfiguration {
	b.Name = &value
	return b
}

// WithDevicePath sets the DevicePath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DevicePath field is set to the value of the last call.
func (b *AttachedVolumeApplyConfiguration) WithDevicePath(value string) *AttachedVolumeApplyConfiguration {
	b.DevicePath = &value
	return b
}
