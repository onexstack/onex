// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	appsv1beta1 "github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// MinerLister helps list Miners.
// All objects returned here must be treated as read-only.
type MinerLister interface {
	// List lists all Miners in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1beta1.Miner, err error)
	// Miners returns an object that can list and get Miners.
	Miners(namespace string) MinerNamespaceLister
	MinerListerExpansion
}

// minerLister implements the MinerLister interface.
type minerLister struct {
	listers.ResourceIndexer[*appsv1beta1.Miner]
}

// NewMinerLister returns a new MinerLister.
func NewMinerLister(indexer cache.Indexer) MinerLister {
	return &minerLister{listers.New[*appsv1beta1.Miner](indexer, appsv1beta1.Resource("miner"))}
}

// Miners returns an object that can list and get Miners.
func (s *minerLister) Miners(namespace string) MinerNamespaceLister {
	return minerNamespaceLister{listers.NewNamespaced[*appsv1beta1.Miner](s.ResourceIndexer, namespace)}
}

// MinerNamespaceLister helps list and get Miners.
// All objects returned here must be treated as read-only.
type MinerNamespaceLister interface {
	// List lists all Miners in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1beta1.Miner, err error)
	// Get retrieves the Miner from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*appsv1beta1.Miner, error)
	MinerNamespaceListerExpansion
}

// minerNamespaceLister implements the MinerNamespaceLister
// interface.
type minerNamespaceLister struct {
	listers.ResourceIndexer[*appsv1beta1.Miner]
}
