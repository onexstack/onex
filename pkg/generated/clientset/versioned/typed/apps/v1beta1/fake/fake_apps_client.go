// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/onexstack/onex/pkg/generated/clientset/versioned/typed/apps/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAppsV1beta1 struct {
	*testing.Fake
}

func (c *FakeAppsV1beta1) Chains(namespace string) v1beta1.ChainInterface {
	return &FakeChains{c, namespace}
}

func (c *FakeAppsV1beta1) ChargeRequests(namespace string) v1beta1.ChargeRequestInterface {
	return &FakeChargeRequests{c, namespace}
}

func (c *FakeAppsV1beta1) Miners(namespace string) v1beta1.MinerInterface {
	return &FakeMiners{c, namespace}
}

func (c *FakeAppsV1beta1) MinerSets(namespace string) v1beta1.MinerSetInterface {
	return &FakeMinerSets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAppsV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
