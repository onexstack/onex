// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	versioned "github.com/superproj/onex/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/superproj/onex/pkg/generated/informers/internalinterfaces"
	corev1 "github.com/superproj/onex/pkg/generated/listers/core/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecretInformer provides access to a shared informer and lister for
// Secrets.
type SecretInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1.SecretLister
}

type secretInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecretInformer constructs a new informer for Secret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecretInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecretInformer constructs a new informer for Secret type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecretInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Secrets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().Secrets(namespace).Watch(context.TODO(), options)
			},
		},
		&apicorev1.Secret{},
		resyncPeriod,
		indexers,
	)
}

func (f *secretInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecretInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secretInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apicorev1.Secret{}, f.defaultInformer)
}

func (f *secretInformer) Lister() corev1.SecretLister {
	return corev1.NewSecretLister(f.Informer().GetIndexer())
}
