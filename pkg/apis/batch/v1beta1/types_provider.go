package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// ProviderSpec defines the configuration to use during node creation.
type ProviderSpec struct {
	// No more than one of the following may be specified.

	// value is an inlined, serialized representation of the resource
	// configuration. It is recommended that providers maintain their own
	// versioned API types that should be serialized/deserialized from this
	// field, akin to component config.
	// +optional
	// +kubebuilder:validation:XPreserveUnknownFields
	Value *runtime.RawExtension `json:"value,omitempty" protobuf:"bytes,1,opt,name=value"`
}
