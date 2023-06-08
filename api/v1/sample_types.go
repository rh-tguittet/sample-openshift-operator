package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SampleSpec defines the desired state of Sample
type SampleSpec struct {
	// +kubebuilder:validation:MinLength=1
	Foo string `json:"foo,omitempty"`
}

// SampleStatus defines the observed state of Sample
type SampleStatus struct {
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Sample is the Schema for the samples API
type Sample struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SampleSpec   `json:"spec,omitempty"`
	Status SampleStatus `json:"status,omitempty"`
	Bar    string       `json:"pouet,omitempty"`
}

//+kubebuilder:object:root=true

// SampleList contains a list of Sample
type SampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sample `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sample{}, &SampleList{})
}
