// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// HostIPApplyConfiguration represents a declarative configuration of the HostIP type for use
// with apply.
type HostIPApplyConfiguration struct {
	IP *string `json:"ip,omitempty"`
}

// HostIPApplyConfiguration constructs a declarative configuration of the HostIP type for use with
// apply.
func HostIP() *HostIPApplyConfiguration {
	return &HostIPApplyConfiguration{}
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *HostIPApplyConfiguration) WithIP(value string) *HostIPApplyConfiguration {
	b.IP = &value
	return b
}
