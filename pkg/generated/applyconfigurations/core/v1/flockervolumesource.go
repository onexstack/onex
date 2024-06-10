// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// FlockerVolumeSourceApplyConfiguration represents an declarative configuration of the FlockerVolumeSource type for use
// with apply.
type FlockerVolumeSourceApplyConfiguration struct {
	DatasetName *string `json:"datasetName,omitempty"`
	DatasetUUID *string `json:"datasetUUID,omitempty"`
}

// FlockerVolumeSourceApplyConfiguration constructs an declarative configuration of the FlockerVolumeSource type for use with
// apply.
func FlockerVolumeSource() *FlockerVolumeSourceApplyConfiguration {
	return &FlockerVolumeSourceApplyConfiguration{}
}

// WithDatasetName sets the DatasetName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DatasetName field is set to the value of the last call.
func (b *FlockerVolumeSourceApplyConfiguration) WithDatasetName(value string) *FlockerVolumeSourceApplyConfiguration {
	b.DatasetName = &value
	return b
}

// WithDatasetUUID sets the DatasetUUID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DatasetUUID field is set to the value of the last call.
func (b *FlockerVolumeSourceApplyConfiguration) WithDatasetUUID(value string) *FlockerVolumeSourceApplyConfiguration {
	b.DatasetUUID = &value
	return b
}
