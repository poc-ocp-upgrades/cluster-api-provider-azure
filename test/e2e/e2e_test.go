package e2e

import (
	"fmt"
	"os"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators/machine"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/resources"
)

type Clients struct {
	kube	KubeClient
	azure	actuators.AzureClients
}

func TestMasterMachineCreated(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kubeConfig := os.Getenv("KUBE_CONFIG")
	if kubeConfig == "" {
		t.Skip("KUBE_CONFIG environment variable is not set")
	}
	clients, err := createTestClients(kubeConfig)
	if err != nil {
		t.Fatalf("failed to create test clients: %v", err)
	}
	machineList, err := clients.kube.ListMachine("default", metav1.ListOptions{LabelSelector: "set=master"})
	if err != nil {
		t.Fatalf("error to while trying to retrieve machine list: %v", err)
	}
	if len(machineList.Items) != 1 {
		t.Fatalf("expected only one machine with label master in the default namespace")
	}
	masterMachine := machineList.Items[0]
	resourceGroup := masterMachine.ObjectMeta.Annotations[string(machine.ResourceGroup)]
	vm, err := clients.azure.Compute.VMIfExists(resourceGroup, resources.GetVMName(&masterMachine))
	if err != nil {
		t.Fatalf("error checking if vm exists: %v", err)
	}
	if vm == nil {
		t.Fatalf("couldn't find vm for machine: %v", masterMachine.Name)
	}
}
func createTestClients(kubeConfig string) (*Clients, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kubeClient, err := NewKubeClient(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kubernetes client: %v", err)
	}
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return nil, fmt.Errorf("AZURE_SUBSCRIPTION_ID environment variable is not set")
	}
	azureServicesClient, err := NewAzureServicesClient(subscriptionID)
	if err != nil {
		return nil, fmt.Errorf("failed to create azure services client: %v", err)
	}
	return &Clients{kube: *kubeClient, azure: *azureServicesClient}, nil
}
