// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package config

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	genericconfig "github.com/onexstack/onex/pkg/config"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlockchainControllerConfiguration configures a scheduler.
type BlockchainControllerConfiguration struct {
	// TypeMeta contains the API version and kind.
	metav1.TypeMeta

	// Generic holds configuration for a generic controller-manager
	Generic genericconfig.GenericControllerManagerConfiguration

	// DryRun tells if the dry run mode is enabled, do not create an actual miner pod,
	// but directly set the miner status to Running.
	// If DryRun is set to true, the DryRun mode will be prioritized.
	// +optional
	DryRun bool

	// FeatureGates is a map of feature names to bools that enable or disable alpha/experimental features.
	FeatureGates map[string]bool

	// Path to miner provider kubeconfig file with authorization and master location information.
	// +optional
	ProviderKubeconfig string

	// Create miner pod in the cluster where miner controller is located.
	// +optional
	InCluster bool

	// Redis defines the configuration of redis client.
	Redis genericconfig.RedisConfiguration

	// MySQL defines the configuration of mysql client.
	MySQL genericconfig.MySQLConfiguration

	// Types specifies the configuration of the cloud mining machine.
	Types map[string]MinerProfile

	// ChainControllerConfiguration holds configuration for ChainController related features.
	ChainController ChainControllerConfiguration

	// Logs *logs.Options `json:"logs,omitempty"`
	// Metrics            *metrics.Options
	// Cloud options
	// Cloud *cloud.CloudOptions `json:"cloud,omitempty"`
}

type MinerProfile struct {
	CPU              resource.Quantity
	Memory           resource.Quantity
	MiningDifficulty int
}

type ChainControllerConfiguration struct {
	// Image specify the blockchain node image.
	Image string
}
