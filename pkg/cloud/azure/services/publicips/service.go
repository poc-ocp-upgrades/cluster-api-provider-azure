package publicips

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	network.PublicIPAddressesClient
	Scope	*actuators.Scope
}

func getPublicIPAddressesClient(subscriptionID string, authorizer autorest.Authorizer) network.PublicIPAddressesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	publicIPsClient := network.NewPublicIPAddressesClient(subscriptionID)
	publicIPsClient.Authorizer = authorizer
	publicIPsClient.AddToUserAgent(azure.UserAgent)
	return publicIPsClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getPublicIPAddressesClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
