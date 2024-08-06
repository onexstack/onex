//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChainControllerConfiguration) DeepCopyInto(out *ChainControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChainControllerConfiguration.
func (in *ChainControllerConfiguration) DeepCopy() *ChainControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ChainControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneXControllerManagerConfiguration) DeepCopyInto(out *OneXControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.MySQL = in.MySQL
	in.Generic.DeepCopyInto(&out.Generic)
	in.GarbageCollectorController.DeepCopyInto(&out.GarbageCollectorController)
	out.ChainController = in.ChainController
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneXControllerManagerConfiguration.
func (in *OneXControllerManagerConfiguration) DeepCopy() *OneXControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(OneXControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OneXControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
