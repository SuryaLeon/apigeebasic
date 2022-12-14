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

type ResourceFileObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResourceFileParameters struct {

	// +kubebuilder:validation:Required
	File *string `json:"file" tf:"file,omitempty"`

	// +kubebuilder:validation:Required
	FileHash *string `json:"fileHash" tf:"file_hash,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// ResourceFileSpec defines the desired state of ResourceFile
type ResourceFileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceFileParameters `json:"forProvider"`
}

// ResourceFileStatus defines the observed state of ResourceFile.
type ResourceFileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceFileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceFile is the Schema for the ResourceFiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigjet}
type ResourceFile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceFileSpec   `json:"spec"`
	Status            ResourceFileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceFileList contains a list of ResourceFiles
type ResourceFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceFile `json:"items"`
}

// Repository type metadata.
var (
	ResourceFile_Kind             = "ResourceFile"
	ResourceFile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceFile_Kind}.String()
	ResourceFile_KindAPIVersion   = ResourceFile_Kind + "." + CRDGroupVersion.String()
	ResourceFile_GroupVersionKind = CRDGroupVersion.WithKind(ResourceFile_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceFile{}, &ResourceFileList{})
}
