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

type KvmObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KvmParameters struct {

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Entry map[string]*string `json:"entry,omitempty" tf:"entry,omitempty"`

	// +kubebuilder:validation:Optional
	SensitiveEntrySecretRef map[string]v1.SecretKeySelector `json:"sensitiveEntrySecretRef" tf:"-"`
}

// KvmSpec defines the desired state of Kvm
type KvmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KvmParameters `json:"forProvider"`
}

// KvmStatus defines the observed state of Kvm.
type KvmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KvmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Kvm is the Schema for the Kvms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigjet}
type Kvm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KvmSpec   `json:"spec"`
	Status            KvmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KvmList contains a list of Kvms
type KvmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kvm `json:"items"`
}

// Repository type metadata.
var (
	Kvm_Kind             = "Kvm"
	Kvm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Kvm_Kind}.String()
	Kvm_KindAPIVersion   = Kvm_Kind + "." + CRDGroupVersion.String()
	Kvm_GroupVersionKind = CRDGroupVersion.WithKind(Kvm_Kind)
)

func init() {
	SchemeBuilder.Register(&Kvm{}, &KvmList{})
}
