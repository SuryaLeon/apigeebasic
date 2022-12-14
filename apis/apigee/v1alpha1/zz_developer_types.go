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

type DeveloperObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DeveloperParameters struct {

	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// +kubebuilder:validation:Required
	Email *string `json:"email" tf:"email,omitempty"`

	// +kubebuilder:validation:Required
	FirstName *string `json:"firstName" tf:"first_name,omitempty"`

	// +kubebuilder:validation:Required
	LastName *string `json:"lastName" tf:"last_name,omitempty"`

	// +kubebuilder:validation:Required
	UserName *string `json:"userName" tf:"user_name,omitempty"`
}

// DeveloperSpec defines the desired state of Developer
type DeveloperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeveloperParameters `json:"forProvider"`
}

// DeveloperStatus defines the observed state of Developer.
type DeveloperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeveloperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Developer is the Schema for the Developers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigjet}
type Developer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeveloperSpec   `json:"spec"`
	Status            DeveloperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeveloperList contains a list of Developers
type DeveloperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Developer `json:"items"`
}

// Repository type metadata.
var (
	Developer_Kind             = "Developer"
	Developer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Developer_Kind}.String()
	Developer_KindAPIVersion   = Developer_Kind + "." + CRDGroupVersion.String()
	Developer_GroupVersionKind = CRDGroupVersion.WithKind(Developer_Kind)
)

func init() {
	SchemeBuilder.Register(&Developer{}, &DeveloperList{})
}
