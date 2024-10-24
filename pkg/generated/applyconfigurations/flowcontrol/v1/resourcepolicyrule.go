// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ResourcePolicyRuleApplyConfiguration represents a declarative configuration of the ResourcePolicyRule type for use
// with apply.
type ResourcePolicyRuleApplyConfiguration struct {
	Verbs        []string `json:"verbs,omitempty"`
	APIGroups    []string `json:"apiGroups,omitempty"`
	Resources    []string `json:"resources,omitempty"`
	ClusterScope *bool    `json:"clusterScope,omitempty"`
	Namespaces   []string `json:"namespaces,omitempty"`
}

// ResourcePolicyRuleApplyConfiguration constructs a declarative configuration of the ResourcePolicyRule type for use with
// apply.
func ResourcePolicyRule() *ResourcePolicyRuleApplyConfiguration {
	return &ResourcePolicyRuleApplyConfiguration{}
}

// WithVerbs adds the given value to the Verbs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Verbs field.
func (b *ResourcePolicyRuleApplyConfiguration) WithVerbs(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Verbs = append(b.Verbs, values[i])
	}
	return b
}

// WithAPIGroups adds the given value to the APIGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIGroups field.
func (b *ResourcePolicyRuleApplyConfiguration) WithAPIGroups(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.APIGroups = append(b.APIGroups, values[i])
	}
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *ResourcePolicyRuleApplyConfiguration) WithResources(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithClusterScope sets the ClusterScope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterScope field is set to the value of the last call.
func (b *ResourcePolicyRuleApplyConfiguration) WithClusterScope(value bool) *ResourcePolicyRuleApplyConfiguration {
	b.ClusterScope = &value
	return b
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *ResourcePolicyRuleApplyConfiguration) WithNamespaces(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}
