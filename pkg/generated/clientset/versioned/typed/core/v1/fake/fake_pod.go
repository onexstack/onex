// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	corev1 "github.com/superproj/onex/pkg/generated/applyconfigurations/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePods implements PodInterface
type FakePods struct {
	Fake *FakeCoreV1
	ns   string
}

var podsResource = v1.SchemeGroupVersion.WithResource("pods")

var podsKind = v1.SchemeGroupVersion.WithKind("Pod")

// Get takes name of the pod, and returns the corresponding pod object, and an error if there is any.
func (c *FakePods) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Pod, err error) {
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(podsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// List takes label and field selectors, and returns the list of Pods that match those selectors.
func (c *FakePods) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PodList, err error) {
	emptyResult := &v1.PodList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(podsResource, podsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PodList{ListMeta: obj.(*v1.PodList).ListMeta}
	for _, item := range obj.(*v1.PodList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pods.
func (c *FakePods) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(podsResource, c.ns, opts))

}

// Create takes the representation of a pod and creates it.  Returns the server's representation of the pod, and an error, if there is any.
func (c *FakePods) Create(ctx context.Context, pod *v1.Pod, opts metav1.CreateOptions) (result *v1.Pod, err error) {
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(podsResource, c.ns, pod, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// Update takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *FakePods) Update(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (result *v1.Pod, err error) {
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(podsResource, c.ns, pod, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePods) UpdateStatus(ctx context.Context, pod *v1.Pod, opts metav1.UpdateOptions) (result *v1.Pod, err error) {
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(podsResource, "status", c.ns, pod, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// Delete takes name of the pod and deletes it. Returns an error if one occurs.
func (c *FakePods) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(podsResource, c.ns, name, opts), &v1.Pod{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePods) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(podsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PodList{})
	return err
}

// Patch applies the patch and returns the patched pod.
func (c *FakePods) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Pod, err error) {
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(podsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied pod.
func (c *FakePods) Apply(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	if pod == nil {
		return nil, fmt.Errorf("pod provided to Apply must not be nil")
	}
	data, err := json.Marshal(pod)
	if err != nil {
		return nil, err
	}
	name := pod.Name
	if name == nil {
		return nil, fmt.Errorf("pod.Name must be provided to Apply")
	}
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(podsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakePods) ApplyStatus(ctx context.Context, pod *corev1.PodApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Pod, err error) {
	if pod == nil {
		return nil, fmt.Errorf("pod provided to Apply must not be nil")
	}
	data, err := json.Marshal(pod)
	if err != nil {
		return nil, err
	}
	name := pod.Name
	if name == nil {
		return nil, fmt.Errorf("pod.Name must be provided to Apply")
	}
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(podsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}

// UpdateEphemeralContainers takes the representation of a pod and updates it. Returns the server's representation of the pod, and an error, if there is any.
func (c *FakePods) UpdateEphemeralContainers(ctx context.Context, podName string, pod *v1.Pod, opts metav1.UpdateOptions) (result *v1.Pod, err error) {
	emptyResult := &v1.Pod{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(podsResource, "ephemeralcontainers", c.ns, pod, opts), &v1.Pod{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Pod), err
}
