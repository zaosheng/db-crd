/*


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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DBcrdSpec defines the desired state of DBcrd
type DBcrdSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Replicas     *int32                      `json:"replicas"`
	Image        string                      `json:"image"`
	Port         int32                       `json:"port"`
	Nodeport     int32                       `json:"nodeport"`
	Envs         []corev1.EnvVar             `json:"envs,omitempty"`
	Resources    corev1.ResourceRequirements `json:"resources,omitempty"`
	VolumeMounts []corev1.VolumeMount        `json:"volumeMounts"`
	Volumes      []corev1.Volume             `json:"volumes"`
}

// DBcrdStatus defines the observed state of DBcrd
type DBcrdStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status"`
}

// +kubebuilder:object:root=true

// DBcrd is the Schema for the dbcrds API
type DBcrd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DBcrdSpec   `json:"spec,omitempty"`
	Status DBcrdStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DBcrdList contains a list of DBcrd
type DBcrdList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DBcrd `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DBcrd{}, &DBcrdList{})
}
