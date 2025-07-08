// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	genericconfig "github.com/onexstack/onex/pkg/config"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JobControllerConfiguration configures a scheduler.
type JobControllerConfiguration struct {
	// TypeMeta contains the API version and kind.
	metav1.TypeMeta `json:",inline"`

	// Generic holds configuration for a generic controller-manager
	Generic genericconfig.GenericControllerManagerConfiguration

	// ConcurrentCronJobSyncs is the number of cron job objects that are
	// allowed to sync concurrently. Larger number = more responsive jobs,
	// but more CPU (and network) load.
	ConcurrentCronJobSyncs int32

	// concurrentJobSyncs is the number of job objects that are
	// allowed to sync concurrently. Larger number = more responsive jobs,
	// but more CPU (and network) load.
	ConcurrentJobSyncs int32
}
