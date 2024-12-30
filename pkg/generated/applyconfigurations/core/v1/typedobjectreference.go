// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// TypedObjectReferenceApplyConfiguration represents a declarative configuration of the TypedObjectReference type for use
// with apply.
type TypedObjectReferenceApplyConfiguration struct {
	APIGroup  *string `json:"apiGroup,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// TypedObjectReferenceApplyConfiguration constructs a declarative configuration of the TypedObjectReference type for use with
// apply.
func TypedObjectReference() *TypedObjectReferenceApplyConfiguration {
	return &TypedObjectReferenceApplyConfiguration{}
}

// WithAPIGroup sets the APIGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIGroup field is set to the value of the last call.
func (b *TypedObjectReferenceApplyConfiguration) WithAPIGroup(value string) *TypedObjectReferenceApplyConfiguration {
	b.APIGroup = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *TypedObjectReferenceApplyConfiguration) WithKind(value string) *TypedObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *TypedObjectReferenceApplyConfiguration) WithName(value string) *TypedObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *TypedObjectReferenceApplyConfiguration) WithNamespace(value string) *TypedObjectReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}
