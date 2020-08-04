/*
Copyright The Kubeform Authors.

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ManagedApplicationDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedApplicationDefinitionSpec   `json:"spec,omitempty"`
	Status            ManagedApplicationDefinitionStatus `json:"status,omitempty"`
}

type ManagedApplicationDefinitionSpecAuthorization struct {
	RoleDefinitionID   string `json:"roleDefinitionID" tf:"role_definition_id"`
	ServicePrincipalID string `json:"servicePrincipalID" tf:"service_principal_id"`
}

type ManagedApplicationDefinitionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MinItems=1
	Authorization []ManagedApplicationDefinitionSpecAuthorization `json:"authorization,omitempty" tf:"authorization,omitempty"`
	// +optional
	CreateUiDefinition string `json:"createUiDefinition,omitempty" tf:"create_ui_definition,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	Location    string `json:"location" tf:"location"`
	LockLevel   string `json:"lockLevel" tf:"lock_level"`
	// +optional
	MainTemplate string `json:"mainTemplate,omitempty" tf:"main_template,omitempty"`
	Name         string `json:"name" tf:"name"`
	// +optional
	PackageEnabled bool `json:"packageEnabled,omitempty" tf:"package_enabled,omitempty"`
	// +optional
	PackageFileURI    string `json:"packageFileURI,omitempty" tf:"package_file_uri,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagedApplicationDefinitionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ManagedApplicationDefinitionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ManagedApplicationDefinitionList is a list of ManagedApplicationDefinitions
type ManagedApplicationDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedApplicationDefinition CRD objects
	Items []ManagedApplicationDefinition `json:"items,omitempty"`
}