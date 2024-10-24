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

// LimitRangeInformer provides access to a shared informer and lister for
// LimitRanges.
type LimitRangeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1.LimitRangeLister
}

type limitRangeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLimitRangeInformer constructs a new informer for LimitRange type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLimitRangeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLimitRangeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLimitRangeInformer constructs a new informer for LimitRange type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLimitRangeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().LimitRanges(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().LimitRanges(namespace).Watch(context.TODO(), options)
			},
		},
		&apicorev1.LimitRange{},
		resyncPeriod,
		indexers,
	)
}

func (f *limitRangeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLimitRangeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *limitRangeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apicorev1.LimitRange{}, f.defaultInformer)
}

func (f *limitRangeInformer) Lister() corev1.LimitRangeLister {
	return corev1.NewLimitRangeLister(f.Informer().GetIndexer())
}
