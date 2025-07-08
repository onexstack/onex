// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	genericconfigv1beta1 "github.com/onexstack/onex/pkg/config/v1beta1"
)

const (
	// MinerControllerDefaultLockObjectNamespace defines default miner controller lock object namespace ("kube-system").
	MinerControllerDefaultLockObjectNamespace string = metav1.NamespaceSystem

	// MinerControllerDefaultLockObjectName defines default miner controller lock object name ("onex-miner-controller").
	MinerControllerDefaultLockObjectName = "onex-miner-controller"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlockchainControllerConfiguration configures a scheduler.
type BlockchainControllerConfiguration struct {
	// TypeMeta contains the API version and kind.
	metav1.TypeMeta `json:",inline"`

	// Generic holds configuration for a generic controller-manager
	Generic genericconfigv1beta1.GenericControllerManagerConfiguration `json:"generic,omitempty"`

	// DryRun tells if the dry run mode is enabled, do not create an actual miner pod,
	// but directly set the miner status to Running.
	// If DryRun is set to true, the DryRun mode will be prioritized.
	// +optional
	DryRun bool `json:"dryRun,omitempty"`

	// FeatureGates is a map of feature names to bools that enable or disable alpha/experimental features.
	FeatureGates map[string]bool `json:"featureGates,omitempty"`

	// Path to miner provider kubeconfig file with authorization and master location information.
	// +optional
	ProviderKubeconfig string `json:"providerKubeconfig,omitempty"`

	// Create miner pod in the cluster where miner controller is located.
	// +optional
	InCluster bool `json:"inCluster,omitempty"`

	// Redis defines the configuration of redis client.
	Redis genericconfigv1beta1.RedisConfiguration `json:"redis,omitempty"`

	// MySQL defines the configuration of mysql client.
	MySQL genericconfigv1beta1.MySQLConfiguration `json:"mysql,omitempty"`

	// Types specifies the configuration of the cloud mining machine.
	Types map[string]MinerProfile `json:"types,omitempty"`

	// ChainControllerConfiguration holds configuration for ChainController related features.
	ChainController ChainControllerConfiguration `json:"chainController,omitempty"`
}

type MinerProfile struct {
	CPU              resource.Quantity `json:"cpu,omitempty"`
	Memory           resource.Quantity `json:"memory,omitempty"`
	MiningDifficulty int               `json:"miningDifficulty,omitempty"`
}

type ChainControllerConfiguration struct {
	// Image specify the blockchain node image.
	Image string `json:"image,omitempty"`
}
