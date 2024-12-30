// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package validation

import (
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/util/sets"

	genericknown "github.com/onexstack/onex/internal/pkg/known"
	known "github.com/onexstack/onex/internal/pkg/known/nightwatch"
	nwv1 "github.com/onexstack/onex/pkg/api/nightwatch/v1"
	"github.com/onexstack/onex/pkg/api/zerrors"
)

var availableScope = sets.New(
	known.LLMJobScope,
)

func validateJob(ctx *gin.Context, job *nwv1.Job) error {
	if job.Name == "" {
		return zerrors.ErrorInvalidParameter("job.name cannot be empty")
	}
	if job.Scope == "" {
		return zerrors.ErrorInvalidParameter("job.scope cannot be empty")
	}

	if !availableScope.Has(job.Scope) {
		return zerrors.ErrorInvalidParameter("invalid job.scope: %s", job.Scope)
	}

	if job.Watcher == "" {
		return zerrors.ErrorInvalidParameter("job.watcher cannot be empty")
	}

	job.UserID = ctx.GetString(genericknown.XUserID)

	return nil
}
