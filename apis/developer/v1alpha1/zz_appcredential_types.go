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

type AppCredentialObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppCredentialParameters struct {

	// +kubebuilder:validation:Optional
	APIProducts []*string `json:"apiProducts,omitempty" tf:"api_products,omitempty"`

	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// +kubebuilder:validation:Required
	ConsumerKey *string `json:"consumerKey" tf:"consumer_key,omitempty"`

	// +kubebuilder:validation:Required
	ConsumerSecretSecretRef v1.SecretKeySelector `json:"consumerSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DeveloperAppName *string `json:"developerAppName" tf:"developer_app_name,omitempty"`

	// +kubebuilder:validation:Required
	DeveloperEmail *string `json:"developerEmail" tf:"developer_email,omitempty"`

	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

// AppCredentialSpec defines the desired state of AppCredential
type AppCredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppCredentialParameters `json:"forProvider"`
}

// AppCredentialStatus defines the observed state of AppCredential.
type AppCredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppCredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppCredential is the Schema for the AppCredentials API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,apigjet}
type AppCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCredentialSpec   `json:"spec"`
	Status            AppCredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppCredentialList contains a list of AppCredentials
type AppCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppCredential `json:"items"`
}

// Repository type metadata.
var (
	AppCredential_Kind             = "AppCredential"
	AppCredential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppCredential_Kind}.String()
	AppCredential_KindAPIVersion   = AppCredential_Kind + "." + CRDGroupVersion.String()
	AppCredential_GroupVersionKind = CRDGroupVersion.WithKind(AppCredential_Kind)
)

func init() {
	SchemeBuilder.Register(&AppCredential{}, &AppCredentialList{})
}
