// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package validation

import (
	genericvalidation "github.com/onexstack/onex/pkg/config/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	componentbasevalidation "k8s.io/component-base/config/validation"

	"github.com/onexstack/onex/internal/controller/blockchain/apis/config"
)

// Validate ensures validation of the BlockchainControllerConfiguration struct.
func Validate(cc *config.BlockchainControllerConfiguration) field.ErrorList {
	allErrs := field.ErrorList{}
	newPath := field.NewPath("BlockchainControllerConfiguration")

	effectiveFeatures := utilfeature.DefaultFeatureGate.DeepCopy()
	if err := effectiveFeatures.SetFromMap(cc.FeatureGates); err != nil {
		allErrs = append(allErrs, field.Invalid(newPath.Child("featureGates"), cc.FeatureGates, err.Error()))
	}
	allErrs = append(allErrs, componentbasevalidation.ValidateLeaderElectionConfiguration(&cc.Generic.LeaderElection, field.NewPath("generic", "leaderElection"))...)
	allErrs = append(allErrs, genericvalidation.ValidateMySQLConfiguration(&cc.MySQL, field.NewPath("mysql"))...)
	allErrs = append(allErrs, genericvalidation.ValidateGenericControllerManagerConfiguration(&cc.Generic, field.NewPath("generic"))...)

	return allErrs
}
