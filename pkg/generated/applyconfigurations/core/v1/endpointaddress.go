// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// EndpointAddressApplyConfiguration represents a declarative configuration of the EndpointAddress type for use
// with apply.
type EndpointAddressApplyConfiguration struct {
	IP        *string                            `json:"ip,omitempty"`
	Hostname  *string                            `json:"hostname,omitempty"`
	NodeName  *string                            `json:"nodeName,omitempty"`
	TargetRef *ObjectReferenceApplyConfiguration `json:"targetRef,omitempty"`
}

// EndpointAddressApplyConfiguration constructs a declarative configuration of the EndpointAddress type for use with
// apply.
func EndpointAddress() *EndpointAddressApplyConfiguration {
	return &EndpointAddressApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *EndpointAddressApplyConfiguration) WithIP(value string) *EndpointAddressApplyConfiguration {
	b.IP = &value
	return b
}

// WithHostname sets the Hostname field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hostname field is set to the value of the last call.
func (b *EndpointAddressApplyConfiguration) WithHostname(value string) *EndpointAddressApplyConfiguration {
	b.Hostname = &value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *EndpointAddressApplyConfiguration) WithNodeName(value string) *EndpointAddressApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithTargetRef sets the TargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetRef field is set to the value of the last call.
func (b *EndpointAddressApplyConfiguration) WithTargetRef(value *ObjectReferenceApplyConfiguration) *EndpointAddressApplyConfiguration {
	b.TargetRef = value
	return b
}
