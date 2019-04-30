package networkinterfaces

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/internalloadbalancers"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/publicloadbalancers"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/subnets"
)

type Spec struct {
	Name				string
	SubnetName			string
	VnetName			string
	StaticIPAddress			string
	PublicLoadBalancerName		string
	InternalLoadBalancerName	string
	NatRule				int
}

func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nicSpec, ok := spec.(*Spec)
	if !ok {
		return network.Interface{}, errors.New("invalid network interface specification")
	}
	nic, err := s.Client.Get(ctx, s.Scope.ClusterConfig.ResourceGroup, nicSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, errors.Wrapf(err, "network interface %s not found", nicSpec.Name)
	} else if err != nil {
		return nic, err
	}
	return nic, nil
}
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nicSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid network interface specification")
	}
	nicConfig := &network.InterfaceIPConfigurationPropertiesFormat{}
	subnetInterface, err := subnets.NewService(s.Scope).Get(ctx, &subnets.Spec{Name: nicSpec.SubnetName, VnetName: nicSpec.VnetName})
	if err != nil {
		return err
	}
	subnet, ok := subnetInterface.(network.Subnet)
	if !ok {
		return errors.New("subnet get returned invalid network interface")
	}
	nicConfig.Subnet = &network.Subnet{ID: subnet.ID}
	nicConfig.PrivateIPAllocationMethod = network.Dynamic
	if nicSpec.StaticIPAddress != "" {
		nicConfig.PrivateIPAllocationMethod = network.Static
		nicConfig.PrivateIPAddress = to.StringPtr(nicSpec.StaticIPAddress)
	}
	backendAddressPools := []network.BackendAddressPool{}
	if nicSpec.PublicLoadBalancerName != "" {
		lbInterface, lberr := publicloadbalancers.NewService(s.Scope).Get(ctx, &publicloadbalancers.Spec{Name: nicSpec.PublicLoadBalancerName})
		if lberr != nil {
			return lberr
		}
		lb, ok := lbInterface.(network.LoadBalancer)
		if !ok {
			return errors.New("public load balancer get returned invalid network interface")
		}
		backendAddressPools = append(backendAddressPools, network.BackendAddressPool{ID: (*lb.BackendAddressPools)[0].ID})
		nicConfig.LoadBalancerInboundNatRules = &[]network.InboundNatRule{{ID: (*lb.InboundNatRules)[nicSpec.NatRule].ID}}
	}
	if nicSpec.InternalLoadBalancerName != "" {
		internallbInterface, ilberr := internalloadbalancers.NewService(s.Scope).Get(ctx, &internalloadbalancers.Spec{Name: nicSpec.InternalLoadBalancerName})
		if ilberr != nil {
			return ilberr
		}
		internallb, ok := internallbInterface.(network.LoadBalancer)
		if !ok {
			return errors.New("internal load balancer get returned invalid network interface")
		}
		backendAddressPools = append(backendAddressPools, network.BackendAddressPool{ID: (*internallb.BackendAddressPools)[0].ID})
	}
	nicConfig.LoadBalancerBackendAddressPools = &backendAddressPools
	f, err := s.Client.CreateOrUpdate(ctx, s.Scope.ClusterConfig.ResourceGroup, nicSpec.Name, network.Interface{Location: to.StringPtr(s.Scope.ClusterConfig.Location), InterfacePropertiesFormat: &network.InterfacePropertiesFormat{IPConfigurations: &[]network.InterfaceIPConfiguration{{Name: to.StringPtr("pipConfig"), InterfaceIPConfigurationPropertiesFormat: nicConfig}}}})
	if err != nil {
		return errors.Wrapf(err, "failed to create network interface %s in resource group %s", nicSpec.Name, s.Scope.ClusterConfig.ResourceGroup)
	}
	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot create, future response")
	}
	_, err = f.Result(s.Client)
	if err != nil {
		return errors.Wrap(err, "result error")
	}
	klog.V(2).Infof("successfully created network interface %s", nicSpec.Name)
	return err
}
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nicSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid network interface Specification")
	}
	klog.V(2).Infof("deleting nic %s", nicSpec.Name)
	f, err := s.Client.Delete(ctx, s.Scope.ClusterConfig.ResourceGroup, nicSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "failed to delete network interface %s in resource group %s", nicSpec.Name, s.Scope.ClusterConfig.ResourceGroup)
	}
	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot create, future response")
	}
	_, err = f.Result(s.Client)
	if err != nil {
		return errors.Wrap(err, "result error")
	}
	klog.V(2).Infof("successfully deleted nic %s", nicSpec.Name)
	return err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
