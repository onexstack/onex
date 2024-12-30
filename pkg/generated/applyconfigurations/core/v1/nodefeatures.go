// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// NodeFeaturesApplyConfiguration represents a declarative configuration of the NodeFeatures type for use
// with apply.
type NodeFeaturesApplyConfiguration struct {
	SupplementalGroupsPolicy *bool `json:"supplementalGroupsPolicy,omitempty"`
}

// NodeFeaturesApplyConfiguration constructs a declarative configuration of the NodeFeatures type for use with
// apply.
func NodeFeatures() *NodeFeaturesApplyConfiguration {
	return &NodeFeaturesApplyConfiguration{}
}

// WithSupplementalGroupsPolicy sets the SupplementalGroupsPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SupplementalGroupsPolicy field is set to the value of the last call.
func (b *NodeFeaturesApplyConfiguration) WithSupplementalGroupsPolicy(value bool) *NodeFeaturesApplyConfiguration {
	b.SupplementalGroupsPolicy = &value
	return b
}
