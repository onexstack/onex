// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/flowcontrol/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FlowSchemaConditionApplyConfiguration represents a declarative configuration of the FlowSchemaCondition type for use
// with apply.
type FlowSchemaConditionApplyConfiguration struct {
	Type               *v1.FlowSchemaConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus         `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                `json:"lastTransitionTime,omitempty"`
	Reason             *string                     `json:"reason,omitempty"`
	Message            *string                     `json:"message,omitempty"`
}

// FlowSchemaConditionApplyConfiguration constructs a declarative configuration of the FlowSchemaCondition type for use with
// apply.
func FlowSchemaCondition() *FlowSchemaConditionApplyConfiguration {
	return &FlowSchemaConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *FlowSchemaConditionApplyConfiguration) WithType(value v1.FlowSchemaConditionType) *FlowSchemaConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *FlowSchemaConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *FlowSchemaConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *FlowSchemaConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *FlowSchemaConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *FlowSchemaConditionApplyConfiguration) WithReason(value string) *FlowSchemaConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *FlowSchemaConditionApplyConfiguration) WithMessage(value string) *FlowSchemaConditionApplyConfiguration {
	b.Message = &value
	return b
}
