// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventSeriesApplyConfiguration represents a declarative configuration of the EventSeries type for use
// with apply.
type EventSeriesApplyConfiguration struct {
	Count            *int32            `json:"count,omitempty"`
	LastObservedTime *metav1.MicroTime `json:"lastObservedTime,omitempty"`
}

// EventSeriesApplyConfiguration constructs a declarative configuration of the EventSeries type for use with
// apply.
func EventSeries() *EventSeriesApplyConfiguration {
	return &EventSeriesApplyConfiguration{}
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *EventSeriesApplyConfiguration) WithCount(value int32) *EventSeriesApplyConfiguration {
	b.Count = &value
	return b
}

// WithLastObservedTime sets the LastObservedTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastObservedTime field is set to the value of the last call.
func (b *EventSeriesApplyConfiguration) WithLastObservedTime(value metav1.MicroTime) *EventSeriesApplyConfiguration {
	b.LastObservedTime = &value
	return b
}
