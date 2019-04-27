package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeadmv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
)

type AzureMachineProviderSpec struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	UserDataSecret		*corev1.SecretReference	`json:"userDataSecret,omitempty"`
	CredentialsSecret	*corev1.SecretReference	`json:"credentialsSecret,omitempty"`
	Location		string			`json:"location"`
	VMSize			string			`json:"vmSize"`
	Image			Image			`json:"image"`
	OSDisk			OSDisk			`json:"osDisk"`
	SSHPublicKey		string			`json:"sshPublicKey"`
	SSHPrivateKey		string			`json:"sshPrivateKey"`
}
type KubeadmConfiguration struct {
	Join	kubeadmv1beta1.JoinConfiguration	`json:"join,omitempty"`
	Init	kubeadmv1beta1.InitConfiguration	`json:"init,omitempty"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&AzureMachineProviderSpec{})
}
