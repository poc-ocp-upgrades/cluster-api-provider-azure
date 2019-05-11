package routetables

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	network.RouteTablesClient
	Scope	*actuators.Scope
}

func getRouteTablesClient(subscriptionID string, authorizer autorest.Authorizer) network.RouteTablesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	routeTablesClient := network.NewRouteTablesClient(subscriptionID)
	routeTablesClient.Authorizer = authorizer
	routeTablesClient.AddToUserAgent(azure.UserAgent)
	return routeTablesClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getRouteTablesClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
