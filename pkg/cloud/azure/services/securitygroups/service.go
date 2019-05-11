package securitygroups

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	network.SecurityGroupsClient
	Scope	*actuators.Scope
}

func getSecurityGroupsClient(subscriptionID string, authorizer autorest.Authorizer) network.SecurityGroupsClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	securityGroupsClient := network.NewSecurityGroupsClient(subscriptionID)
	securityGroupsClient.Authorizer = authorizer
	securityGroupsClient.AddToUserAgent(azure.UserAgent)
	return securityGroupsClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getSecurityGroupsClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
