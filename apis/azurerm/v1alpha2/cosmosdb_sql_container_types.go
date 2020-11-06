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

type CosmosdbSQLContainer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbSQLContainerSpec   `json:"spec,omitempty"`
	Status            CosmosdbSQLContainerStatus `json:"status,omitempty"`
}

type CosmosdbSQLContainerSpecUniqueKey struct {
	Paths []string `json:"paths" tf:"paths"`
}

type CosmosdbSQLContainerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountName  string `json:"accountName" tf:"account_name"`
	DatabaseName string `json:"databaseName" tf:"database_name"`
	// +optional
	DefaultTtl int64  `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	PartitionKeyPath  string `json:"partitionKeyPath,omitempty" tf:"partition_key_path,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Throughput int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
	// +optional
	UniqueKey []CosmosdbSQLContainerSpecUniqueKey `json:"uniqueKey,omitempty" tf:"unique_key,omitempty"`
}

type CosmosdbSQLContainerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CosmosdbSQLContainerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CosmosdbSQLContainerList is a list of CosmosdbSQLContainers
type CosmosdbSQLContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CosmosdbSQLContainer CRD objects
	Items []CosmosdbSQLContainer `json:"items,omitempty"`
}