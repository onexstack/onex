// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package validation

import (
	genericvalidation "github.com/onexstack/onex/pkg/config/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	componentbasevalidation "k8s.io/component-base/config/validation"

	"github.com/onexstack/onex/internal/controller/job/apis/config"
)

// Validate ensures validation of the JobControllerConfiguration struct.
func Validate(cc *config.JobControllerConfiguration) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, componentbasevalidation.ValidateLeaderElectionConfiguration(&cc.Generic.LeaderElection, field.NewPath("generic", "leaderElection"))...)
	allErrs = append(allErrs, genericvalidation.ValidateGenericControllerManagerConfiguration(&cc.Generic, field.NewPath("generic"))...)

	return allErrs
}
