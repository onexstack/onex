// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	applyconfigurationscorev1 "github.com/onexstack/onex/pkg/generated/applyconfigurations/core/v1"
	scheme "github.com/onexstack/onex/pkg/generated/clientset/versioned/scheme"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ServiceAccountsGetter has a method to return a ServiceAccountInterface.
// A group's client should implement this interface.
type ServiceAccountsGetter interface {
	ServiceAccounts(namespace string) ServiceAccountInterface
}

// ServiceAccountInterface has methods to work with ServiceAccount resources.
type ServiceAccountInterface interface {
	Create(ctx context.Context, serviceAccount *corev1.ServiceAccount, opts metav1.CreateOptions) (*corev1.ServiceAccount, error)
	Update(ctx context.Context, serviceAccount *corev1.ServiceAccount, opts metav1.UpdateOptions) (*corev1.ServiceAccount, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.ServiceAccount, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceAccountList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.ServiceAccount, err error)
	Apply(ctx context.Context, serviceAccount *applyconfigurationscorev1.ServiceAccountApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.ServiceAccount, err error)
	CreateToken(ctx context.Context, serviceAccountName string, tokenRequest *authenticationv1.TokenRequest, opts metav1.CreateOptions) (*authenticationv1.TokenRequest, error)

	ServiceAccountExpansion
}

// serviceAccounts implements ServiceAccountInterface
type serviceAccounts struct {
	*gentype.ClientWithListAndApply[*corev1.ServiceAccount, *corev1.ServiceAccountList, *applyconfigurationscorev1.ServiceAccountApplyConfiguration]
}

// newServiceAccounts returns a ServiceAccounts
func newServiceAccounts(c *CoreV1Client, namespace string) *serviceAccounts {
	return &serviceAccounts{
		gentype.NewClientWithListAndApply[*corev1.ServiceAccount, *corev1.ServiceAccountList, *applyconfigurationscorev1.ServiceAccountApplyConfiguration](
			"serviceaccounts",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.ServiceAccount { return &corev1.ServiceAccount{} },
			func() *corev1.ServiceAccountList { return &corev1.ServiceAccountList{} },
			gentype.PrefersProtobuf[*corev1.ServiceAccount](),
		),
	}
}

// CreateToken takes the representation of a tokenRequest and creates it.  Returns the server's representation of the tokenRequest, and an error, if there is any.
func (c *serviceAccounts) CreateToken(ctx context.Context, serviceAccountName string, tokenRequest *authenticationv1.TokenRequest, opts metav1.CreateOptions) (result *authenticationv1.TokenRequest, err error) {
	result = &authenticationv1.TokenRequest{}
	err = c.GetClient().Post().
		UseProtobufAsDefault().
		Namespace(c.GetNamespace()).
		Resource("serviceaccounts").
		Name(serviceAccountName).
		SubResource("token").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tokenRequest).
		Do(ctx).
		Into(result)
	return
}
