// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	appsv1beta1 "github.com/superproj/onex/pkg/apis/apps/v1beta1"
	internal "github.com/superproj/onex/pkg/generated/applyconfigurations/internal"
	v1 "github.com/superproj/onex/pkg/generated/applyconfigurations/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// ChargeRequestApplyConfiguration represents a declarative configuration of the ChargeRequest type for use
// with apply.
type ChargeRequestApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *ChargeRequestSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *ChargeRequestStatusApplyConfiguration `json:"status,omitempty"`
}

// ChargeRequest constructs a declarative configuration of the ChargeRequest type for use with
// apply.
func ChargeRequest(name, namespace string) *ChargeRequestApplyConfiguration {
	b := &ChargeRequestApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("ChargeRequest")
	b.WithAPIVersion("apps.onex.io/v1beta1")
	return b
}

// ExtractChargeRequest extracts the applied configuration owned by fieldManager from
// chargeRequest. If no managedFields are found in chargeRequest for fieldManager, a
// ChargeRequestApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// chargeRequest must be a unmodified ChargeRequest API object that was retrieved from the Kubernetes API.
// ExtractChargeRequest provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractChargeRequest(chargeRequest *appsv1beta1.ChargeRequest, fieldManager string) (*ChargeRequestApplyConfiguration, error) {
	return extractChargeRequest(chargeRequest, fieldManager, "")
}

// ExtractChargeRequestStatus is the same as ExtractChargeRequest except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractChargeRequestStatus(chargeRequest *appsv1beta1.ChargeRequest, fieldManager string) (*ChargeRequestApplyConfiguration, error) {
	return extractChargeRequest(chargeRequest, fieldManager, "status")
}

func extractChargeRequest(chargeRequest *appsv1beta1.ChargeRequest, fieldManager string, subresource string) (*ChargeRequestApplyConfiguration, error) {
	b := &ChargeRequestApplyConfiguration{}
	err := managedfields.ExtractInto(chargeRequest, internal.Parser().Type("com.github.superproj.onex.pkg.apis.apps.v1beta1.ChargeRequest"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(chargeRequest.Name)
	b.WithNamespace(chargeRequest.Namespace)

	b.WithKind("ChargeRequest")
	b.WithAPIVersion("apps.onex.io/v1beta1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithKind(value string) *ChargeRequestApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithAPIVersion(value string) *ChargeRequestApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithName(value string) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithGenerateName(value string) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithNamespace(value string) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithUID(value types.UID) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithResourceVersion(value string) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithGeneration(value int64) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithCreationTimestamp(value metav1.Time) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *ChargeRequestApplyConfiguration) WithLabels(entries map[string]string) *ChargeRequestApplyConfiguration {
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
func (b *ChargeRequestApplyConfiguration) WithAnnotations(entries map[string]string) *ChargeRequestApplyConfiguration {
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
func (b *ChargeRequestApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *ChargeRequestApplyConfiguration {
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
func (b *ChargeRequestApplyConfiguration) WithFinalizers(values ...string) *ChargeRequestApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *ChargeRequestApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithSpec(value *ChargeRequestSpecApplyConfiguration) *ChargeRequestApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ChargeRequestApplyConfiguration) WithStatus(value *ChargeRequestStatusApplyConfiguration) *ChargeRequestApplyConfiguration {
	b.Status = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *ChargeRequestApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
