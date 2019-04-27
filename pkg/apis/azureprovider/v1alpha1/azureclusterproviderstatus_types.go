package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AzureClusterProviderStatus struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Network			Network	`json:"network,omitempty"`
	Bastion			VM	`json:"bastion,omitempty"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&AzureClusterProviderStatus{})
}
