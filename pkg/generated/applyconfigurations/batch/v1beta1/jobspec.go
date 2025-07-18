// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	batchv1beta1 "github.com/onexstack/onex/pkg/apis/batch/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// JobSpecApplyConfiguration represents a declarative configuration of the JobSpec type for use
// with apply.
type JobSpecApplyConfiguration struct {
	Type                  *batchv1beta1.JobType `json:"type,omitempty"`
	Suspend               *bool                 `json:"suspend,omitempty"`
	ActiveDeadlineSeconds *int64                `json:"activeDeadlineSeconds,omitempty"`
	ProviderSpec          *runtime.RawExtension `json:"providerSpec,omitempty"`
}

// JobSpecApplyConfiguration constructs a declarative configuration of the JobSpec type for use with
// apply.
func JobSpec() *JobSpecApplyConfiguration {
	return &JobSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithType(value batchv1beta1.JobType) *JobSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithSuspend sets the Suspend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Suspend field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithSuspend(value bool) *JobSpecApplyConfiguration {
	b.Suspend = &value
	return b
}

// WithActiveDeadlineSeconds sets the ActiveDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActiveDeadlineSeconds field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithActiveDeadlineSeconds(value int64) *JobSpecApplyConfiguration {
	b.ActiveDeadlineSeconds = &value
	return b
}

// WithProviderSpec sets the ProviderSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProviderSpec field is set to the value of the last call.
func (b *JobSpecApplyConfiguration) WithProviderSpec(value runtime.RawExtension) *JobSpecApplyConfiguration {
	b.ProviderSpec = &value
	return b
}
