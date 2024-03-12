//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by defaulter-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1.Pod{}, func(obj any) { SetObjectDefaults_Pod(obj.(*v1.Pod)) })
	scheme.AddTypeDefaultingFunc(&v1.PodList{}, func(obj any) { SetObjectDefaults_PodList(obj.(*v1.PodList)) })
	scheme.AddTypeDefaultingFunc(&v1.PodTemplate{}, func(obj any) { SetObjectDefaults_PodTemplate(obj.(*v1.PodTemplate)) })
	scheme.AddTypeDefaultingFunc(&v1.PodTemplateList{}, func(obj any) { SetObjectDefaults_PodTemplateList(obj.(*v1.PodTemplateList)) })
	scheme.AddTypeDefaultingFunc(&v1.ReplicationController{}, func(obj any) { SetObjectDefaults_ReplicationController(obj.(*v1.ReplicationController)) })
	scheme.AddTypeDefaultingFunc(&v1.ReplicationControllerList{}, func(obj any) {
		SetObjectDefaults_ReplicationControllerList(obj.(*v1.ReplicationControllerList))
	})
	scheme.AddTypeDefaultingFunc(&v1.Service{}, func(obj any) { SetObjectDefaults_Service(obj.(*v1.Service)) })
	scheme.AddTypeDefaultingFunc(&v1.ServiceList{}, func(obj any) { SetObjectDefaults_ServiceList(obj.(*v1.ServiceList)) })
	return nil
}

func SetObjectDefaults_Pod(in *v1.Pod) {
	for i := range in.Spec.InitContainers {
		a := &in.Spec.InitContainers[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		if a.LivenessProbe != nil {
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
	}
	for i := range in.Spec.Containers {
		a := &in.Spec.Containers[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		if a.LivenessProbe != nil {
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
	}
	for i := range in.Spec.EphemeralContainers {
		a := &in.Spec.EphemeralContainers[i]
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
	}
}

func SetObjectDefaults_PodList(in *v1.PodList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Pod(a)
	}
}

func SetObjectDefaults_PodTemplate(in *v1.PodTemplate) {
	for i := range in.Template.Spec.InitContainers {
		a := &in.Template.Spec.InitContainers[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		if a.LivenessProbe != nil {
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
	}
	for i := range in.Template.Spec.Containers {
		a := &in.Template.Spec.Containers[i]
		for j := range a.Ports {
			b := &a.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		if a.LivenessProbe != nil {
			if a.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.ReadinessProbe != nil {
			if a.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.StartupProbe != nil {
			if a.StartupProbe.ProbeHandler.GRPC != nil {
				if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
	}
	for i := range in.Template.Spec.EphemeralContainers {
		a := &in.Template.Spec.EphemeralContainers[i]
		for j := range a.EphemeralContainerCommon.Ports {
			b := &a.EphemeralContainerCommon.Ports[j]
			if b.Protocol == "" {
				b.Protocol = "TCP"
			}
		}
		if a.EphemeralContainerCommon.LivenessProbe != nil {
			if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.ReadinessProbe != nil {
			if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
		if a.EphemeralContainerCommon.StartupProbe != nil {
			if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
					var ptrVar1 string = ""
					a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
				}
			}
		}
	}
}

func SetObjectDefaults_PodTemplateList(in *v1.PodTemplateList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PodTemplate(a)
	}
}

func SetObjectDefaults_ReplicationController(in *v1.ReplicationController) {
	if in.Spec.Template != nil {
		for i := range in.Spec.Template.Spec.InitContainers {
			a := &in.Spec.Template.Spec.InitContainers[i]
			for j := range a.Ports {
				b := &a.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.LivenessProbe != nil {
				if a.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.ReadinessProbe != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.StartupProbe != nil {
				if a.StartupProbe.ProbeHandler.GRPC != nil {
					if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
		for i := range in.Spec.Template.Spec.Containers {
			a := &in.Spec.Template.Spec.Containers[i]
			for j := range a.Ports {
				b := &a.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.LivenessProbe != nil {
				if a.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.ReadinessProbe != nil {
				if a.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.StartupProbe != nil {
				if a.StartupProbe.ProbeHandler.GRPC != nil {
					if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
		for i := range in.Spec.Template.Spec.EphemeralContainers {
			a := &in.Spec.Template.Spec.EphemeralContainers[i]
			for j := range a.EphemeralContainerCommon.Ports {
				b := &a.EphemeralContainerCommon.Ports[j]
				if b.Protocol == "" {
					b.Protocol = "TCP"
				}
			}
			if a.EphemeralContainerCommon.LivenessProbe != nil {
				if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.EphemeralContainerCommon.ReadinessProbe != nil {
				if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
			if a.EphemeralContainerCommon.StartupProbe != nil {
				if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC != nil {
					if a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service == nil {
						var ptrVar1 string = ""
						a.EphemeralContainerCommon.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
					}
				}
			}
		}
	}
}

func SetObjectDefaults_ReplicationControllerList(in *v1.ReplicationControllerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ReplicationController(a)
	}
}

func SetObjectDefaults_Service(in *v1.Service) {
	for i := range in.Spec.Ports {
		a := &in.Spec.Ports[i]
		if a.Protocol == "" {
			a.Protocol = "TCP"
		}
	}
}

func SetObjectDefaults_ServiceList(in *v1.ServiceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Service(a)
	}
}
