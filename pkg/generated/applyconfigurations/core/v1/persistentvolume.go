// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	internal "github.com/superproj/onex/pkg/generated/applyconfigurations/internal"
	v1 "github.com/superproj/onex/pkg/generated/applyconfigurations/meta/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// PersistentVolumeApplyConfiguration represents an declarative configuration of the PersistentVolume type for use
// with apply.
type PersistentVolumeApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *PersistentVolumeSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *PersistentVolumeStatusApplyConfiguration `json:"status,omitempty"`
}

// PersistentVolume constructs an declarative configuration of the PersistentVolume type for use with
// apply.
func PersistentVolume(name string) *PersistentVolumeApplyConfiguration {
	b := &PersistentVolumeApplyConfiguration{}
	b.WithName(name)
	b.WithKind("PersistentVolume")
	b.WithAPIVersion("v1")
	return b
}

// ExtractPersistentVolume extracts the applied configuration owned by fieldManager from
// persistentVolume. If no managedFields are found in persistentVolume for fieldManager, a
// PersistentVolumeApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// persistentVolume must be a unmodified PersistentVolume API object that was retrieved from the Kubernetes API.
// ExtractPersistentVolume provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractPersistentVolume(persistentVolume *apicorev1.PersistentVolume, fieldManager string) (*PersistentVolumeApplyConfiguration, error) {
	return extractPersistentVolume(persistentVolume, fieldManager, "")
}

// ExtractPersistentVolumeStatus is the same as ExtractPersistentVolume except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractPersistentVolumeStatus(persistentVolume *apicorev1.PersistentVolume, fieldManager string) (*PersistentVolumeApplyConfiguration, error) {
	return extractPersistentVolume(persistentVolume, fieldManager, "status")
}

func extractPersistentVolume(persistentVolume *apicorev1.PersistentVolume, fieldManager string, subresource string) (*PersistentVolumeApplyConfiguration, error) {
	b := &PersistentVolumeApplyConfiguration{}
	err := managedfields.ExtractInto(persistentVolume, internal.Parser().Type("io.k8s.api.core.v1.PersistentVolume"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(persistentVolume.Name)

	b.WithKind("PersistentVolume")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithKind(value string) *PersistentVolumeApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithAPIVersion(value string) *PersistentVolumeApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithName(value string) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithGenerateName(value string) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithNamespace(value string) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithUID(value types.UID) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithResourceVersion(value string) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithGeneration(value int64) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithCreationTimestamp(value metav1.Time) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *PersistentVolumeApplyConfiguration) WithLabels(entries map[string]string) *PersistentVolumeApplyConfiguration {
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
func (b *PersistentVolumeApplyConfiguration) WithAnnotations(entries map[string]string) *PersistentVolumeApplyConfiguration {
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
func (b *PersistentVolumeApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *PersistentVolumeApplyConfiguration {
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
func (b *PersistentVolumeApplyConfiguration) WithFinalizers(values ...string) *PersistentVolumeApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *PersistentVolumeApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithSpec(value *PersistentVolumeSpecApplyConfiguration) *PersistentVolumeApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PersistentVolumeApplyConfiguration) WithStatus(value *PersistentVolumeStatusApplyConfiguration) *PersistentVolumeApplyConfiguration {
	b.Status = value
	return b
}
