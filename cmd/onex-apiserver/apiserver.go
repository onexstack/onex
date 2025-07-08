// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// APIServer is the main API server and master for the onex.
// It is responsible for serving the onex management API.
package main

import (
	"context"
	"os"

	_ "go.uber.org/automaxprocs/maxprocs"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"
	_ "k8s.io/component-base/logs/json/register"          // for JSON log format registration
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration
	"k8s.io/klog/v2"

	"github.com/onexstack/onex/cmd/onex-apiserver/app"
	"github.com/onexstack/onex/internal/apiserver/admission/plugin/minerset"

	"github.com/onexstack/onex/internal/apiserver/admission/initializer"
	appsrest "github.com/onexstack/onex/internal/apiserver/registry/apps/rest"
	batchrest "github.com/onexstack/onex/internal/apiserver/registry/batch/rest"
	"github.com/onexstack/onex/internal/pkg/config/minerprofile"
	appsv1beta1 "github.com/onexstack/onex/pkg/apis/apps/v1beta1"
	batchv1beta1 "github.com/onexstack/onex/pkg/apis/batch/v1beta1"
	"github.com/onexstack/onex/pkg/generated/clientset/versioned"
	"github.com/onexstack/onex/pkg/generated/informers"
	generatedopenapi "github.com/onexstack/onex/pkg/generated/openapi"
)

func main() {
	var informerFactory informers.SharedInformerFactory

	// Please note that the following WithOptions are all required.
	command := app.NewAPIServerCommand(
		// Add custom etcd options.
		app.WithEtcdOptions("/registry/onex.io", appsv1beta1.SchemeGroupVersion, batchv1beta1.SchemeGroupVersion),
		// Add custom resource storage.
		app.WithRESTStorageProviders(appsrest.RESTStorageProvider{}, batchrest.RESTStorageProvider{}),
		// Add custom dns address.
		app.WithAlternateDNS("onex.io"),
		// Add custom admission plugins.
		app.WithAdmissionPlugin(minerset.PluginName, minerset.Register),
		// Add custom admission plugins initializer.
		app.WithGetOpenAPIDefinitions(generatedopenapi.GetOpenAPIDefinitions),
		app.WithAdmissionInitializers(func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
			client, err := versioned.NewForConfig(c.LoopbackClientConfig)
			if err != nil {
				return nil, err
			}
			informerFactory = informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
			// NOTICE: As we create a shared informer, we need to start it later.
			// We can usually start it by adding a PostStartHook.
			return []admission.PluginInitializer{initializer.New(informerFactory, client)}, nil
		}),
		app.WithPostStartHook(
			"start-external-informers",
			func(ctx genericapiserver.PostStartHookContext) error {
				if informerFactory != nil {
					informerFactory.Start(ctx.Done())
				}
				return nil
			}),
		app.WithPostStartHook(
			"initialize-instance-config-client",
			func(ctx genericapiserver.PostStartHookContext) error {
				client, err := versioned.NewForConfig(ctx.LoopbackClientConfig)
				if err != nil {
					return err
				}

				if err := minerprofile.Init(context.Background(), client); err != nil {
					// When returning 'NotFound' error, we should not report an error, otherwise we can not
					// create 'MinerTypesConfigMapName' configmap via onex-apiserver
					if apierrors.IsNotFound(err) {
						return nil
					}

					klog.ErrorS(err, "Failed to init miner type cache")
					return err
				}

				return nil
			},
		),
	)

	code := cli.Run(command)
	os.Exit(code)
}
