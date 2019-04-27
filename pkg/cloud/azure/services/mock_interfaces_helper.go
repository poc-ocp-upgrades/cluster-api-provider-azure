package services

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	providerv1 "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

func MockVMExists() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockVMIfExists: func(resourceGroup string, name string) (*compute.VirtualMachine, error) {
		networkProfile := compute.NetworkProfile{NetworkInterfaces: &[]compute.NetworkInterfaceReference{{ID: to.StringPtr("001")}}}
		OsDiskName := fmt.Sprintf("OS_Disk_%v", name)
		storageProfile := compute.StorageProfile{OsDisk: &compute.OSDisk{Name: &OsDiskName}}
		vmProperties := compute.VirtualMachineProperties{StorageProfile: &storageProfile, NetworkProfile: &networkProfile}
		return &compute.VirtualMachine{Name: &name, VirtualMachineProperties: &vmProperties}, nil
	}}
}
func MockVMExistsNICInvalid() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockVMIfExists: func(resourceGroup string, name string) (*compute.VirtualMachine, error) {
		networkProfile := compute.NetworkProfile{NetworkInterfaces: &[]compute.NetworkInterfaceReference{{ID: to.StringPtr("")}}}
		OsDiskName := fmt.Sprintf("OS_Disk_%v", name)
		storageProfile := compute.StorageProfile{OsDisk: &compute.OSDisk{Name: &OsDiskName}}
		vmProperties := compute.VirtualMachineProperties{StorageProfile: &storageProfile, NetworkProfile: &networkProfile}
		return &compute.VirtualMachine{Name: &name, VirtualMachineProperties: &vmProperties}, nil
	}}
}
func MockVMNotExists() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockVMIfExists: func(resourceGroup string, name string) (*compute.VirtualMachine, error) {
		return nil, nil
	}}
}
func MockVMCheckFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockVMIfExists: func(resourceGroup string, name string) (*compute.VirtualMachine, error) {
		return &compute.VirtualMachine{}, errors.New("error while checking if vm exists")
	}}
}
func MockVMDeleteFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockDeleteVM: func(resourceGroup string, name string) (compute.VirtualMachinesDeleteFuture, error) {
		return compute.VirtualMachinesDeleteFuture{}, errors.New("error while deleting vm")
	}}
}
func MockVMDeleteFutureFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockWaitForVMDeletionFuture: func(future compute.VirtualMachinesDeleteFuture) error {
		return errors.New("failed on waiting for VirtualMachinesDeleteFuture")
	}}
}
func MockDisksDeleteFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockDeleteManagedDisk: func(resourceGroup string, name string) (compute.DisksDeleteFuture, error) {
		return compute.DisksDeleteFuture{}, errors.New("error while deleting managed disk")
	}}
}
func MockDisksDeleteFutureFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockWaitForDisksDeleteFuture: func(future compute.DisksDeleteFuture) error {
		return errors.New("failed on waiting for VirtualMachinesDeleteFuture")
	}}
}
func MockRunCommandFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockRunCommand: func(resourceGroup string, name string, cmd string) (compute.VirtualMachinesRunCommandFuture, error) {
		return compute.VirtualMachinesRunCommandFuture{}, errors.New("error while running command on vm")
	}}
}
func MockRunCommandFutureFailure() MockAzureComputeClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureComputeClient{MockWaitForVMRunCommandFuture: func(future compute.VirtualMachinesRunCommandFuture) error {
		return errors.New("failed on waiting for VirtualMachinesRunCommandFuture")
	}}
}
func MockNsgCreateOrUpdateSuccess() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockCreateOrUpdateNetworkSecurityGroup: func(resourceGroupName string, networkSecurityGroupName string, location string) (*network.SecurityGroupsCreateOrUpdateFuture, error) {
		return &network.SecurityGroupsCreateOrUpdateFuture{}, nil
	}}
}
func MockNsgCreateOrUpdateFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockCreateOrUpdateNetworkSecurityGroup: func(resourceGroupName string, networkSecurityGroupName string, location string) (*network.SecurityGroupsCreateOrUpdateFuture, error) {
		return nil, errors.New("failed to create or update network security group")
	}}
}
func MockVnetCreateOrUpdateSuccess() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockCreateOrUpdateVnet: func(resourceGroupName string, virtualNetworkName string, location string) (*network.VirtualNetworksCreateOrUpdateFuture, error) {
		return &network.VirtualNetworksCreateOrUpdateFuture{}, nil
	}}
}
func MockVnetCreateOrUpdateFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockCreateOrUpdateVnet: func(resourceGroupName string, virtualNetworkName string, location string) (*network.VirtualNetworksCreateOrUpdateFuture, error) {
		return nil, errors.New("failed to create or update vnet")
	}}
}
func MockNsgCreateOrUpdateFutureFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockWaitForNetworkSGsCreateOrUpdateFuture: func(future network.SecurityGroupsCreateOrUpdateFuture) error {
		return errors.New("failed on waiting for SecurityGroupsCreateOrUpdateFuture")
	}}
}
func MockVnetCreateOrUpdateFutureFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockWaitForVnetCreateOrUpdateFuture: func(future network.VirtualNetworksCreateOrUpdateFuture) error {
		return errors.New("failed on waiting for VirtualNetworksCreateOrUpdateFuture")
	}}
}
func MockNicDeleteFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockDeleteNetworkInterface: func(resourceGroup string, networkInterfaceName string) (network.InterfacesDeleteFuture, error) {
		return network.InterfacesDeleteFuture{}, errors.New("failed to delete network interface")
	}}
}
func MockNicDeleteFutureFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockWaitForNetworkInterfacesDeleteFuture: func(future network.InterfacesDeleteFuture) error {
		return errors.New("failed on waiting for InterfacesDeleteFuture")
	}}
}
func MockPublicIPDeleteFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockDeletePublicIPAddress: func(resourceGroup string, IPName string) (network.PublicIPAddressesDeleteFuture, error) {
		return network.PublicIPAddressesDeleteFuture{}, errors.New("failed to delete public ip address")
	}}
}
func MockPublicIPDeleteFutureFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockWaitForPublicIPAddressDeleteFuture: func(future network.PublicIPAddressesDeleteFuture) error {
		return errors.New("failed on waiting for PublicIPAddressesDeleteFuture")
	}}
}
func MockCreateOrUpdatePublicIPAddress(ip string) MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockCreateOrUpdatePublicIPAddress: func(resourceGroup string, IPName string) (network.PublicIPAddress, error) {
		publicIPAddress := network.PublicIPAddress{PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{}}
		publicIPAddress.IPAddress = to.StringPtr(ip)
		return publicIPAddress, nil
	}}
}
func MockCreateOrUpdatePublicIPAddressFailure() MockAzureNetworkClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureNetworkClient{MockCreateOrUpdatePublicIPAddress: func(resourceGroup string, IPName string) (network.PublicIPAddress, error) {
		return network.PublicIPAddress{}, errors.New("failed to get public ip address")
	}}
}
func MockRgExists() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockCheckGroupExistence: func(rgName string) (autorest.Response, error) {
		return autorest.Response{Response: &http.Response{StatusCode: 200}}, nil
	}}
}
func MockRgNotExists() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockCheckGroupExistence: func(rgName string) (autorest.Response, error) {
		return autorest.Response{Response: &http.Response{StatusCode: 404}}, nil
	}}
}
func MockRgCheckFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockCheckGroupExistence: func(rgName string) (autorest.Response, error) {
		return autorest.Response{Response: &http.Response{StatusCode: 200}}, errors.New("failed to check resource group existence")
	}}
}
func MockRgCreateOrUpdateFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockCreateOrUpdateGroup: func(resourceGroupName string, location string) (resources.Group, error) {
		return resources.Group{}, errors.New("failed to create resource group")
	}}
}
func MockRgDeleteSuccess() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockDeleteGroup: func(resourceGroupName string) (resources.GroupsDeleteFuture, error) {
		return resources.GroupsDeleteFuture{}, nil
	}}
}
func MockRgDeleteFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockDeleteGroup: func(resourceGroupName string) (resources.GroupsDeleteFuture, error) {
		return resources.GroupsDeleteFuture{}, errors.New("failed to delete resource group")
	}}
}
func MockRgDeleteFutureFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockWaitForGroupsDeleteFuture: func(future resources.GroupsDeleteFuture) error {
		return errors.New("error waiting for GroupsDeleteFuture")
	}}
}
func MockDeploymentCreateOrUpdateSuccess() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockCreateOrUpdateDeployment: func(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) (*resources.DeploymentsCreateOrUpdateFuture, error) {
		return &resources.DeploymentsCreateOrUpdateFuture{}, nil
	}}
}
func MockDeploymentCreateOrUpdateFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockCreateOrUpdateDeployment: func(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) (*resources.DeploymentsCreateOrUpdateFuture, error) {
		return nil, errors.New("failed to create resource")
	}}
}
func MockDeploymentCreateOrUpdateFutureFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockWaitForDeploymentsCreateOrUpdateFuture: func(future resources.DeploymentsCreateOrUpdateFuture) error {
		return errors.New("failed on waiting for DeploymentsCreateOrUpdateFuture")
	}}
}
func MockDeloymentGetResultSuccess() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockGetDeploymentResult: func(future resources.DeploymentsCreateOrUpdateFuture) (resources.DeploymentExtended, error) {
		return resources.DeploymentExtended{Name: to.StringPtr("deployment-test")}, nil
	}}
}
func MockDeloymentGetResultFailure() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockGetDeploymentResult: func(future resources.DeploymentsCreateOrUpdateFuture) (resources.DeploymentExtended, error) {
		return resources.DeploymentExtended{}, errors.New("error getting deployment result")
	}}
}
func MockDeploymentValidate() MockAzureResourcesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return MockAzureResourcesClient{MockValidateDeployment: func(machine *clusterv1.Machine, clusterConfig *providerv1.AzureClusterProviderSpec, machineConfig *providerv1.AzureMachineProviderSpec) error {
		return errors.New("error validating deployment")
	}}
}
