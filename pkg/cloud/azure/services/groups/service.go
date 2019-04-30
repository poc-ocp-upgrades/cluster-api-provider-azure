package groups

import (
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	resources.GroupsClient
	Scope	*actuators.Scope
}

func getGroupsClient(subscriptionID string, authorizer autorest.Authorizer) resources.GroupsClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupsClient := resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = authorizer
	groupsClient.AddToUserAgent(azure.UserAgent)
	return groupsClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getGroupsClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
