// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	types "k8s.io/apimachinery/pkg/types"
)

// ObjectReferenceApplyConfiguration represents a declarative configuration of the ObjectReference type for use
// with apply.
type ObjectReferenceApplyConfiguration struct {
	Kind            *string    `json:"kind,omitempty"`
	Namespace       *string    `json:"namespace,omitempty"`
	Name            *string    `json:"name,omitempty"`
	UID             *types.UID `json:"uid,omitempty"`
	APIVersion      *string    `json:"apiVersion,omitempty"`
	ResourceVersion *string    `json:"resourceVersion,omitempty"`
	FieldPath       *string    `json:"fieldPath,omitempty"`
}

// ObjectReferenceApplyConfiguration constructs a declarative configuration of the ObjectReference type for use with
// apply.
func ObjectReference() *ObjectReferenceApplyConfiguration {
	return &ObjectReferenceApplyConfiguration{}
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithKind(value string) *ObjectReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithNamespace(value string) *ObjectReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithName(value string) *ObjectReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithUID(value types.UID) *ObjectReferenceApplyConfiguration {
	b.UID = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithAPIVersion(value string) *ObjectReferenceApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithResourceVersion(value string) *ObjectReferenceApplyConfiguration {
	b.ResourceVersion = &value
	return b
}

// WithFieldPath sets the FieldPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FieldPath field is set to the value of the last call.
func (b *ObjectReferenceApplyConfiguration) WithFieldPath(value string) *ObjectReferenceApplyConfiguration {
	b.FieldPath = &value
	return b
}
