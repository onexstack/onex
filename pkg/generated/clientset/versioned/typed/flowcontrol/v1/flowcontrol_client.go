// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	http "net/http"

	scheme "github.com/onexstack/onex/pkg/generated/clientset/versioned/scheme"
	flowcontrolv1 "k8s.io/api/flowcontrol/v1"
	rest "k8s.io/client-go/rest"
)

type FlowcontrolV1Interface interface {
	RESTClient() rest.Interface
	FlowSchemasGetter
	PriorityLevelConfigurationsGetter
}

// FlowcontrolV1Client is used to interact with features provided by the flowcontrol.apiserver.k8s.io group.
type FlowcontrolV1Client struct {
	restClient rest.Interface
}

func (c *FlowcontrolV1Client) FlowSchemas() FlowSchemaInterface {
	return newFlowSchemas(c)
}

func (c *FlowcontrolV1Client) PriorityLevelConfigurations() PriorityLevelConfigurationInterface {
	return newPriorityLevelConfigurations(c)
}

// NewForConfig creates a new FlowcontrolV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*FlowcontrolV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new FlowcontrolV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*FlowcontrolV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &FlowcontrolV1Client{client}, nil
}

// NewForConfigOrDie creates a new FlowcontrolV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *FlowcontrolV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new FlowcontrolV1Client for the given RESTClient.
func New(c rest.Interface) *FlowcontrolV1Client {
	return &FlowcontrolV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := flowcontrolv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FlowcontrolV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
