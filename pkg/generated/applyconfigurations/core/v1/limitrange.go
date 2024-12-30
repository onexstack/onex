// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	internal "github.com/onexstack/onex/pkg/generated/applyconfigurations/internal"
	v1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/meta/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// LimitRangeApplyConfiguration represents a declarative configuration of the LimitRange type for use
// with apply.
type LimitRangeApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *LimitRangeSpecApplyConfiguration `json:"spec,omitempty"`
}

// LimitRange constructs a declarative configuration of the LimitRange type for use with
// apply.
func LimitRange(name, namespace string) *LimitRangeApplyConfiguration {
	b := &LimitRangeApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("LimitRange")
	b.WithAPIVersion("v1")
	return b
}

// ExtractLimitRange extracts the applied configuration owned by fieldManager from
// limitRange. If no managedFields are found in limitRange for fieldManager, a
// LimitRangeApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// limitRange must be a unmodified LimitRange API object that was retrieved from the Kubernetes API.
// ExtractLimitRange provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractLimitRange(limitRange *apicorev1.LimitRange, fieldManager string) (*LimitRangeApplyConfiguration, error) {
	return extractLimitRange(limitRange, fieldManager, "")
}

// ExtractLimitRangeStatus is the same as ExtractLimitRange except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractLimitRangeStatus(limitRange *apicorev1.LimitRange, fieldManager string) (*LimitRangeApplyConfiguration, error) {
	return extractLimitRange(limitRange, fieldManager, "status")
}

func extractLimitRange(limitRange *apicorev1.LimitRange, fieldManager string, subresource string) (*LimitRangeApplyConfiguration, error) {
	b := &LimitRangeApplyConfiguration{}
	err := managedfields.ExtractInto(limitRange, internal.Parser().Type("io.k8s.api.core.v1.LimitRange"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(limitRange.Name)
	b.WithNamespace(limitRange.Namespace)

	b.WithKind("LimitRange")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithKind(value string) *LimitRangeApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithAPIVersion(value string) *LimitRangeApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithName(value string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithGenerateName(value string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithNamespace(value string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithUID(value types.UID) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithResourceVersion(value string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithGeneration(value int64) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithCreationTimestamp(value metav1.Time) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *LimitRangeApplyConfiguration) WithLabels(entries map[string]string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *LimitRangeApplyConfiguration) WithAnnotations(entries map[string]string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *LimitRangeApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *LimitRangeApplyConfiguration) WithFinalizers(values ...string) *LimitRangeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *LimitRangeApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *LimitRangeApplyConfiguration) WithSpec(value *LimitRangeSpecApplyConfiguration) *LimitRangeApplyConfiguration {
	b.Spec = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *LimitRangeApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
