// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PodDNSConfigApplyConfiguration represents a declarative configuration of the PodDNSConfig type for use
// with apply.
type PodDNSConfigApplyConfiguration struct {
	Nameservers []string                               `json:"nameservers,omitempty"`
	Searches    []string                               `json:"searches,omitempty"`
	Options     []PodDNSConfigOptionApplyConfiguration `json:"options,omitempty"`
}

// PodDNSConfigApplyConfiguration constructs a declarative configuration of the PodDNSConfig type for use with
// apply.
func PodDNSConfig() *PodDNSConfigApplyConfiguration {
	return &PodDNSConfigApplyConfiguration{}
}

// WithNameservers adds the given value to the Nameservers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Nameservers field.
func (b *PodDNSConfigApplyConfiguration) WithNameservers(values ...string) *PodDNSConfigApplyConfiguration {
	for i := range values {
		b.Nameservers = append(b.Nameservers, values[i])
	}
	return b
}

// WithSearches adds the given value to the Searches field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Searches field.
func (b *PodDNSConfigApplyConfiguration) WithSearches(values ...string) *PodDNSConfigApplyConfiguration {
	for i := range values {
		b.Searches = append(b.Searches, values[i])
	}
	return b
}

// WithOptions adds the given value to the Options field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Options field.
func (b *PodDNSConfigApplyConfiguration) WithOptions(values ...*PodDNSConfigOptionApplyConfiguration) *PodDNSConfigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOptions")
		}
		b.Options = append(b.Options, *values[i])
	}
	return b
}
