package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AzureMachineProviderStatus struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	VMID			*string				`json:"vmId,omitempty"`
	VMState			*VMState			`json:"vmState,omitempty"`
	Conditions		[]AzureMachineProviderCondition	`json:"conditions,omitempty"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&AzureMachineProviderStatus{})
}
