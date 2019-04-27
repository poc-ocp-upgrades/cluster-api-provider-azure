package availabilityzones

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	compute.ResourceSkusClient
	Scope	*actuators.Scope
}

func getResourceSkusClient(subscriptionID string, authorizer autorest.Authorizer) compute.ResourceSkusClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	skusClient := compute.NewResourceSkusClient(subscriptionID)
	skusClient.Authorizer = authorizer
	skusClient.AddToUserAgent(azure.UserAgent)
	return skusClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getResourceSkusClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
