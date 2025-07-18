// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// AppArmorProfileApplyConfiguration represents a declarative configuration of the AppArmorProfile type for use
// with apply.
type AppArmorProfileApplyConfiguration struct {
	Type             *corev1.AppArmorProfileType `json:"type,omitempty"`
	LocalhostProfile *string                     `json:"localhostProfile,omitempty"`
}

// AppArmorProfileApplyConfiguration constructs a declarative configuration of the AppArmorProfile type for use with
// apply.
func AppArmorProfile() *AppArmorProfileApplyConfiguration {
	return &AppArmorProfileApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AppArmorProfileApplyConfiguration) WithType(value corev1.AppArmorProfileType) *AppArmorProfileApplyConfiguration {
	b.Type = &value
	return b
}

// WithLocalhostProfile sets the LocalhostProfile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LocalhostProfile field is set to the value of the last call.
func (b *AppArmorProfileApplyConfiguration) WithLocalhostProfile(value string) *AppArmorProfileApplyConfiguration {
	b.LocalhostProfile = &value
	return b
}
