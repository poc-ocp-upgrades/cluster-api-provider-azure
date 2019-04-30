package networkinterfaces

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	network.InterfacesClient
	Scope	*actuators.Scope
}

func getNetworkInterfacesClient(subscriptionID string, authorizer autorest.Authorizer) network.InterfacesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nicClient := network.NewInterfacesClient(subscriptionID)
	nicClient.Authorizer = authorizer
	nicClient.AddToUserAgent(azure.UserAgent)
	return nicClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getNetworkInterfacesClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
