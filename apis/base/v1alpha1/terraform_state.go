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
/*
Copyright Royal Dutch Shell
*/
package v1alpha1

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces
type TerraFormState struct {
	// +optional
	InitLogs string `json:"init_logs,omitempty"`
	// +optional
	InitErrLogs string `json:"init_err_logs,omitempty"`
	// +optional
	ApplyLogs string `json:"apply_logs,omitempty"`
	// +optional
	ApplyErrLogs string `json:"apply_err_logs,omitempty"`
	// +optional
	Plan string `json:"plan,omitempty"`
}
