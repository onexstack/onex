// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PodResourceClaimApplyConfiguration represents a declarative configuration of the PodResourceClaim type for use
// with apply.
type PodResourceClaimApplyConfiguration struct {
	Name                      *string `json:"name,omitempty"`
	ResourceClaimName         *string `json:"resourceClaimName,omitempty"`
	ResourceClaimTemplateName *string `json:"resourceClaimTemplateName,omitempty"`
}

// PodResourceClaimApplyConfiguration constructs a declarative configuration of the PodResourceClaim type for use with
// apply.
func PodResourceClaim() *PodResourceClaimApplyConfiguration {
	return &PodResourceClaimApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PodResourceClaimApplyConfiguration) WithName(value string) *PodResourceClaimApplyConfiguration {
	b.Name = &value
	return b
}

// WithResourceClaimName sets the ResourceClaimName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceClaimName field is set to the value of the last call.
func (b *PodResourceClaimApplyConfiguration) WithResourceClaimName(value string) *PodResourceClaimApplyConfiguration {
	b.ResourceClaimName = &value
	return b
}

// WithResourceClaimTemplateName sets the ResourceClaimTemplateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceClaimTemplateName field is set to the value of the last call.
func (b *PodResourceClaimApplyConfiguration) WithResourceClaimTemplateName(value string) *PodResourceClaimApplyConfiguration {
	b.ResourceClaimTemplateName = &value
	return b
}
