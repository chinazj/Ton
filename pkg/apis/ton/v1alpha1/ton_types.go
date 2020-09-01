package v1alpha1
import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FlowStatus struct {

}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Flow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlowSpec `json:"spec"`
	Status            FlowStatus `json:"status,omitempty"`
}

type FlowSpec struct {
	Tasks     []Task `json:"Tasks"`
}

type Task struct {
	Name    string `json:"name"`
	Commond string `json:"commond"`
	Image  string `json:"image"`
	ServiceAccountName string `json:"serviceaccountname"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FlowList is a list of Ton resources
type FlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Flow `json:"items"`
}
func init() {
	SchemeBuilder.Register(&Flow{}, &FlowList{})
}