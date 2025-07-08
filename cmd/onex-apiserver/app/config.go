// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package app

import (
	apiextensionsapiserver "k8s.io/apiextensions-apiserver/pkg/apiserver"
	"k8s.io/apiserver/pkg/util/webhook"
	aggregatorapiserver "k8s.io/kube-aggregator/pkg/apiserver"

	"github.com/onexstack/onex/cmd/onex-apiserver/app/options"
	"github.com/onexstack/onex/internal/controlplane"
	"github.com/onexstack/onex/internal/controlplane/apiserver"
)

type Config struct {
	Options options.CompletedOptions

	Aggregator    *aggregatorapiserver.Config
	KubeAPIs      *controlplane.Config
	ApiExtensions *apiextensionsapiserver.Config

	ExtraConfig
}

type ExtraConfig struct{}

type completedConfig struct {
	Options options.CompletedOptions

	Aggregator    aggregatorapiserver.CompletedConfig
	KubeAPIs      controlplane.CompletedConfig
	ApiExtensions apiextensionsapiserver.CompletedConfig

	ExtraConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

func (c *Config) Complete() (CompletedConfig, error) {
	return CompletedConfig{&completedConfig{
		Options: c.Options,

		Aggregator:    c.Aggregator.Complete(),
		KubeAPIs:      c.KubeAPIs.Complete(),
		ApiExtensions: c.ApiExtensions.Complete(),

		ExtraConfig: c.ExtraConfig,
	}}, nil
}

// NewConfig creates all the resources for running kube-apiserver, but runs none of them.
func NewConfig(opts options.CompletedOptions) (*Config, error) {
	c := &Config{
		Options: opts,
	}

	controlPlane, serviceResolver, err := CreateOneXAPIServerConfig(opts)
	if err != nil {
		return nil, err
	}
	c.KubeAPIs = controlPlane

	apiExtensions, err := apiserver.CreateAPIExtensionsConfig(
		controlPlane.Generic.Config,
		controlPlane.Extra.InternalVersionedInformers,
		nil,
		opts.CompletedOptions,
		3,
		serviceResolver,
		webhook.NewDefaultAuthenticationInfoResolverWrapper(
			controlPlane.Extra.ProxyTransport,
			controlPlane.Generic.EgressSelector,
			controlPlane.Generic.LoopbackClientConfig,
			controlPlane.Generic.TracerProvider,
		),
	)
	if err != nil {
		return nil, err
	}
	c.ApiExtensions = apiExtensions

	aggregator, err := createAggregatorConfig(
		controlPlane.Generic.Config,
		opts.CompletedOptions,
		controlPlane.Extra.InternalVersionedInformers,
		serviceResolver,
		controlPlane.Extra.ProxyTransport,
		controlPlane.Extra.PeerProxy,
		nil,
	)
	if err != nil {
		return nil, err
	}
	c.Aggregator = aggregator

	return c, nil
}
