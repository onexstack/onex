// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// EphemeralContainerApplyConfiguration represents a declarative configuration of the EphemeralContainer type for use
// with apply.
type EphemeralContainerApplyConfiguration struct {
	EphemeralContainerCommonApplyConfiguration `json:",inline"`
	TargetContainerName                        *string `json:"targetContainerName,omitempty"`
}

// EphemeralContainerApplyConfiguration constructs a declarative configuration of the EphemeralContainer type for use with
// apply.
func EphemeralContainer() *EphemeralContainerApplyConfiguration {
	return &EphemeralContainerApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithName(value string) *EphemeralContainerApplyConfiguration {
	b.Name = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithImage(value string) *EphemeralContainerApplyConfiguration {
	b.Image = &value
	return b
}

// WithCommand adds the given value to the Command field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Command field.
func (b *EphemeralContainerApplyConfiguration) WithCommand(values ...string) *EphemeralContainerApplyConfiguration {
	for i := range values {
		b.Command = append(b.Command, values[i])
	}
	return b
}

// WithArgs adds the given value to the Args field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Args field.
func (b *EphemeralContainerApplyConfiguration) WithArgs(values ...string) *EphemeralContainerApplyConfiguration {
	for i := range values {
		b.Args = append(b.Args, values[i])
	}
	return b
}

// WithWorkingDir sets the WorkingDir field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WorkingDir field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithWorkingDir(value string) *EphemeralContainerApplyConfiguration {
	b.WorkingDir = &value
	return b
}

// WithPorts adds the given value to the Ports field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Ports field.
func (b *EphemeralContainerApplyConfiguration) WithPorts(values ...*ContainerPortApplyConfiguration) *EphemeralContainerApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPorts")
		}
		b.Ports = append(b.Ports, *values[i])
	}
	return b
}

// WithEnvFrom adds the given value to the EnvFrom field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EnvFrom field.
func (b *EphemeralContainerApplyConfiguration) WithEnvFrom(values ...*EnvFromSourceApplyConfiguration) *EphemeralContainerApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEnvFrom")
		}
		b.EnvFrom = append(b.EnvFrom, *values[i])
	}
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *EphemeralContainerApplyConfiguration) WithEnv(values ...*EnvVarApplyConfiguration) *EphemeralContainerApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEnv")
		}
		b.Env = append(b.Env, *values[i])
	}
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithResources(value *ResourceRequirementsApplyConfiguration) *EphemeralContainerApplyConfiguration {
	b.Resources = value
	return b
}

// WithResizePolicy adds the given value to the ResizePolicy field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ResizePolicy field.
func (b *EphemeralContainerApplyConfiguration) WithResizePolicy(values ...*ContainerResizePolicyApplyConfiguration) *EphemeralContainerApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithResizePolicy")
		}
		b.ResizePolicy = append(b.ResizePolicy, *values[i])
	}
	return b
}

// WithRestartPolicy sets the RestartPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RestartPolicy field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithRestartPolicy(value corev1.ContainerRestartPolicy) *EphemeralContainerApplyConfiguration {
	b.RestartPolicy = &value
	return b
}

// WithVolumeMounts adds the given value to the VolumeMounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VolumeMounts field.
func (b *EphemeralContainerApplyConfiguration) WithVolumeMounts(values ...*VolumeMountApplyConfiguration) *EphemeralContainerApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumeMounts")
		}
		b.VolumeMounts = append(b.VolumeMounts, *values[i])
	}
	return b
}

// WithVolumeDevices adds the given value to the VolumeDevices field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the VolumeDevices field.
func (b *EphemeralContainerApplyConfiguration) WithVolumeDevices(values ...*VolumeDeviceApplyConfiguration) *EphemeralContainerApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithVolumeDevices")
		}
		b.VolumeDevices = append(b.VolumeDevices, *values[i])
	}
	return b
}

// WithLivenessProbe sets the LivenessProbe field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LivenessProbe field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithLivenessProbe(value *ProbeApplyConfiguration) *EphemeralContainerApplyConfiguration {
	b.LivenessProbe = value
	return b
}

// WithReadinessProbe sets the ReadinessProbe field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadinessProbe field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithReadinessProbe(value *ProbeApplyConfiguration) *EphemeralContainerApplyConfiguration {
	b.ReadinessProbe = value
	return b
}

// WithStartupProbe sets the StartupProbe field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartupProbe field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithStartupProbe(value *ProbeApplyConfiguration) *EphemeralContainerApplyConfiguration {
	b.StartupProbe = value
	return b
}

// WithLifecycle sets the Lifecycle field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Lifecycle field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithLifecycle(value *LifecycleApplyConfiguration) *EphemeralContainerApplyConfiguration {
	b.Lifecycle = value
	return b
}

// WithTerminationMessagePath sets the TerminationMessagePath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TerminationMessagePath field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithTerminationMessagePath(value string) *EphemeralContainerApplyConfiguration {
	b.TerminationMessagePath = &value
	return b
}

// WithTerminationMessagePolicy sets the TerminationMessagePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TerminationMessagePolicy field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithTerminationMessagePolicy(value corev1.TerminationMessagePolicy) *EphemeralContainerApplyConfiguration {
	b.TerminationMessagePolicy = &value
	return b
}

// WithImagePullPolicy sets the ImagePullPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImagePullPolicy field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithImagePullPolicy(value corev1.PullPolicy) *EphemeralContainerApplyConfiguration {
	b.ImagePullPolicy = &value
	return b
}

// WithSecurityContext sets the SecurityContext field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecurityContext field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithSecurityContext(value *SecurityContextApplyConfiguration) *EphemeralContainerApplyConfiguration {
	b.SecurityContext = value
	return b
}

// WithStdin sets the Stdin field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Stdin field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithStdin(value bool) *EphemeralContainerApplyConfiguration {
	b.Stdin = &value
	return b
}

// WithStdinOnce sets the StdinOnce field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StdinOnce field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithStdinOnce(value bool) *EphemeralContainerApplyConfiguration {
	b.StdinOnce = &value
	return b
}

// WithTTY sets the TTY field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TTY field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithTTY(value bool) *EphemeralContainerApplyConfiguration {
	b.TTY = &value
	return b
}

// WithTargetContainerName sets the TargetContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetContainerName field is set to the value of the last call.
func (b *EphemeralContainerApplyConfiguration) WithTargetContainerName(value string) *EphemeralContainerApplyConfiguration {
	b.TargetContainerName = &value
	return b
}
