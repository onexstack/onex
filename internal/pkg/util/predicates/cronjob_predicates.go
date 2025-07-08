// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

// Package predicates implements predicate utilities.
package predicates

import (
	"github.com/go-logr/logr"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	batchv1beta1 "github.com/onexstack/onex/pkg/apis/batch/v1beta1"
)

func CronJobNotSuspend(logger logr.Logger) predicate.Funcs {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			return processIfNotPaused(logger.WithValues("predicate", "ResourceNotPaused", "eventType", "update"), e.ObjectNew)
		},
		CreateFunc: func(e event.CreateEvent) bool {
			return processIfNotPaused(logger.WithValues("predicate", "ResourceNotPaused", "eventType", "create"), e.Object)
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return processIfNotPaused(logger.WithValues("predicate", "ResourceNotPaused", "eventType", "delete"), e.Object)
		},
		GenericFunc: func(e event.GenericEvent) bool {
			return processIfNotPaused(logger.WithValues("predicate", "ResourceNotPaused", "eventType", "generic"), e.Object)
		},
	}
}

func processIfNotSuspend(logger logr.Logger, obj client.Object) bool {
	cronJob, ok := obj.(*batchv1beta1.CronJob)
	if !ok {
		return false
	}

	if ptr.Deref(cronJob.Spec.Suspend, false) {
		logger.V(6).Info("Resource is not suspend, will attempt to map resource")
		return true
	}

	logger.V(4).Info("Resource is suspend, will not attempt to map resource")
	return false
}
