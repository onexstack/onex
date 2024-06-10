//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	coordination "github.com/superproj/onex/pkg/apis/coordination"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Lease)(nil), (*coordination.Lease)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Lease_To_coordination_Lease(a.(*Lease), b.(*coordination.Lease), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordination.Lease)(nil), (*Lease)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coordination_Lease_To_v1_Lease(a.(*coordination.Lease), b.(*Lease), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LeaseList)(nil), (*coordination.LeaseList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LeaseList_To_coordination_LeaseList(a.(*LeaseList), b.(*coordination.LeaseList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordination.LeaseList)(nil), (*LeaseList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coordination_LeaseList_To_v1_LeaseList(a.(*coordination.LeaseList), b.(*LeaseList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LeaseSpec)(nil), (*coordination.LeaseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LeaseSpec_To_coordination_LeaseSpec(a.(*LeaseSpec), b.(*coordination.LeaseSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*coordination.LeaseSpec)(nil), (*LeaseSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_coordination_LeaseSpec_To_v1_LeaseSpec(a.(*coordination.LeaseSpec), b.(*LeaseSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Lease_To_coordination_Lease(in *Lease, out *coordination.Lease, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_LeaseSpec_To_coordination_LeaseSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Lease_To_coordination_Lease is an autogenerated conversion function.
func Convert_v1_Lease_To_coordination_Lease(in *Lease, out *coordination.Lease, s conversion.Scope) error {
	return autoConvert_v1_Lease_To_coordination_Lease(in, out, s)
}

func autoConvert_coordination_Lease_To_v1_Lease(in *coordination.Lease, out *Lease, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_coordination_LeaseSpec_To_v1_LeaseSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_coordination_Lease_To_v1_Lease is an autogenerated conversion function.
func Convert_coordination_Lease_To_v1_Lease(in *coordination.Lease, out *Lease, s conversion.Scope) error {
	return autoConvert_coordination_Lease_To_v1_Lease(in, out, s)
}

func autoConvert_v1_LeaseList_To_coordination_LeaseList(in *LeaseList, out *coordination.LeaseList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]coordination.Lease)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_LeaseList_To_coordination_LeaseList is an autogenerated conversion function.
func Convert_v1_LeaseList_To_coordination_LeaseList(in *LeaseList, out *coordination.LeaseList, s conversion.Scope) error {
	return autoConvert_v1_LeaseList_To_coordination_LeaseList(in, out, s)
}

func autoConvert_coordination_LeaseList_To_v1_LeaseList(in *coordination.LeaseList, out *LeaseList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Lease)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_coordination_LeaseList_To_v1_LeaseList is an autogenerated conversion function.
func Convert_coordination_LeaseList_To_v1_LeaseList(in *coordination.LeaseList, out *LeaseList, s conversion.Scope) error {
	return autoConvert_coordination_LeaseList_To_v1_LeaseList(in, out, s)
}

func autoConvert_v1_LeaseSpec_To_coordination_LeaseSpec(in *LeaseSpec, out *coordination.LeaseSpec, s conversion.Scope) error {
	out.HolderIdentity = (*string)(unsafe.Pointer(in.HolderIdentity))
	out.LeaseDurationSeconds = (*int32)(unsafe.Pointer(in.LeaseDurationSeconds))
	out.AcquireTime = (*metav1.MicroTime)(unsafe.Pointer(in.AcquireTime))
	out.RenewTime = (*metav1.MicroTime)(unsafe.Pointer(in.RenewTime))
	out.LeaseTransitions = (*int32)(unsafe.Pointer(in.LeaseTransitions))
	return nil
}

// Convert_v1_LeaseSpec_To_coordination_LeaseSpec is an autogenerated conversion function.
func Convert_v1_LeaseSpec_To_coordination_LeaseSpec(in *LeaseSpec, out *coordination.LeaseSpec, s conversion.Scope) error {
	return autoConvert_v1_LeaseSpec_To_coordination_LeaseSpec(in, out, s)
}

func autoConvert_coordination_LeaseSpec_To_v1_LeaseSpec(in *coordination.LeaseSpec, out *LeaseSpec, s conversion.Scope) error {
	out.HolderIdentity = (*string)(unsafe.Pointer(in.HolderIdentity))
	out.LeaseDurationSeconds = (*int32)(unsafe.Pointer(in.LeaseDurationSeconds))
	out.AcquireTime = (*metav1.MicroTime)(unsafe.Pointer(in.AcquireTime))
	out.RenewTime = (*metav1.MicroTime)(unsafe.Pointer(in.RenewTime))
	out.LeaseTransitions = (*int32)(unsafe.Pointer(in.LeaseTransitions))
	return nil
}

// Convert_coordination_LeaseSpec_To_v1_LeaseSpec is an autogenerated conversion function.
func Convert_coordination_LeaseSpec_To_v1_LeaseSpec(in *coordination.LeaseSpec, out *LeaseSpec, s conversion.Scope) error {
	return autoConvert_coordination_LeaseSpec_To_v1_LeaseSpec(in, out, s)
}
