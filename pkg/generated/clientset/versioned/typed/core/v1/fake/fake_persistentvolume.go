// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	corev1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePersistentVolumes implements PersistentVolumeInterface
type FakePersistentVolumes struct {
	Fake *FakeCoreV1
}

var persistentvolumesResource = v1.SchemeGroupVersion.WithResource("persistentvolumes")

var persistentvolumesKind = v1.SchemeGroupVersion.WithKind("PersistentVolume")

// Get takes name of the persistentVolume, and returns the corresponding persistentVolume object, and an error if there is any.
func (c *FakePersistentVolumes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PersistentVolume, err error) {
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(persistentvolumesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}

// List takes label and field selectors, and returns the list of PersistentVolumes that match those selectors.
func (c *FakePersistentVolumes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PersistentVolumeList, err error) {
	emptyResult := &v1.PersistentVolumeList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(persistentvolumesResource, persistentvolumesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PersistentVolumeList{ListMeta: obj.(*v1.PersistentVolumeList).ListMeta}
	for _, item := range obj.(*v1.PersistentVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested persistentVolumes.
func (c *FakePersistentVolumes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(persistentvolumesResource, opts))
}

// Create takes the representation of a persistentVolume and creates it.  Returns the server's representation of the persistentVolume, and an error, if there is any.
func (c *FakePersistentVolumes) Create(ctx context.Context, persistentVolume *v1.PersistentVolume, opts metav1.CreateOptions) (result *v1.PersistentVolume, err error) {
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(persistentvolumesResource, persistentVolume, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}

// Update takes the representation of a persistentVolume and updates it. Returns the server's representation of the persistentVolume, and an error, if there is any.
func (c *FakePersistentVolumes) Update(ctx context.Context, persistentVolume *v1.PersistentVolume, opts metav1.UpdateOptions) (result *v1.PersistentVolume, err error) {
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(persistentvolumesResource, persistentVolume, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePersistentVolumes) UpdateStatus(ctx context.Context, persistentVolume *v1.PersistentVolume, opts metav1.UpdateOptions) (result *v1.PersistentVolume, err error) {
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(persistentvolumesResource, "status", persistentVolume, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}

// Delete takes name of the persistentVolume and deletes it. Returns an error if one occurs.
func (c *FakePersistentVolumes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(persistentvolumesResource, name, opts), &v1.PersistentVolume{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePersistentVolumes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(persistentvolumesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PersistentVolumeList{})
	return err
}

// Patch applies the patch and returns the patched persistentVolume.
func (c *FakePersistentVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PersistentVolume, err error) {
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(persistentvolumesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied persistentVolume.
func (c *FakePersistentVolumes) Apply(ctx context.Context, persistentVolume *corev1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolume, err error) {
	if persistentVolume == nil {
		return nil, fmt.Errorf("persistentVolume provided to Apply must not be nil")
	}
	data, err := json.Marshal(persistentVolume)
	if err != nil {
		return nil, err
	}
	name := persistentVolume.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolume.Name must be provided to Apply")
	}
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(persistentvolumesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePersistentVolumes) ApplyStatus(ctx context.Context, persistentVolume *corev1.PersistentVolumeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PersistentVolume, err error) {
	if persistentVolume == nil {
		return nil, fmt.Errorf("persistentVolume provided to Apply must not be nil")
	}
	data, err := json.Marshal(persistentVolume)
	if err != nil {
		return nil, err
	}
	name := persistentVolume.Name
	if name == nil {
		return nil, fmt.Errorf("persistentVolume.Name must be provided to Apply")
	}
	emptyResult := &v1.PersistentVolume{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(persistentvolumesResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.PersistentVolume), err
}
