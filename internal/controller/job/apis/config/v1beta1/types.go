// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	genericconfigv1beta1 "github.com/onexstack/onex/pkg/config/v1beta1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// JobControllerConfiguration configures a scheduler.
type JobControllerConfiguration struct {
	// TypeMeta contains the API version and kind.
	metav1.TypeMeta `json:",inline"`

	// Generic holds configuration for a generic controller-manager
	Generic genericconfigv1beta1.GenericControllerManagerConfiguration `json:"generic,omitempty"`

	// ConcurrentCronJobSyncs is the number of cron job objects that are
	// allowed to sync concurrently. Larger number = more responsive jobs,
	// but more CPU (and network) load.
	ConcurrentCronJobSyncs int32 `json:"concurrentCronJobSyncs,omitempty"`

	// concurrentJobSyncs is the number of job objects that are
	// allowed to sync concurrently. Larger number = more responsive jobs,
	// but more CPU (and network) load.
	ConcurrentJobSyncs int32 `json:"concurrentJobSyncs,omitempty"`
}
