//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	config "github.com/onexstack/onex/internal/controller/blockchain/apis/config"
	configv1beta1 "github.com/onexstack/onex/pkg/config/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*BlockchainControllerConfiguration)(nil), (*config.BlockchainControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_BlockchainControllerConfiguration_To_config_BlockchainControllerConfiguration(a.(*BlockchainControllerConfiguration), b.(*config.BlockchainControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.BlockchainControllerConfiguration)(nil), (*BlockchainControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_BlockchainControllerConfiguration_To_v1beta1_BlockchainControllerConfiguration(a.(*config.BlockchainControllerConfiguration), b.(*BlockchainControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChainControllerConfiguration)(nil), (*config.ChainControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ChainControllerConfiguration_To_config_ChainControllerConfiguration(a.(*ChainControllerConfiguration), b.(*config.ChainControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ChainControllerConfiguration)(nil), (*ChainControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ChainControllerConfiguration_To_v1beta1_ChainControllerConfiguration(a.(*config.ChainControllerConfiguration), b.(*ChainControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MinerProfile)(nil), (*config.MinerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerProfile_To_config_MinerProfile(a.(*MinerProfile), b.(*config.MinerProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MinerProfile)(nil), (*MinerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MinerProfile_To_v1beta1_MinerProfile(a.(*config.MinerProfile), b.(*MinerProfile), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_BlockchainControllerConfiguration_To_config_BlockchainControllerConfiguration(in *BlockchainControllerConfiguration, out *config.BlockchainControllerConfiguration, s conversion.Scope) error {
	if err := configv1beta1.Convert_v1beta1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	out.DryRun = in.DryRun
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ProviderKubeconfig = in.ProviderKubeconfig
	out.InCluster = in.InCluster
	if err := configv1beta1.Convert_v1beta1_RedisConfiguration_To_config_RedisConfiguration(&in.Redis, &out.Redis, s); err != nil {
		return err
	}
	if err := configv1beta1.Convert_v1beta1_MySQLConfiguration_To_config_MySQLConfiguration(&in.MySQL, &out.MySQL, s); err != nil {
		return err
	}
	out.Types = *(*map[string]config.MinerProfile)(unsafe.Pointer(&in.Types))
	if err := Convert_v1beta1_ChainControllerConfiguration_To_config_ChainControllerConfiguration(&in.ChainController, &out.ChainController, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_BlockchainControllerConfiguration_To_config_BlockchainControllerConfiguration is an autogenerated conversion function.
func Convert_v1beta1_BlockchainControllerConfiguration_To_config_BlockchainControllerConfiguration(in *BlockchainControllerConfiguration, out *config.BlockchainControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_BlockchainControllerConfiguration_To_config_BlockchainControllerConfiguration(in, out, s)
}

func autoConvert_config_BlockchainControllerConfiguration_To_v1beta1_BlockchainControllerConfiguration(in *config.BlockchainControllerConfiguration, out *BlockchainControllerConfiguration, s conversion.Scope) error {
	if err := configv1beta1.Convert_config_GenericControllerManagerConfiguration_To_v1beta1_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	out.DryRun = in.DryRun
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.ProviderKubeconfig = in.ProviderKubeconfig
	out.InCluster = in.InCluster
	if err := configv1beta1.Convert_config_RedisConfiguration_To_v1beta1_RedisConfiguration(&in.Redis, &out.Redis, s); err != nil {
		return err
	}
	if err := configv1beta1.Convert_config_MySQLConfiguration_To_v1beta1_MySQLConfiguration(&in.MySQL, &out.MySQL, s); err != nil {
		return err
	}
	out.Types = *(*map[string]MinerProfile)(unsafe.Pointer(&in.Types))
	if err := Convert_config_ChainControllerConfiguration_To_v1beta1_ChainControllerConfiguration(&in.ChainController, &out.ChainController, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_BlockchainControllerConfiguration_To_v1beta1_BlockchainControllerConfiguration is an autogenerated conversion function.
func Convert_config_BlockchainControllerConfiguration_To_v1beta1_BlockchainControllerConfiguration(in *config.BlockchainControllerConfiguration, out *BlockchainControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_BlockchainControllerConfiguration_To_v1beta1_BlockchainControllerConfiguration(in, out, s)
}

func autoConvert_v1beta1_ChainControllerConfiguration_To_config_ChainControllerConfiguration(in *ChainControllerConfiguration, out *config.ChainControllerConfiguration, s conversion.Scope) error {
	out.Image = in.Image
	return nil
}

// Convert_v1beta1_ChainControllerConfiguration_To_config_ChainControllerConfiguration is an autogenerated conversion function.
func Convert_v1beta1_ChainControllerConfiguration_To_config_ChainControllerConfiguration(in *ChainControllerConfiguration, out *config.ChainControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_ChainControllerConfiguration_To_config_ChainControllerConfiguration(in, out, s)
}

func autoConvert_config_ChainControllerConfiguration_To_v1beta1_ChainControllerConfiguration(in *config.ChainControllerConfiguration, out *ChainControllerConfiguration, s conversion.Scope) error {
	out.Image = in.Image
	return nil
}

// Convert_config_ChainControllerConfiguration_To_v1beta1_ChainControllerConfiguration is an autogenerated conversion function.
func Convert_config_ChainControllerConfiguration_To_v1beta1_ChainControllerConfiguration(in *config.ChainControllerConfiguration, out *ChainControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_ChainControllerConfiguration_To_v1beta1_ChainControllerConfiguration(in, out, s)
}

func autoConvert_v1beta1_MinerProfile_To_config_MinerProfile(in *MinerProfile, out *config.MinerProfile, s conversion.Scope) error {
	out.CPU = in.CPU
	out.Memory = in.Memory
	out.MiningDifficulty = in.MiningDifficulty
	return nil
}

// Convert_v1beta1_MinerProfile_To_config_MinerProfile is an autogenerated conversion function.
func Convert_v1beta1_MinerProfile_To_config_MinerProfile(in *MinerProfile, out *config.MinerProfile, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerProfile_To_config_MinerProfile(in, out, s)
}

func autoConvert_config_MinerProfile_To_v1beta1_MinerProfile(in *config.MinerProfile, out *MinerProfile, s conversion.Scope) error {
	out.CPU = in.CPU
	out.Memory = in.Memory
	out.MiningDifficulty = in.MiningDifficulty
	return nil
}

// Convert_config_MinerProfile_To_v1beta1_MinerProfile is an autogenerated conversion function.
func Convert_config_MinerProfile_To_v1beta1_MinerProfile(in *config.MinerProfile, out *MinerProfile, s conversion.Scope) error {
	return autoConvert_config_MinerProfile_To_v1beta1_MinerProfile(in, out, s)
}
