package cluster

import (
	"context"
	"testing"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

func newFakeScope() *actuators.Scope {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &actuators.Scope{Context: context.Background(), Cluster: &clusterv1.Cluster{ObjectMeta: metav1.ObjectMeta{Name: "dummyClusterName"}}, ClusterConfig: &v1alpha1.AzureClusterProviderSpec{}, ClusterStatus: &v1alpha1.AzureClusterProviderStatus{}}
}
func TestReconcileSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeNotFoundSvc := &azure.FakeNotFoundService{}
	fakeReconciler := &Reconciler{scope: newFakeScope(), groupsSvc: fakeSuccessSvc, certificatesSvc: fakeSuccessSvc, vnetSvc: fakeSuccessSvc, securityGroupSvc: fakeSuccessSvc, routeTableSvc: fakeSuccessSvc, subnetsSvc: fakeSuccessSvc, internalLBSvc: fakeSuccessSvc, publicIPSvc: fakeSuccessSvc, publicLBSvc: fakeSuccessSvc}
	if err := fakeReconciler.Reconcile(); err != nil {
		t.Errorf("failed to reconcile cluster services: %+v", err)
	}
	if err := fakeReconciler.Delete(); err != nil {
		t.Errorf("failed to delete cluster services: %+v", err)
	}
	fakeReconciler.groupsSvc = fakeNotFoundSvc
	if err := fakeReconciler.Delete(); err != nil {
		t.Errorf("failed to delete cluster services: %+v", err)
	}
}
func TestReconcileFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeFailureSvc := &azure.FakeFailureService{}
	fakeReconciler := &Reconciler{scope: newFakeScope(), certificatesSvc: fakeFailureSvc, groupsSvc: fakeSuccessSvc, vnetSvc: fakeSuccessSvc, securityGroupSvc: fakeFailureSvc, routeTableSvc: fakeSuccessSvc, subnetsSvc: fakeSuccessSvc, internalLBSvc: fakeFailureSvc, publicIPSvc: fakeSuccessSvc, publicLBSvc: fakeSuccessSvc}
	if err := fakeReconciler.Reconcile(); err == nil {
		t.Errorf("expected reconcile to fail")
	}
	if err := fakeReconciler.Delete(); err == nil {
		t.Errorf("expected delete to fail")
	}
}
func TestPublicIPNonEmpty(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeReconciler := &Reconciler{scope: newFakeScope(), groupsSvc: fakeSuccessSvc, certificatesSvc: fakeSuccessSvc, vnetSvc: fakeSuccessSvc, securityGroupSvc: fakeSuccessSvc, routeTableSvc: fakeSuccessSvc, subnetsSvc: fakeSuccessSvc, internalLBSvc: fakeSuccessSvc, publicIPSvc: fakeSuccessSvc, publicLBSvc: fakeSuccessSvc}
	if err := fakeReconciler.Reconcile(); err != nil {
		t.Errorf("failed to reconcile cluster services: %+v", err)
	}
	ipName := fakeReconciler.scope.Network().APIServerIP.Name
	if ipName == "" {
		t.Errorf("public ip still empty, expected to be refilled")
	}
	if err := fakeReconciler.Reconcile(); err != nil {
		t.Errorf("failed to reconcile cluster services: %+v", err)
	}
	if fakeReconciler.scope.Network().APIServerIP.Name != ipName {
		t.Errorf("expected public ip to be not generated again")
	}
}
func TestServicesCreatedCount(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cache := make(map[string]int)
	fakeSuccessSvc := &azure.FakeCachedService{Cache: &cache}
	fakeReconciler := &Reconciler{scope: newFakeScope(), groupsSvc: fakeSuccessSvc, certificatesSvc: fakeSuccessSvc, vnetSvc: fakeSuccessSvc, securityGroupSvc: fakeSuccessSvc, routeTableSvc: fakeSuccessSvc, subnetsSvc: fakeSuccessSvc, internalLBSvc: fakeSuccessSvc, publicIPSvc: fakeSuccessSvc, publicLBSvc: fakeSuccessSvc}
	if err := fakeReconciler.Reconcile(); err != nil {
		t.Errorf("failed to reconcile cluster services: %+v", err)
	}
	if cache[azure.GenerateVnetName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateVnetName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GenerateControlPlaneSecurityGroupName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateControlPlaneSecurityGroupName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GenerateNodeSecurityGroupName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateNodeSecurityGroupName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GenerateNodeRouteTableName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateNodeRouteTableName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GenerateControlPlaneSubnetName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateControlPlaneSubnetName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GenerateNodeSubnetName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateNodeSubnetName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GenerateInternalLBName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GenerateInternalLBName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[azure.GeneratePublicLBName(fakeReconciler.scope.Cluster.Name)] != 1 {
		t.Errorf("Expected 1 count of %s service", azure.GeneratePublicLBName(fakeReconciler.scope.Cluster.Name))
	}
	if cache[fakeReconciler.scope.Network().APIServerIP.Name] != 1 {
		t.Errorf("Expected 1 count of %s service", fakeReconciler.scope.Network().APIServerIP.Name)
	}
}
