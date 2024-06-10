// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ScaleIOVolumeSourceApplyConfiguration represents an declarative configuration of the ScaleIOVolumeSource type for use
// with apply.
type ScaleIOVolumeSourceApplyConfiguration struct {
	Gateway          *string                                 `json:"gateway,omitempty"`
	System           *string                                 `json:"system,omitempty"`
	SecretRef        *LocalObjectReferenceApplyConfiguration `json:"secretRef,omitempty"`
	SSLEnabled       *bool                                   `json:"sslEnabled,omitempty"`
	ProtectionDomain *string                                 `json:"protectionDomain,omitempty"`
	StoragePool      *string                                 `json:"storagePool,omitempty"`
	StorageMode      *string                                 `json:"storageMode,omitempty"`
	VolumeName       *string                                 `json:"volumeName,omitempty"`
	FSType           *string                                 `json:"fsType,omitempty"`
	ReadOnly         *bool                                   `json:"readOnly,omitempty"`
}

// ScaleIOVolumeSourceApplyConfiguration constructs an declarative configuration of the ScaleIOVolumeSource type for use with
// apply.
func ScaleIOVolumeSource() *ScaleIOVolumeSourceApplyConfiguration {
	return &ScaleIOVolumeSourceApplyConfiguration{}
}

// WithGateway sets the Gateway field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Gateway field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithGateway(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.Gateway = &value
	return b
}

// WithSystem sets the System field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the System field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithSystem(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.System = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithSecretRef(value *LocalObjectReferenceApplyConfiguration) *ScaleIOVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}

// WithSSLEnabled sets the SSLEnabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SSLEnabled field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithSSLEnabled(value bool) *ScaleIOVolumeSourceApplyConfiguration {
	b.SSLEnabled = &value
	return b
}

// WithProtectionDomain sets the ProtectionDomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProtectionDomain field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithProtectionDomain(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.ProtectionDomain = &value
	return b
}

// WithStoragePool sets the StoragePool field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StoragePool field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithStoragePool(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.StoragePool = &value
	return b
}

// WithStorageMode sets the StorageMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageMode field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithStorageMode(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.StorageMode = &value
	return b
}

// WithVolumeName sets the VolumeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeName field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithVolumeName(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.VolumeName = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithFSType(value string) *ScaleIOVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *ScaleIOVolumeSourceApplyConfiguration) WithReadOnly(value bool) *ScaleIOVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}
