// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"
)

type pluginInitializer struct {
	informers informers.SharedInformerFactory
	client    clientset.Interface
	//authorizer        authorizer.Authorizer
	//featureGates      featuregate.FeatureGate
	stopCh <-chan struct{}
}

var _ admission.PluginInitializer = pluginInitializer{}

// New creates an instance of node admission plugins initializer.
func New(
	informers informers.SharedInformerFactory,
	client clientset.Interface,
) pluginInitializer {
	return pluginInitializer{
		informers: informers,
		client:    client,
	}
}

// Initialize checks the initialization interfaces implemented by a plugin
// and provide the appropriate initialization data.
func (i pluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsInternalInformerFactory); ok {
		wants.SetInternalInformerFactory(i.informers)
	}

	if wants, ok := plugin.(WantsInternalClientSet); ok {
		wants.SetInternalClientSet(i.client)
	}
}
