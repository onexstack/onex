// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/conversion"

	"github.com/superproj/onex/pkg/config"
)

// Important! The public back-and-forth conversion functions for the types in this generic
// package with ComponentConfig types need to be manually exposed like this in order for
// other packages that reference this package to be able to call these conversion functions
// in an autogenerated manner.
// TODO: Fix the bug in conversion-gen so it automatically discovers these Convert_* functions
// in autogenerated code as well.

func Convert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in *GenericControllerManagerConfiguration, out *config.GenericControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(in, out, s)

}

func Convert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(in *config.GenericControllerManagerConfiguration, out *GenericControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(in, out, s)
}

func Convert_v1beta1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in *GarbageCollectorControllerConfiguration, out *config.GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_GarbageCollectorControllerConfiguration_To_config_GarbageCollectorControllerConfiguration(in, out, s)
}

func Convert_config_GarbageCollectorControllerConfiguration_To_v1beta1_GarbageCollectorControllerConfiguration(in *config.GarbageCollectorControllerConfiguration, out *GarbageCollectorControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_GarbageCollectorControllerConfiguration_To_v1beta1_GarbageCollectorControllerConfiguration(in, out, s)
}

func Convert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(in *MySQLConfiguration, out *config.MySQLConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(in, out, s)
}

func Convert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(in *config.MySQLConfiguration, out *MySQLConfiguration, s conversion.Scope) error {
	return autoConvert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(in, out, s)
}
func Convert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(in *RedisConfiguration, out *config.RedisConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(in, out, s)
}
func Convert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(in *config.RedisConfiguration, out *RedisConfiguration, s conversion.Scope) error {
	return autoConvert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(in, out, s)
}
