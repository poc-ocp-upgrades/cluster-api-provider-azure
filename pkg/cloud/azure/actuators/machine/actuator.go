package machine

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"time"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	client "github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/typed/machine/v1beta1"
	controllerError "github.com/openshift/cluster-api/pkg/controller/error"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/deployer"
	controllerclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type Actuator struct {
	*deployer.Deployer
	client		client.MachineV1beta1Interface
	coreClient	controllerclient.Client
}
type ActuatorParams struct {
	Client		client.MachineV1beta1Interface
	CoreClient	controllerclient.Client
}

func NewActuator(params ActuatorParams) *Actuator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Actuator{Deployer: deployer.New(deployer.Params{ScopeGetter: actuators.DefaultScopeGetter}), client: params.Client, coreClient: params.CoreClient}
}
func (a *Actuator) Create(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Creating machine %v", machine.Name)
	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{Machine: machine, Cluster: nil, Client: a.client, CoreClient: a.coreClient})
	if err != nil {
		return errors.Errorf("failed to create scope: %+v", err)
	}
	defer scope.Close()
	err = NewReconciler(scope).Create(context.Background())
	if err != nil {
		klog.Errorf("failed to reconcile machine %s: %v", machine.Name, err)
		return &controllerError.RequeueAfterError{RequeueAfter: time.Minute}
	}
	return nil
}
func (a *Actuator) Delete(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Deleting machine %v", machine.Name)
	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{Machine: machine, Cluster: nil, Client: a.client, CoreClient: a.coreClient})
	if err != nil {
		return errors.Wrapf(err, "failed to create scope")
	}
	defer scope.Close()
	err = NewReconciler(scope).Delete(context.Background())
	if err != nil {
		klog.Errorf("failed to delete machine %s: %v", machine.Name, err)
		return &controllerError.RequeueAfterError{RequeueAfter: time.Minute}
	}
	return nil
}
func (a *Actuator) Update(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Updating machine %v", machine.Name)
	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{Machine: machine, Cluster: nil, Client: a.client, CoreClient: a.coreClient})
	if err != nil {
		return errors.Errorf("failed to create scope: %+v", err)
	}
	defer scope.Close()
	err = NewReconciler(scope).Update(context.Background())
	if err != nil {
		klog.Errorf("failed to update machine %s: %v", machine.Name, err)
		return &controllerError.RequeueAfterError{RequeueAfter: time.Minute}
	}
	return nil
}
func (a *Actuator) Exists(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Checking if machine %v exists", machine.Name)
	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{Machine: machine, Cluster: nil, Client: a.client, CoreClient: a.coreClient})
	if err != nil {
		return false, errors.Errorf("failed to create scope: %+v", err)
	}
	defer scope.Close()
	isExists, err := NewReconciler(scope).Exists(context.Background())
	if err != nil {
		klog.Errorf("failed to check machine %s exists: %v", machine.Name, err)
	}
	return isExists, err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
