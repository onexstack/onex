// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package config

import (
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/cluster"

	"github.com/onexstack/onex/internal/controller/miner/apis/config"
)

// Config is the main context object for the controller.
type Config struct {
	ComponentConfig *config.MinerControllerConfiguration

	// the rest config for the master
	Kubeconfig *restclient.Config

	// Kubernetes clientset used to create miner pods.
	ProviderClient kubernetes.Interface

	ProviderCluster cluster.Cluster
}

// CompletedConfig same as Config, just to swap private object.
type CompletedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() *CompletedConfig {
	return &CompletedConfig{c}
}
