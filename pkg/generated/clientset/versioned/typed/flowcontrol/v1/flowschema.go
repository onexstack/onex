// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	flowcontrolv1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/flowcontrol/v1"
	scheme "github.com/onexstack/onex/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/api/flowcontrol/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// FlowSchemasGetter has a method to return a FlowSchemaInterface.
// A group's client should implement this interface.
type FlowSchemasGetter interface {
	FlowSchemas() FlowSchemaInterface
}

// FlowSchemaInterface has methods to work with FlowSchema resources.
type FlowSchemaInterface interface {
	Create(ctx context.Context, flowSchema *v1.FlowSchema, opts metav1.CreateOptions) (*v1.FlowSchema, error)
	Update(ctx context.Context, flowSchema *v1.FlowSchema, opts metav1.UpdateOptions) (*v1.FlowSchema, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, flowSchema *v1.FlowSchema, opts metav1.UpdateOptions) (*v1.FlowSchema, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.FlowSchema, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.FlowSchemaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.FlowSchema, err error)
	Apply(ctx context.Context, flowSchema *flowcontrolv1.FlowSchemaApplyConfiguration, opts metav1.ApplyOptions) (result *v1.FlowSchema, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, flowSchema *flowcontrolv1.FlowSchemaApplyConfiguration, opts metav1.ApplyOptions) (result *v1.FlowSchema, err error)
	FlowSchemaExpansion
}

// flowSchemas implements FlowSchemaInterface
type flowSchemas struct {
	*gentype.ClientWithListAndApply[*v1.FlowSchema, *v1.FlowSchemaList, *flowcontrolv1.FlowSchemaApplyConfiguration]
}

// newFlowSchemas returns a FlowSchemas
func newFlowSchemas(c *FlowcontrolV1Client) *flowSchemas {
	return &flowSchemas{
		gentype.NewClientWithListAndApply[*v1.FlowSchema, *v1.FlowSchemaList, *flowcontrolv1.FlowSchemaApplyConfiguration](
			"flowschemas",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.FlowSchema { return &v1.FlowSchema{} },
			func() *v1.FlowSchemaList { return &v1.FlowSchemaList{} }),
	}
}
