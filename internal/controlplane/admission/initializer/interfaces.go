// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	kubeinformers "k8s.io/client-go/informers"
	kubeclientset "k8s.io/client-go/kubernetes"
)

// WantsInternalInformerFactory defines a function which sets InformerFactory for admission plugins that need it.
type WantsInternalInformerFactory interface {
	admission.InitializationValidator
	SetInternalInformerFactory(kubeinformers.SharedInformerFactory)
}

// WantsInternalClientSet defines a function which sets external ClientSet for admission plugins that need it.
type WantsInternalClientSet interface {
	admission.InitializationValidator
	SetInternalClientSet(kubeclientset.Interface)
}
