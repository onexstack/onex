// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// WeightedPodAffinityTermApplyConfiguration represents a declarative configuration of the WeightedPodAffinityTerm type for use
// with apply.
type WeightedPodAffinityTermApplyConfiguration struct {
	Weight          *int32                             `json:"weight,omitempty"`
	PodAffinityTerm *PodAffinityTermApplyConfiguration `json:"podAffinityTerm,omitempty"`
}

// WeightedPodAffinityTermApplyConfiguration constructs a declarative configuration of the WeightedPodAffinityTerm type for use with
// apply.
func WeightedPodAffinityTerm() *WeightedPodAffinityTermApplyConfiguration {
	return &WeightedPodAffinityTermApplyConfiguration{}
}

// WithWeight sets the Weight field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Weight field is set to the value of the last call.
func (b *WeightedPodAffinityTermApplyConfiguration) WithWeight(value int32) *WeightedPodAffinityTermApplyConfiguration {
	b.Weight = &value
	return b
}

// WithPodAffinityTerm sets the PodAffinityTerm field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodAffinityTerm field is set to the value of the last call.
func (b *WeightedPodAffinityTermApplyConfiguration) WithPodAffinityTerm(value *PodAffinityTermApplyConfiguration) *WeightedPodAffinityTermApplyConfiguration {
	b.PodAffinityTerm = value
	return b
}
