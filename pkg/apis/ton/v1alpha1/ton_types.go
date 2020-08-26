package v1alpha1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Line struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LineSpec `json:"spec"`
}

type LineSpec struct {
	Jobs     []Jog `json:"jobs"`
}

type Jog struct {
	Name    string `json:"name"`
	Commond string `json:"commond"`
	Image  string `json:"image"`
	ServiceAccountName string `json:"serviceaccountname"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TonList is a list of Ton resources
type LineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Line `json:"items"`
}