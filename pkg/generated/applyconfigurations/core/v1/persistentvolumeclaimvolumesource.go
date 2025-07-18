// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PersistentVolumeClaimVolumeSourceApplyConfiguration represents a declarative configuration of the PersistentVolumeClaimVolumeSource type for use
// with apply.
type PersistentVolumeClaimVolumeSourceApplyConfiguration struct {
	ClaimName *string `json:"claimName,omitempty"`
	ReadOnly  *bool   `json:"readOnly,omitempty"`
}

// PersistentVolumeClaimVolumeSourceApplyConfiguration constructs a declarative configuration of the PersistentVolumeClaimVolumeSource type for use with
// apply.
func PersistentVolumeClaimVolumeSource() *PersistentVolumeClaimVolumeSourceApplyConfiguration {
	return &PersistentVolumeClaimVolumeSourceApplyConfiguration{}
}

// WithClaimName sets the ClaimName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClaimName field is set to the value of the last call.
func (b *PersistentVolumeClaimVolumeSourceApplyConfiguration) WithClaimName(value string) *PersistentVolumeClaimVolumeSourceApplyConfiguration {
	b.ClaimName = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *PersistentVolumeClaimVolumeSourceApplyConfiguration) WithReadOnly(value bool) *PersistentVolumeClaimVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
