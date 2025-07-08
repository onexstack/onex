/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/ptr"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

func SetDefaults_Job(obj *Job) {}

func SetDefaults_CronJob(obj *CronJob) {
	if obj.Spec.ConcurrencyPolicy == "" {
		obj.Spec.ConcurrencyPolicy = AllowConcurrent
	}
	if obj.Spec.Suspend == nil {
		obj.Spec.Suspend = ptr.To(false)
	}
	if obj.Spec.SuccessfulJobsHistoryLimit == nil {
		obj.Spec.SuccessfulJobsHistoryLimit = ptr.To[int32](3)
	}
	if obj.Spec.FailedJobsHistoryLimit == nil {
		obj.Spec.FailedJobsHistoryLimit = ptr.To[int32](1)
	}
}
