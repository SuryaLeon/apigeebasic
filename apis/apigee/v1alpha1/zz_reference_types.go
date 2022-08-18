/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ReferenceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReferenceParameters struct {

	// +kubebuilder:validation:Required
	EnvironmentName *string `json:"environmentName" tf:"environment_name,omitempty"`

	// +kubebuilder:validation:Required
	Refers *string `json:"refers" tf:"refers,omitempty"`

	// +kubebuilder:validation:Required
	ResourceType *string `json:"resourceType" tf:"resource_type,omitempty"`
}

// ReferenceSpec defines the desired state of Reference
type ReferenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReferenceParameters `json:"forProvider"`
}

// ReferenceStatus defines the observed state of Reference.
type ReferenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReferenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Reference is the Schema for the References API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigjet}
type Reference struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReferenceSpec   `json:"spec"`
	Status            ReferenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceList contains a list of References
type ReferenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reference `json:"items"`
}

// Repository type metadata.
var (
	Reference_Kind             = "Reference"
	Reference_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Reference_Kind}.String()
	Reference_KindAPIVersion   = Reference_Kind + "." + CRDGroupVersion.String()
	Reference_GroupVersionKind = CRDGroupVersion.WithKind(Reference_Kind)
)

func init() {
	SchemeBuilder.Register(&Reference{}, &ReferenceList{})
}
