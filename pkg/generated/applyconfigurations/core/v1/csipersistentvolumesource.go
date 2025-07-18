// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CSIPersistentVolumeSourceApplyConfiguration represents a declarative configuration of the CSIPersistentVolumeSource type for use
// with apply.
type CSIPersistentVolumeSourceApplyConfiguration struct {
	Driver                     *string                            `json:"driver,omitempty"`
	VolumeHandle               *string                            `json:"volumeHandle,omitempty"`
	ReadOnly                   *bool                              `json:"readOnly,omitempty"`
	FSType                     *string                            `json:"fsType,omitempty"`
	VolumeAttributes           map[string]string                  `json:"volumeAttributes,omitempty"`
	ControllerPublishSecretRef *SecretReferenceApplyConfiguration `json:"controllerPublishSecretRef,omitempty"`
	NodeStageSecretRef         *SecretReferenceApplyConfiguration `json:"nodeStageSecretRef,omitempty"`
	NodePublishSecretRef       *SecretReferenceApplyConfiguration `json:"nodePublishSecretRef,omitempty"`
	ControllerExpandSecretRef  *SecretReferenceApplyConfiguration `json:"controllerExpandSecretRef,omitempty"`
	NodeExpandSecretRef        *SecretReferenceApplyConfiguration `json:"nodeExpandSecretRef,omitempty"`
}

// CSIPersistentVolumeSourceApplyConfiguration constructs a declarative configuration of the CSIPersistentVolumeSource type for use with
// apply.
func CSIPersistentVolumeSource() *CSIPersistentVolumeSourceApplyConfiguration {
	return &CSIPersistentVolumeSourceApplyConfiguration{}
}

// WithDriver sets the Driver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Driver field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithDriver(value string) *CSIPersistentVolumeSourceApplyConfiguration {
	b.Driver = &value
	return b
}

// WithVolumeHandle sets the VolumeHandle field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeHandle field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithVolumeHandle(value string) *CSIPersistentVolumeSourceApplyConfiguration {
	b.VolumeHandle = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithReadOnly(value bool) *CSIPersistentVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithFSType(value string) *CSIPersistentVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithVolumeAttributes puts the entries into the VolumeAttributes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the VolumeAttributes field,
// overwriting an existing map entries in VolumeAttributes field with the same key.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithVolumeAttributes(entries map[string]string) *CSIPersistentVolumeSourceApplyConfiguration {
	if b.VolumeAttributes == nil && len(entries) > 0 {
		b.VolumeAttributes = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.VolumeAttributes[k] = v
	}
	return b
}

// WithControllerPublishSecretRef sets the ControllerPublishSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControllerPublishSecretRef field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithControllerPublishSecretRef(value *SecretReferenceApplyConfiguration) *CSIPersistentVolumeSourceApplyConfiguration {
	b.ControllerPublishSecretRef = value
	return b
}

// WithNodeStageSecretRef sets the NodeStageSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeStageSecretRef field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithNodeStageSecretRef(value *SecretReferenceApplyConfiguration) *CSIPersistentVolumeSourceApplyConfiguration {
	b.NodeStageSecretRef = value
	return b
}

// WithNodePublishSecretRef sets the NodePublishSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodePublishSecretRef field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithNodePublishSecretRef(value *SecretReferenceApplyConfiguration) *CSIPersistentVolumeSourceApplyConfiguration {
	b.NodePublishSecretRef = value
	return b
}

// WithControllerExpandSecretRef sets the ControllerExpandSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ControllerExpandSecretRef field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithControllerExpandSecretRef(value *SecretReferenceApplyConfiguration) *CSIPersistentVolumeSourceApplyConfiguration {
	b.ControllerExpandSecretRef = value
	return b
}

// WithNodeExpandSecretRef sets the NodeExpandSecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeExpandSecretRef field is set to the value of the last call.
func (b *CSIPersistentVolumeSourceApplyConfiguration) WithNodeExpandSecretRef(value *SecretReferenceApplyConfiguration) *CSIPersistentVolumeSourceApplyConfiguration {
	b.NodeExpandSecretRef = value
	return b
}
