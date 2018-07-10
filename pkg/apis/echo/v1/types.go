package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EchoAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []EchoApp `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EchoApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              EchoAppSpec   `json:"spec"`
	Status            EchoAppStatus `json:"status,omitempty"`
}

type EchoAppSpec struct {
	Size  int32  `json:"size"`
	Image string `json:"image"`
}
type EchoAppStatus struct {
	Nodes []string `json:"nodes"`
}
