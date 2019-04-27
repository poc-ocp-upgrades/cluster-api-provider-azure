package certificates

import (
	"context"
	"reflect"
	"testing"
	"time"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

func TestCreateOrUpdateCertificates(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope := actuators.Scope{ClusterConfig: &v1alpha1.AzureClusterProviderSpec{Location: "eastus"}, ClusterStatus: &v1alpha1.AzureClusterProviderStatus{}, Cluster: &clusterv1.Cluster{ObjectMeta: metav1.ObjectMeta{Name: "dummyclustername"}}}
	scope.Network().APIServerIP.DNSName = "dummydnsname"
	certsSvc := NewService(&scope)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := certsSvc.CreateOrUpdate(ctx, nil); err != nil {
		t.Errorf("Error creating certificates: %v", err)
	}
	caKeyPair := scope.ClusterConfig.CAKeyPair
	if !scope.ClusterConfig.CAKeyPair.HasCertAndKey() {
		t.Errorf("Error creating ca keypair")
	}
	if !scope.ClusterConfig.SAKeyPair.HasCertAndKey() {
		t.Errorf("Error creating sa keypair")
	}
	if !scope.ClusterConfig.EtcdCAKeyPair.HasCertAndKey() {
		t.Errorf("Error creating etcd ca keypair")
	}
	if scope.ClusterConfig.AdminKubeconfig == "" {
		t.Errorf("Error generating admin kube config")
	}
	if len(scope.ClusterConfig.DiscoveryHashes) <= 0 {
		t.Errorf("Error generating discovery hashes")
	}
	if err := certsSvc.CreateOrUpdate(ctx, nil); err != nil {
		t.Errorf("Error creating certificates: %v", err)
	}
	if !reflect.DeepEqual(scope.ClusterConfig.CAKeyPair, caKeyPair) {
		t.Errorf("Expected ca key pair not be regenerated")
	}
}
