// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodTemplateLister helps list PodTemplates.
// All objects returned here must be treated as read-only.
type PodTemplateLister interface {
	// List lists all PodTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodTemplate, err error)
	// PodTemplates returns an object that can list and get PodTemplates.
	PodTemplates(namespace string) PodTemplateNamespaceLister
	PodTemplateListerExpansion
}

// podTemplateLister implements the PodTemplateLister interface.
type podTemplateLister struct {
	indexer cache.Indexer
}

// NewPodTemplateLister returns a new PodTemplateLister.
func NewPodTemplateLister(indexer cache.Indexer) PodTemplateLister {
	return &podTemplateLister{indexer: indexer}
}

// List lists all PodTemplates in the indexer.
func (s *podTemplateLister) List(selector labels.Selector) (ret []*v1.PodTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodTemplate))
	})
	return ret, err
}

// PodTemplates returns an object that can list and get PodTemplates.
func (s *podTemplateLister) PodTemplates(namespace string) PodTemplateNamespaceLister {
	return podTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodTemplateNamespaceLister helps list and get PodTemplates.
// All objects returned here must be treated as read-only.
type PodTemplateNamespaceLister interface {
	// List lists all PodTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PodTemplate, err error)
	// Get retrieves the PodTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PodTemplate, error)
	PodTemplateNamespaceListerExpansion
}

// podTemplateNamespaceLister implements the PodTemplateNamespaceLister
// interface.
type podTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodTemplates in the indexer for a given namespace.
func (s podTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1.PodTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodTemplate))
	})
	return ret, err
}

// Get retrieves the PodTemplate from the indexer for a given namespace and name.
func (s podTemplateNamespaceLister) Get(name string) (*v1.PodTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("podtemplate"), name)
	}
	return obj.(*v1.PodTemplate), nil
}
