package internalloadbalancers

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	network.LoadBalancersClient
	Scope	*actuators.Scope
}

func getLoadbalancersClient(subscriptionID string, authorizer autorest.Authorizer) network.LoadBalancersClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	loadBalancersClient := network.NewLoadBalancersClient(subscriptionID)
	loadBalancersClient.Authorizer = authorizer
	loadBalancersClient.AddToUserAgent(azure.UserAgent)
	return loadBalancersClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getLoadbalancersClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
