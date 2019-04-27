package machine

import (
	"context"
	"encoding/base64"
	"strings"
	"testing"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/ghodss/yaml"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/fake"
	"github.com/openshift/cluster-api/pkg/controller/machine"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	providerv1 "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/virtualmachines"
	controllerfake "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var (
	_ machine.Actuator = (*Actuator)(nil)
)

func newClusterProviderSpec() providerv1.AzureClusterProviderSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return providerv1.AzureClusterProviderSpec{ResourceGroup: "resource-group-test", Location: "southcentralus"}
}
func providerSpecFromMachine(in *providerv1.AzureMachineProviderSpec) (*machinev1.ProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := yaml.Marshal(in)
	if err != nil {
		return nil, err
	}
	return &machinev1.ProviderSpec{Value: &runtime.RawExtension{Raw: bytes}}, nil
}
func providerSpecFromCluster(in *providerv1.AzureClusterProviderSpec) (*clusterv1.ProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := yaml.Marshal(in)
	if err != nil {
		return nil, err
	}
	return &clusterv1.ProviderSpec{Value: &runtime.RawExtension{Raw: bytes}}, nil
}
func newMachine(t *testing.T, machineConfig providerv1.AzureMachineProviderSpec, labels map[string]string) *machinev1.Machine {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	providerSpec, err := providerSpecFromMachine(&machineConfig)
	if err != nil {
		t.Fatalf("error encoding provider config: %v", err)
	}
	return &machinev1.Machine{ObjectMeta: metav1.ObjectMeta{Name: "machine-test", Labels: labels}, Spec: machinev1.MachineSpec{ProviderSpec: *providerSpec, Versions: machinev1.MachineVersionInfo{Kubelet: "1.9.4", ControlPlane: "1.9.4"}}}
}
func newCluster(t *testing.T) *clusterv1.Cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterProviderSpec := newClusterProviderSpec()
	providerSpec, err := providerSpecFromCluster(&clusterProviderSpec)
	if err != nil {
		t.Fatalf("error encoding provider config: %v", err)
	}
	return &clusterv1.Cluster{TypeMeta: metav1.TypeMeta{Kind: "Cluster"}, ObjectMeta: metav1.ObjectMeta{Name: "cluster-test"}, Spec: clusterv1.ClusterSpec{ClusterNetwork: clusterv1.ClusterNetworkingConfig{Services: clusterv1.NetworkRanges{CIDRBlocks: []string{"10.96.0.0/12"}}, Pods: clusterv1.NetworkRanges{CIDRBlocks: []string{"192.168.0.0/16"}}}, ProviderSpec: *providerSpec}}
}
func newFakeScope(t *testing.T, label string) *actuators.MachineScope {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope := &actuators.Scope{Context: context.Background(), Cluster: newCluster(t), ClusterConfig: &v1alpha1.AzureClusterProviderSpec{ResourceGroup: "dummyResourceGroup", Location: "dummyLocation", CAKeyPair: v1alpha1.KeyPair{Cert: []byte("cert"), Key: []byte("key")}, EtcdCAKeyPair: v1alpha1.KeyPair{Cert: []byte("cert"), Key: []byte("key")}, FrontProxyCAKeyPair: v1alpha1.KeyPair{Cert: []byte("cert"), Key: []byte("key")}, SAKeyPair: v1alpha1.KeyPair{Cert: []byte("cert"), Key: []byte("key")}, DiscoveryHashes: []string{"discoveryhash0"}}, ClusterStatus: &v1alpha1.AzureClusterProviderStatus{}}
	scope.Network().APIServerIP.DNSName = "DummyDNSName"
	labels := make(map[string]string)
	labels[v1alpha1.MachineRoleLabel] = label
	machineConfig := providerv1.AzureMachineProviderSpec{}
	m := newMachine(t, machineConfig, labels)
	c := fake.NewSimpleClientset(m).MachineV1beta1()
	return &actuators.MachineScope{Scope: scope, Machine: m, MachineClient: c.Machines("dummyNamespace"), MachineConfig: &v1alpha1.AzureMachineProviderSpec{}, MachineStatus: &v1alpha1.AzureMachineProviderStatus{}}
}
func newFakeReconciler(t *testing.T) *Reconciler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeVMSuccessSvc := &FakeVMService{Name: "machine-test", ID: "machine-test-ID", ProvisioningState: "Succeeded"}
	return &Reconciler{scope: newFakeScope(t, v1alpha1.ControlPlane), availabilityZonesSvc: fakeSuccessSvc, networkInterfacesSvc: fakeSuccessSvc, virtualMachinesSvc: fakeVMSuccessSvc, virtualMachinesExtSvc: fakeSuccessSvc}
}
func newFakeReconcilerWithScope(t *testing.T, scope *actuators.MachineScope) *Reconciler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeVMSuccessSvc := &FakeVMService{Name: "machine-test", ID: "machine-test-ID", ProvisioningState: "Succeeded"}
	return &Reconciler{scope: scope, availabilityZonesSvc: fakeSuccessSvc, networkInterfacesSvc: fakeSuccessSvc, virtualMachinesSvc: fakeVMSuccessSvc, virtualMachinesExtSvc: fakeSuccessSvc}
}

type FakeVMService struct {
	Name			string
	ID			string
	ProvisioningState	string
	GetCallCount		int
	CreateOrUpdateCallCount	int
	DeleteCallCount		int
}

func (s *FakeVMService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.GetCallCount++
	return compute.VirtualMachine{ID: to.StringPtr(s.ID), Name: to.StringPtr(s.Name), VirtualMachineProperties: &compute.VirtualMachineProperties{ProvisioningState: to.StringPtr(s.ProvisioningState)}}, nil
}
func (s *FakeVMService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.CreateOrUpdateCallCount++
	return nil
}
func (s *FakeVMService) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.DeleteCallCount++
	return nil
}
func TestReconcilerSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeReconciler := newFakeReconciler(t)
	if err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}
	if err := fakeReconciler.Update(context.Background()); err != nil {
		t.Errorf("failed to update machine: %+v", err)
	}
	if _, err := fakeReconciler.Exists(context.Background()); err != nil {
		t.Errorf("failed to check if machine exists: %+v", err)
	}
	if err := fakeReconciler.Delete(context.Background()); err != nil {
		t.Errorf("failed to delete machine: %+v", err)
	}
}
func TestReconcileFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeFailureSvc := &azure.FakeFailureService{}
	fakeReconciler := newFakeReconciler(t)
	fakeReconciler.networkInterfacesSvc = fakeFailureSvc
	fakeReconciler.virtualMachinesSvc = fakeFailureSvc
	fakeReconciler.virtualMachinesExtSvc = fakeFailureSvc
	if err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}
	if err := fakeReconciler.Update(context.Background()); err == nil {
		t.Errorf("expected update to fail")
	}
	if _, err := fakeReconciler.Exists(context.Background()); err == nil {
		t.Errorf("expected exists to fail")
	}
	if err := fakeReconciler.Delete(context.Background()); err == nil {
		t.Errorf("expected delete to fail")
	}
}
func TestReconcileVMFailedState(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeReconciler := newFakeReconciler(t)
	fakeVMService := &FakeVMService{Name: "machine-test", ID: "machine-test-ID", ProvisioningState: "Failed"}
	fakeReconciler.virtualMachinesSvc = fakeVMService
	if err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}
	if fakeVMService.GetCallCount != 1 {
		t.Errorf("expected get to be called just once")
	}
	if fakeVMService.DeleteCallCount != 1 {
		t.Errorf("expected delete to be called just once")
	}
	if fakeVMService.CreateOrUpdateCallCount != 0 {
		t.Errorf("expected createorupdate not to be called")
	}
}
func TestReconcileVMUpdatingState(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeReconciler := newFakeReconciler(t)
	fakeVMService := &FakeVMService{Name: "machine-test", ID: "machine-test-ID", ProvisioningState: "Updating"}
	fakeReconciler.virtualMachinesSvc = fakeVMService
	if err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}
	if fakeVMService.GetCallCount != 1 {
		t.Errorf("expected get to be called just once")
	}
	if fakeVMService.DeleteCallCount != 0 {
		t.Errorf("expected delete not to be called")
	}
	if fakeVMService.CreateOrUpdateCallCount != 0 {
		t.Errorf("expected createorupdate not to be called")
	}
}
func TestReconcileVMSuceededState(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeReconciler := newFakeReconciler(t)
	fakeVMService := &FakeVMService{Name: "machine-test", ID: "machine-test-ID", ProvisioningState: "Succeeded"}
	fakeReconciler.virtualMachinesSvc = fakeVMService
	if err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}
	if fakeVMService.GetCallCount != 1 {
		t.Errorf("expected get to be called just once")
	}
	if fakeVMService.DeleteCallCount != 0 {
		t.Errorf("expected delete not to be called")
	}
	if fakeVMService.CreateOrUpdateCallCount != 0 {
		t.Errorf("expected createorupdate not to be called")
	}
}
func TestNodeJoinFirstControlPlane(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeReconciler := newFakeReconciler(t)
	if isNodeJoin, err := fakeReconciler.isNodeJoin(); err != nil {
		t.Errorf("isNodeJoin failed to create machine: %+v", err)
	} else if isNodeJoin {
		t.Errorf("Expected isNodeJoin to be false since its first VM")
	}
}
func TestNodeJoinNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeScope := newFakeScope(t, v1alpha1.Node)
	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope)
	if isNodeJoin, err := fakeReconciler.isNodeJoin(); err != nil {
		t.Errorf("isNodeJoin failed to create machine: %+v", err)
	} else if !isNodeJoin {
		t.Errorf("Expected isNodeJoin to be true since its a node")
	}
}
func TestNodeJoinSecondControlPlane(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeScope := newFakeScope(t, v1alpha1.ControlPlane)
	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope)
	if _, err := fakeScope.MachineClient.Create(fakeScope.Machine); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}
	if isNodeJoin, err := fakeReconciler.isNodeJoin(); err != nil {
		t.Errorf("isNodeJoin failed to create machine: %+v", err)
	} else if !isNodeJoin {
		t.Errorf("Expected isNodeJoin to be true since its second control plane vm")
	}
}

type FakeVMCheckZonesService struct{ checkZones []string }

func (s *FakeVMCheckZonesService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.New("vm not found")
}
func (s *FakeVMCheckZonesService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmSpec, ok := spec.(*virtualmachines.Spec)
	if !ok {
		return errors.New("invalid vm specification")
	}
	if len(s.checkZones) <= 0 {
		return nil
	}
	for _, zone := range s.checkZones {
		if strings.EqualFold(zone, vmSpec.Zone) {
			return nil
		}
	}
	return errors.New("invalid input zone")
}
func (s *FakeVMCheckZonesService) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type FakeAvailabilityZonesService struct{ zonesResponse []string }

func (s *FakeAvailabilityZonesService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.zonesResponse, nil
}
func (s *FakeAvailabilityZonesService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *FakeAvailabilityZonesService) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func TestAvailabilityZones(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeScope := newFakeScope(t, v1alpha1.ControlPlane)
	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope)
	zones := []string{"1", "2", "3"}
	fakeReconciler.availabilityZonesSvc = &FakeAvailabilityZonesService{zonesResponse: zones}
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{checkZones: zones}
	if err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}
	fakeReconciler.availabilityZonesSvc = &FakeAvailabilityZonesService{zonesResponse: []string{}}
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{checkZones: []string{}}
	if err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}
	fakeReconciler.availabilityZonesSvc = &FakeAvailabilityZonesService{zonesResponse: []string{"2"}}
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{checkZones: []string{"3"}}
	if err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail due to zone mismatch")
	}
}
func TestCustomUserData(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeScope := newFakeScope(t, v1alpha1.Node)
	userDataSecret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "testCustomUserData", Namespace: "dummyNamespace"}, Data: map[string][]byte{"userData": []byte("test-userdata")}}
	fakeScope.CoreClient = controllerfake.NewFakeClient(userDataSecret)
	fakeScope.MachineConfig.UserDataSecret = &corev1.SecretReference{Name: "testCustomUserData"}
	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope)
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{}
	if err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("expected create to succeed %v", err)
	}
	userData, err := fakeReconciler.getCustomUserData()
	if err != nil {
		t.Errorf("expected get custom data to succeed %v", err)
	}
	if userData != base64.StdEncoding.EncodeToString([]byte("test-userdata")) {
		t.Errorf("expected userdata to be test-userdata, but found %s", userData)
	}
}
func TestCustomDataFailures(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeScope := newFakeScope(t, v1alpha1.Node)
	userDataSecret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "testCustomUserData", Namespace: "dummyNamespace"}, Data: map[string][]byte{"userData": []byte("test-userdata")}}
	fakeScope.CoreClient = controllerfake.NewFakeClient(userDataSecret)
	fakeScope.MachineConfig.UserDataSecret = &corev1.SecretReference{Name: "testCustomUserData"}
	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope)
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{}
	fakeScope.MachineConfig.UserDataSecret = &corev1.SecretReference{Name: "testFailure"}
	if err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}
	if _, err := fakeReconciler.getCustomUserData(); err == nil {
		t.Errorf("expected get custom data to fail")
	}
	userDataSecret.Data = map[string][]byte{"notUserData": []byte("test-notuserdata")}
	fakeScope.CoreClient = controllerfake.NewFakeClient(userDataSecret)
	if _, err := fakeReconciler.getCustomUserData(); err == nil {
		t.Errorf("expected get custom data to fail, due to missing userdata")
	}
}
