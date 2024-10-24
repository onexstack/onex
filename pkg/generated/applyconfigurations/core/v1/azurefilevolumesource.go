// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AzureFileVolumeSourceApplyConfiguration represents a declarative configuration of the AzureFileVolumeSource type for use
// with apply.
type AzureFileVolumeSourceApplyConfiguration struct {
	SecretName *string `json:"secretName,omitempty"`
	ShareName  *string `json:"shareName,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty"`
}

// AzureFileVolumeSourceApplyConfiguration constructs a declarative configuration of the AzureFileVolumeSource type for use with
// apply.
func AzureFileVolumeSource() *AzureFileVolumeSourceApplyConfiguration {
	return &AzureFileVolumeSourceApplyConfiguration{}
}

// WithSecretName sets the SecretName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretName field is set to the value of the last call.
func (b *AzureFileVolumeSourceApplyConfiguration) WithSecretName(value string) *AzureFileVolumeSourceApplyConfiguration {
	b.SecretName = &value
	return b
}

// WithShareName sets the ShareName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ShareName field is set to the value of the last call.
func (b *AzureFileVolumeSourceApplyConfiguration) WithShareName(value string) *AzureFileVolumeSourceApplyConfiguration {
	b.ShareName = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *AzureFileVolumeSourceApplyConfiguration) WithReadOnly(value bool) *AzureFileVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
