// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1beta1 "github.com/superproj/onex/pkg/apis/apps/v1beta1"
	apiextensionsv1 "github.com/superproj/onex/pkg/generated/applyconfigurations/apiextensions/v1"
	appsv1beta1 "github.com/superproj/onex/pkg/generated/applyconfigurations/apps/v1beta1"
	applyconfigurationsautoscalingv1 "github.com/superproj/onex/pkg/generated/applyconfigurations/autoscaling/v1"
	applyconfigurationscoordinationv1 "github.com/superproj/onex/pkg/generated/applyconfigurations/coordination/v1"
	applyconfigurationscorev1 "github.com/superproj/onex/pkg/generated/applyconfigurations/core/v1"
	applyconfigurationsflowcontrolv1 "github.com/superproj/onex/pkg/generated/applyconfigurations/flowcontrol/v1"
	internal "github.com/superproj/onex/pkg/generated/applyconfigurations/internal"
	applyconfigurationsmetav1 "github.com/superproj/onex/pkg/generated/applyconfigurations/meta/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	coordinationv1 "k8s.io/api/coordination/v1"
	corev1 "k8s.io/api/core/v1"
	flowcontrolv1 "k8s.io/api/flowcontrol/v1"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	clientgoapplyconfigurationsmetav1 "k8s.io/client-go/applyconfigurations/meta/v1"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=apiextensions.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("CustomResourceColumnDefinition"):
		return &apiextensionsv1.CustomResourceColumnDefinitionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceConversion"):
		return &apiextensionsv1.CustomResourceConversionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinition"):
		return &apiextensionsv1.CustomResourceDefinitionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionCondition"):
		return &apiextensionsv1.CustomResourceDefinitionConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionNames"):
		return &apiextensionsv1.CustomResourceDefinitionNamesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionSpec"):
		return &apiextensionsv1.CustomResourceDefinitionSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionStatus"):
		return &apiextensionsv1.CustomResourceDefinitionStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceDefinitionVersion"):
		return &apiextensionsv1.CustomResourceDefinitionVersionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceSubresources"):
		return &apiextensionsv1.CustomResourceSubresourcesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceSubresourceScale"):
		return &apiextensionsv1.CustomResourceSubresourceScaleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("CustomResourceValidation"):
		return &apiextensionsv1.CustomResourceValidationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ExternalDocumentation"):
		return &apiextensionsv1.ExternalDocumentationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("JSONSchemaProps"):
		return &apiextensionsv1.JSONSchemaPropsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SelectableField"):
		return &apiextensionsv1.SelectableFieldApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceReference"):
		return &apiextensionsv1.ServiceReferenceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ValidationRule"):
		return &apiextensionsv1.ValidationRuleApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebhookClientConfig"):
		return &apiextensionsv1.WebhookClientConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WebhookConversion"):
		return &apiextensionsv1.WebhookConversionApplyConfiguration{}

		// Group=apps.onex.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithKind("Chain"):
		return &appsv1beta1.ChainApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ChainSpec"):
		return &appsv1beta1.ChainSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ChainStatus"):
		return &appsv1beta1.ChainStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ChargeRequest"):
		return &appsv1beta1.ChargeRequestApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ChargeRequestSpec"):
		return &appsv1beta1.ChargeRequestSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ChargeRequestStatus"):
		return &appsv1beta1.ChargeRequestStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Condition"):
		return &appsv1beta1.ConditionApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("LocalObjectReference"):
		return &appsv1beta1.LocalObjectReferenceApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("Miner"):
		return &appsv1beta1.MinerApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerAddress"):
		return &appsv1beta1.MinerAddressApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerSet"):
		return &appsv1beta1.MinerSetApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerSetSpec"):
		return &appsv1beta1.MinerSetSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerSetStatus"):
		return &appsv1beta1.MinerSetStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerSpec"):
		return &appsv1beta1.MinerSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerStatus"):
		return &appsv1beta1.MinerStatusApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("MinerTemplateSpec"):
		return &appsv1beta1.MinerTemplateSpecApplyConfiguration{}
	case v1beta1.SchemeGroupVersion.WithKind("ObjectMeta"):
		return &appsv1beta1.ObjectMetaApplyConfiguration{}

		// Group=autoscaling, Version=v1
	case autoscalingv1.SchemeGroupVersion.WithKind("CrossVersionObjectReference"):
		return &applyconfigurationsautoscalingv1.CrossVersionObjectReferenceApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("HorizontalPodAutoscaler"):
		return &applyconfigurationsautoscalingv1.HorizontalPodAutoscalerApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("HorizontalPodAutoscalerSpec"):
		return &applyconfigurationsautoscalingv1.HorizontalPodAutoscalerSpecApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("HorizontalPodAutoscalerStatus"):
		return &applyconfigurationsautoscalingv1.HorizontalPodAutoscalerStatusApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("Scale"):
		return &applyconfigurationsautoscalingv1.ScaleApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("ScaleSpec"):
		return &applyconfigurationsautoscalingv1.ScaleSpecApplyConfiguration{}
	case autoscalingv1.SchemeGroupVersion.WithKind("ScaleStatus"):
		return &applyconfigurationsautoscalingv1.ScaleStatusApplyConfiguration{}

		// Group=coordination.k8s.io, Version=v1
	case coordinationv1.SchemeGroupVersion.WithKind("Lease"):
		return &applyconfigurationscoordinationv1.LeaseApplyConfiguration{}
	case coordinationv1.SchemeGroupVersion.WithKind("LeaseSpec"):
		return &applyconfigurationscoordinationv1.LeaseSpecApplyConfiguration{}

		// Group=core, Version=v1
	case corev1.SchemeGroupVersion.WithKind("Affinity"):
		return &applyconfigurationscorev1.AffinityApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("AppArmorProfile"):
		return &applyconfigurationscorev1.AppArmorProfileApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("AttachedVolume"):
		return &applyconfigurationscorev1.AttachedVolumeApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("AWSElasticBlockStoreVolumeSource"):
		return &applyconfigurationscorev1.AWSElasticBlockStoreVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("AzureDiskVolumeSource"):
		return &applyconfigurationscorev1.AzureDiskVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("AzureFilePersistentVolumeSource"):
		return &applyconfigurationscorev1.AzureFilePersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("AzureFileVolumeSource"):
		return &applyconfigurationscorev1.AzureFileVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Capabilities"):
		return &applyconfigurationscorev1.CapabilitiesApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("CephFSPersistentVolumeSource"):
		return &applyconfigurationscorev1.CephFSPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("CephFSVolumeSource"):
		return &applyconfigurationscorev1.CephFSVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("CinderPersistentVolumeSource"):
		return &applyconfigurationscorev1.CinderPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("CinderVolumeSource"):
		return &applyconfigurationscorev1.CinderVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ClientIPConfig"):
		return &applyconfigurationscorev1.ClientIPConfigApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ClusterTrustBundleProjection"):
		return &applyconfigurationscorev1.ClusterTrustBundleProjectionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ComponentCondition"):
		return &applyconfigurationscorev1.ComponentConditionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ComponentStatus"):
		return &applyconfigurationscorev1.ComponentStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ConfigMap"):
		return &applyconfigurationscorev1.ConfigMapApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ConfigMapEnvSource"):
		return &applyconfigurationscorev1.ConfigMapEnvSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ConfigMapKeySelector"):
		return &applyconfigurationscorev1.ConfigMapKeySelectorApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ConfigMapNodeConfigSource"):
		return &applyconfigurationscorev1.ConfigMapNodeConfigSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ConfigMapProjection"):
		return &applyconfigurationscorev1.ConfigMapProjectionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ConfigMapVolumeSource"):
		return &applyconfigurationscorev1.ConfigMapVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Container"):
		return &applyconfigurationscorev1.ContainerApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerImage"):
		return &applyconfigurationscorev1.ContainerImageApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerPort"):
		return &applyconfigurationscorev1.ContainerPortApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerResizePolicy"):
		return &applyconfigurationscorev1.ContainerResizePolicyApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerState"):
		return &applyconfigurationscorev1.ContainerStateApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerStateRunning"):
		return &applyconfigurationscorev1.ContainerStateRunningApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerStateTerminated"):
		return &applyconfigurationscorev1.ContainerStateTerminatedApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerStateWaiting"):
		return &applyconfigurationscorev1.ContainerStateWaitingApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerStatus"):
		return &applyconfigurationscorev1.ContainerStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ContainerUser"):
		return &applyconfigurationscorev1.ContainerUserApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("CSIPersistentVolumeSource"):
		return &applyconfigurationscorev1.CSIPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("CSIVolumeSource"):
		return &applyconfigurationscorev1.CSIVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("DaemonEndpoint"):
		return &applyconfigurationscorev1.DaemonEndpointApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("DownwardAPIProjection"):
		return &applyconfigurationscorev1.DownwardAPIProjectionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("DownwardAPIVolumeFile"):
		return &applyconfigurationscorev1.DownwardAPIVolumeFileApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("DownwardAPIVolumeSource"):
		return &applyconfigurationscorev1.DownwardAPIVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EmptyDirVolumeSource"):
		return &applyconfigurationscorev1.EmptyDirVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EndpointAddress"):
		return &applyconfigurationscorev1.EndpointAddressApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EndpointPort"):
		return &applyconfigurationscorev1.EndpointPortApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Endpoints"):
		return &applyconfigurationscorev1.EndpointsApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EndpointSubset"):
		return &applyconfigurationscorev1.EndpointSubsetApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EnvFromSource"):
		return &applyconfigurationscorev1.EnvFromSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EnvVar"):
		return &applyconfigurationscorev1.EnvVarApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EnvVarSource"):
		return &applyconfigurationscorev1.EnvVarSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EphemeralContainer"):
		return &applyconfigurationscorev1.EphemeralContainerApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EphemeralContainerCommon"):
		return &applyconfigurationscorev1.EphemeralContainerCommonApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EphemeralVolumeSource"):
		return &applyconfigurationscorev1.EphemeralVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Event"):
		return &applyconfigurationscorev1.EventApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EventSeries"):
		return &applyconfigurationscorev1.EventSeriesApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("EventSource"):
		return &applyconfigurationscorev1.EventSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ExecAction"):
		return &applyconfigurationscorev1.ExecActionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("FCVolumeSource"):
		return &applyconfigurationscorev1.FCVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("FlexPersistentVolumeSource"):
		return &applyconfigurationscorev1.FlexPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("FlexVolumeSource"):
		return &applyconfigurationscorev1.FlexVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("FlockerVolumeSource"):
		return &applyconfigurationscorev1.FlockerVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("GCEPersistentDiskVolumeSource"):
		return &applyconfigurationscorev1.GCEPersistentDiskVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("GitRepoVolumeSource"):
		return &applyconfigurationscorev1.GitRepoVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("GlusterfsPersistentVolumeSource"):
		return &applyconfigurationscorev1.GlusterfsPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("GlusterfsVolumeSource"):
		return &applyconfigurationscorev1.GlusterfsVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("GRPCAction"):
		return &applyconfigurationscorev1.GRPCActionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("HostAlias"):
		return &applyconfigurationscorev1.HostAliasApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("HostIP"):
		return &applyconfigurationscorev1.HostIPApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("HostPathVolumeSource"):
		return &applyconfigurationscorev1.HostPathVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("HTTPGetAction"):
		return &applyconfigurationscorev1.HTTPGetActionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("HTTPHeader"):
		return &applyconfigurationscorev1.HTTPHeaderApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ImageVolumeSource"):
		return &applyconfigurationscorev1.ImageVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ISCSIPersistentVolumeSource"):
		return &applyconfigurationscorev1.ISCSIPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ISCSIVolumeSource"):
		return &applyconfigurationscorev1.ISCSIVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("KeyToPath"):
		return &applyconfigurationscorev1.KeyToPathApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Lifecycle"):
		return &applyconfigurationscorev1.LifecycleApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LifecycleHandler"):
		return &applyconfigurationscorev1.LifecycleHandlerApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LimitRange"):
		return &applyconfigurationscorev1.LimitRangeApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LimitRangeItem"):
		return &applyconfigurationscorev1.LimitRangeItemApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LimitRangeSpec"):
		return &applyconfigurationscorev1.LimitRangeSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LinuxContainerUser"):
		return &applyconfigurationscorev1.LinuxContainerUserApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LoadBalancerIngress"):
		return &applyconfigurationscorev1.LoadBalancerIngressApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LoadBalancerStatus"):
		return &applyconfigurationscorev1.LoadBalancerStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LocalObjectReference"):
		return &applyconfigurationscorev1.LocalObjectReferenceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("LocalVolumeSource"):
		return &applyconfigurationscorev1.LocalVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ModifyVolumeStatus"):
		return &applyconfigurationscorev1.ModifyVolumeStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Namespace"):
		return &applyconfigurationscorev1.NamespaceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NamespaceCondition"):
		return &applyconfigurationscorev1.NamespaceConditionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NamespaceSpec"):
		return &applyconfigurationscorev1.NamespaceSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NamespaceStatus"):
		return &applyconfigurationscorev1.NamespaceStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NFSVolumeSource"):
		return &applyconfigurationscorev1.NFSVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Node"):
		return &applyconfigurationscorev1.NodeApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeAddress"):
		return &applyconfigurationscorev1.NodeAddressApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeAffinity"):
		return &applyconfigurationscorev1.NodeAffinityApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeCondition"):
		return &applyconfigurationscorev1.NodeConditionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeConfigSource"):
		return &applyconfigurationscorev1.NodeConfigSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeConfigStatus"):
		return &applyconfigurationscorev1.NodeConfigStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeDaemonEndpoints"):
		return &applyconfigurationscorev1.NodeDaemonEndpointsApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeFeatures"):
		return &applyconfigurationscorev1.NodeFeaturesApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeRuntimeHandler"):
		return &applyconfigurationscorev1.NodeRuntimeHandlerApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeRuntimeHandlerFeatures"):
		return &applyconfigurationscorev1.NodeRuntimeHandlerFeaturesApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeSelector"):
		return &applyconfigurationscorev1.NodeSelectorApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeSelectorRequirement"):
		return &applyconfigurationscorev1.NodeSelectorRequirementApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeSelectorTerm"):
		return &applyconfigurationscorev1.NodeSelectorTermApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeSpec"):
		return &applyconfigurationscorev1.NodeSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeStatus"):
		return &applyconfigurationscorev1.NodeStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("NodeSystemInfo"):
		return &applyconfigurationscorev1.NodeSystemInfoApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ObjectFieldSelector"):
		return &applyconfigurationscorev1.ObjectFieldSelectorApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ObjectReference"):
		return &applyconfigurationscorev1.ObjectReferenceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolume"):
		return &applyconfigurationscorev1.PersistentVolumeApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaim"):
		return &applyconfigurationscorev1.PersistentVolumeClaimApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaimCondition"):
		return &applyconfigurationscorev1.PersistentVolumeClaimConditionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaimSpec"):
		return &applyconfigurationscorev1.PersistentVolumeClaimSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaimStatus"):
		return &applyconfigurationscorev1.PersistentVolumeClaimStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaimTemplate"):
		return &applyconfigurationscorev1.PersistentVolumeClaimTemplateApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeClaimVolumeSource"):
		return &applyconfigurationscorev1.PersistentVolumeClaimVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeSource"):
		return &applyconfigurationscorev1.PersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeSpec"):
		return &applyconfigurationscorev1.PersistentVolumeSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PersistentVolumeStatus"):
		return &applyconfigurationscorev1.PersistentVolumeStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PhotonPersistentDiskVolumeSource"):
		return &applyconfigurationscorev1.PhotonPersistentDiskVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Pod"):
		return &applyconfigurationscorev1.PodApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodAffinity"):
		return &applyconfigurationscorev1.PodAffinityApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodAffinityTerm"):
		return &applyconfigurationscorev1.PodAffinityTermApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodAntiAffinity"):
		return &applyconfigurationscorev1.PodAntiAffinityApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodCondition"):
		return &applyconfigurationscorev1.PodConditionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodDNSConfig"):
		return &applyconfigurationscorev1.PodDNSConfigApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodDNSConfigOption"):
		return &applyconfigurationscorev1.PodDNSConfigOptionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodIP"):
		return &applyconfigurationscorev1.PodIPApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodOS"):
		return &applyconfigurationscorev1.PodOSApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodReadinessGate"):
		return &applyconfigurationscorev1.PodReadinessGateApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodResourceClaim"):
		return &applyconfigurationscorev1.PodResourceClaimApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodResourceClaimStatus"):
		return &applyconfigurationscorev1.PodResourceClaimStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodSchedulingGate"):
		return &applyconfigurationscorev1.PodSchedulingGateApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodSecurityContext"):
		return &applyconfigurationscorev1.PodSecurityContextApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodSpec"):
		return &applyconfigurationscorev1.PodSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodStatus"):
		return &applyconfigurationscorev1.PodStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodTemplate"):
		return &applyconfigurationscorev1.PodTemplateApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PodTemplateSpec"):
		return &applyconfigurationscorev1.PodTemplateSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PortStatus"):
		return &applyconfigurationscorev1.PortStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PortworxVolumeSource"):
		return &applyconfigurationscorev1.PortworxVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("PreferredSchedulingTerm"):
		return &applyconfigurationscorev1.PreferredSchedulingTermApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Probe"):
		return &applyconfigurationscorev1.ProbeApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ProbeHandler"):
		return &applyconfigurationscorev1.ProbeHandlerApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ProjectedVolumeSource"):
		return &applyconfigurationscorev1.ProjectedVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("QuobyteVolumeSource"):
		return &applyconfigurationscorev1.QuobyteVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("RBDPersistentVolumeSource"):
		return &applyconfigurationscorev1.RBDPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("RBDVolumeSource"):
		return &applyconfigurationscorev1.RBDVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ReplicationController"):
		return &applyconfigurationscorev1.ReplicationControllerApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ReplicationControllerCondition"):
		return &applyconfigurationscorev1.ReplicationControllerConditionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ReplicationControllerSpec"):
		return &applyconfigurationscorev1.ReplicationControllerSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ReplicationControllerStatus"):
		return &applyconfigurationscorev1.ReplicationControllerStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceClaim"):
		return &applyconfigurationscorev1.ResourceClaimApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceFieldSelector"):
		return &applyconfigurationscorev1.ResourceFieldSelectorApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceHealth"):
		return &applyconfigurationscorev1.ResourceHealthApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceQuota"):
		return &applyconfigurationscorev1.ResourceQuotaApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceQuotaSpec"):
		return &applyconfigurationscorev1.ResourceQuotaSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceQuotaStatus"):
		return &applyconfigurationscorev1.ResourceQuotaStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceRequirements"):
		return &applyconfigurationscorev1.ResourceRequirementsApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ResourceStatus"):
		return &applyconfigurationscorev1.ResourceStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ScaleIOPersistentVolumeSource"):
		return &applyconfigurationscorev1.ScaleIOPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ScaleIOVolumeSource"):
		return &applyconfigurationscorev1.ScaleIOVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ScopedResourceSelectorRequirement"):
		return &applyconfigurationscorev1.ScopedResourceSelectorRequirementApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ScopeSelector"):
		return &applyconfigurationscorev1.ScopeSelectorApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SeccompProfile"):
		return &applyconfigurationscorev1.SeccompProfileApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Secret"):
		return &applyconfigurationscorev1.SecretApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SecretEnvSource"):
		return &applyconfigurationscorev1.SecretEnvSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SecretKeySelector"):
		return &applyconfigurationscorev1.SecretKeySelectorApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SecretProjection"):
		return &applyconfigurationscorev1.SecretProjectionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SecretReference"):
		return &applyconfigurationscorev1.SecretReferenceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SecretVolumeSource"):
		return &applyconfigurationscorev1.SecretVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SecurityContext"):
		return &applyconfigurationscorev1.SecurityContextApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SELinuxOptions"):
		return &applyconfigurationscorev1.SELinuxOptionsApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Service"):
		return &applyconfigurationscorev1.ServiceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ServiceAccount"):
		return &applyconfigurationscorev1.ServiceAccountApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ServiceAccountTokenProjection"):
		return &applyconfigurationscorev1.ServiceAccountTokenProjectionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ServicePort"):
		return &applyconfigurationscorev1.ServicePortApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ServiceSpec"):
		return &applyconfigurationscorev1.ServiceSpecApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("ServiceStatus"):
		return &applyconfigurationscorev1.ServiceStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SessionAffinityConfig"):
		return &applyconfigurationscorev1.SessionAffinityConfigApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("SleepAction"):
		return &applyconfigurationscorev1.SleepActionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("StorageOSPersistentVolumeSource"):
		return &applyconfigurationscorev1.StorageOSPersistentVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("StorageOSVolumeSource"):
		return &applyconfigurationscorev1.StorageOSVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Sysctl"):
		return &applyconfigurationscorev1.SysctlApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Taint"):
		return &applyconfigurationscorev1.TaintApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("TCPSocketAction"):
		return &applyconfigurationscorev1.TCPSocketActionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Toleration"):
		return &applyconfigurationscorev1.TolerationApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("TopologySpreadConstraint"):
		return &applyconfigurationscorev1.TopologySpreadConstraintApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("TypedLocalObjectReference"):
		return &applyconfigurationscorev1.TypedLocalObjectReferenceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("TypedObjectReference"):
		return &applyconfigurationscorev1.TypedObjectReferenceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("Volume"):
		return &applyconfigurationscorev1.VolumeApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeDevice"):
		return &applyconfigurationscorev1.VolumeDeviceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeMount"):
		return &applyconfigurationscorev1.VolumeMountApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeMountStatus"):
		return &applyconfigurationscorev1.VolumeMountStatusApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeNodeAffinity"):
		return &applyconfigurationscorev1.VolumeNodeAffinityApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeProjection"):
		return &applyconfigurationscorev1.VolumeProjectionApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeResourceRequirements"):
		return &applyconfigurationscorev1.VolumeResourceRequirementsApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VolumeSource"):
		return &applyconfigurationscorev1.VolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("VsphereVirtualDiskVolumeSource"):
		return &applyconfigurationscorev1.VsphereVirtualDiskVolumeSourceApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("WeightedPodAffinityTerm"):
		return &applyconfigurationscorev1.WeightedPodAffinityTermApplyConfiguration{}
	case corev1.SchemeGroupVersion.WithKind("WindowsSecurityContextOptions"):
		return &applyconfigurationscorev1.WindowsSecurityContextOptionsApplyConfiguration{}

		// Group=flowcontrol.apiserver.k8s.io, Version=v1
	case flowcontrolv1.SchemeGroupVersion.WithKind("ExemptPriorityLevelConfiguration"):
		return &applyconfigurationsflowcontrolv1.ExemptPriorityLevelConfigurationApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("FlowDistinguisherMethod"):
		return &applyconfigurationsflowcontrolv1.FlowDistinguisherMethodApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("FlowSchema"):
		return &applyconfigurationsflowcontrolv1.FlowSchemaApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("FlowSchemaCondition"):
		return &applyconfigurationsflowcontrolv1.FlowSchemaConditionApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("FlowSchemaSpec"):
		return &applyconfigurationsflowcontrolv1.FlowSchemaSpecApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("FlowSchemaStatus"):
		return &applyconfigurationsflowcontrolv1.FlowSchemaStatusApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("GroupSubject"):
		return &applyconfigurationsflowcontrolv1.GroupSubjectApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("LimitedPriorityLevelConfiguration"):
		return &applyconfigurationsflowcontrolv1.LimitedPriorityLevelConfigurationApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("LimitResponse"):
		return &applyconfigurationsflowcontrolv1.LimitResponseApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("NonResourcePolicyRule"):
		return &applyconfigurationsflowcontrolv1.NonResourcePolicyRuleApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("PolicyRulesWithSubjects"):
		return &applyconfigurationsflowcontrolv1.PolicyRulesWithSubjectsApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("PriorityLevelConfiguration"):
		return &applyconfigurationsflowcontrolv1.PriorityLevelConfigurationApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("PriorityLevelConfigurationCondition"):
		return &applyconfigurationsflowcontrolv1.PriorityLevelConfigurationConditionApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("PriorityLevelConfigurationReference"):
		return &applyconfigurationsflowcontrolv1.PriorityLevelConfigurationReferenceApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("PriorityLevelConfigurationSpec"):
		return &applyconfigurationsflowcontrolv1.PriorityLevelConfigurationSpecApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("PriorityLevelConfigurationStatus"):
		return &applyconfigurationsflowcontrolv1.PriorityLevelConfigurationStatusApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("QueuingConfiguration"):
		return &applyconfigurationsflowcontrolv1.QueuingConfigurationApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("ResourcePolicyRule"):
		return &applyconfigurationsflowcontrolv1.ResourcePolicyRuleApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("ServiceAccountSubject"):
		return &applyconfigurationsflowcontrolv1.ServiceAccountSubjectApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("Subject"):
		return &applyconfigurationsflowcontrolv1.SubjectApplyConfiguration{}
	case flowcontrolv1.SchemeGroupVersion.WithKind("UserSubject"):
		return &applyconfigurationsflowcontrolv1.UserSubjectApplyConfiguration{}

		// Group=meta.k8s.io, Version=v1
	case metav1.SchemeGroupVersion.WithKind("Condition"):
		return &applyconfigurationsmetav1.ConditionApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("DeleteOptions"):
		return &clientgoapplyconfigurationsmetav1.DeleteOptionsApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("LabelSelector"):
		return &applyconfigurationsmetav1.LabelSelectorApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("LabelSelectorRequirement"):
		return &applyconfigurationsmetav1.LabelSelectorRequirementApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("ManagedFieldsEntry"):
		return &applyconfigurationsmetav1.ManagedFieldsEntryApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("ObjectMeta"):
		return &applyconfigurationsmetav1.ObjectMetaApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("OwnerReference"):
		return &applyconfigurationsmetav1.OwnerReferenceApplyConfiguration{}
	case metav1.SchemeGroupVersion.WithKind("TypeMeta"):
		return &applyconfigurationsmetav1.TypeMetaApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
