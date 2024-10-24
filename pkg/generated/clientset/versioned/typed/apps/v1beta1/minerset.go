// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	json "encoding/json"
	"fmt"

	v1beta1 "github.com/superproj/onex/pkg/apis/apps/v1beta1"
	appsv1beta1 "github.com/superproj/onex/pkg/generated/applyconfigurations/apps/v1beta1"
	applyconfigurationsautoscalingv1 "github.com/superproj/onex/pkg/generated/applyconfigurations/autoscaling/v1"
	scheme "github.com/superproj/onex/pkg/generated/clientset/versioned/scheme"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// MinerSetsGetter has a method to return a MinerSetInterface.
// A group's client should implement this interface.
type MinerSetsGetter interface {
	MinerSets(namespace string) MinerSetInterface
}

// MinerSetInterface has methods to work with MinerSet resources.
type MinerSetInterface interface {
	Create(ctx context.Context, minerSet *v1beta1.MinerSet, opts v1.CreateOptions) (*v1beta1.MinerSet, error)
	Update(ctx context.Context, minerSet *v1beta1.MinerSet, opts v1.UpdateOptions) (*v1beta1.MinerSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, minerSet *v1beta1.MinerSet, opts v1.UpdateOptions) (*v1beta1.MinerSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.MinerSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.MinerSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MinerSet, err error)
	Apply(ctx context.Context, minerSet *appsv1beta1.MinerSetApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.MinerSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, minerSet *appsv1beta1.MinerSetApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.MinerSet, err error)
	GetScale(ctx context.Context, minerSetName string, options v1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, minerSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (*autoscalingv1.Scale, error)
	ApplyScale(ctx context.Context, minerSetName string, scale *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts v1.ApplyOptions) (*autoscalingv1.Scale, error)

	MinerSetExpansion
}

// minerSets implements MinerSetInterface
type minerSets struct {
	*gentype.ClientWithListAndApply[*v1beta1.MinerSet, *v1beta1.MinerSetList, *appsv1beta1.MinerSetApplyConfiguration]
}

// newMinerSets returns a MinerSets
func newMinerSets(c *AppsV1beta1Client, namespace string) *minerSets {
	return &minerSets{
		gentype.NewClientWithListAndApply[*v1beta1.MinerSet, *v1beta1.MinerSetList, *appsv1beta1.MinerSetApplyConfiguration](
			"minersets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.MinerSet { return &v1beta1.MinerSet{} },
			func() *v1beta1.MinerSetList { return &v1beta1.MinerSetList{} }),
	}
}

// GetScale takes name of the minerSet, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *minerSets) GetScale(ctx context.Context, minerSetName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("minersets").
		Name(minerSetName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *minerSets) UpdateScale(ctx context.Context, minerSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("minersets").
		Name(minerSetName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *minerSets) ApplyScale(ctx context.Context, minerSetName string, scale *applyconfigurationsautoscalingv1.ScaleApplyConfiguration, opts v1.ApplyOptions) (result *autoscalingv1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}

	result = &autoscalingv1.Scale{}
	err = c.GetClient().Patch(types.ApplyPatchType).
		Namespace(c.GetNamespace()).
		Resource("minersets").
		Name(minerSetName).
		SubResource("scale").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
