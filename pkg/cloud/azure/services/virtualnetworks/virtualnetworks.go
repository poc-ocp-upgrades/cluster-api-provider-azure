package virtualnetworks

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

type Spec struct {
	Name	string
	CIDR	string
}

func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vnetSpec, ok := spec.(*Spec)
	if !ok {
		return network.VirtualNetwork{}, errors.New("Invalid VNET Specification")
	}
	vnet, err := s.Client.Get(ctx, s.Scope.ClusterConfig.ResourceGroup, vnetSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, errors.Wrapf(err, "vnet %s not found", vnetSpec.Name)
	} else if err != nil {
		return vnet, err
	}
	return vnet, nil
}
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vnetSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("Invalid VNET Specification")
	}
	if _, err := s.Get(ctx, vnetSpec); err == nil {
		return nil
	}
	klog.V(2).Infof("creating vnet %s ", vnetSpec.Name)
	f, err := s.Client.CreateOrUpdate(ctx, s.Scope.ClusterConfig.ResourceGroup, vnetSpec.Name, network.VirtualNetwork{Location: to.StringPtr(s.Scope.ClusterConfig.Location), VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{AddressSpace: &network.AddressSpace{AddressPrefixes: &[]string{vnetSpec.CIDR}}}})
	if err != nil {
		return err
	}
	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return err
	}
	_, err = f.Result(s.Client)
	if err != nil {
		return err
	}
	klog.V(2).Infof("successfully created vnet %s ", vnetSpec.Name)
	return err
}
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vnetSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("Invalid VNET Specification")
	}
	klog.V(2).Infof("deleting vnet %s ", vnetSpec.Name)
	future, err := s.Client.Delete(ctx, s.Scope.ClusterConfig.ResourceGroup, vnetSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "failed to delete vnet %s in resource group %s", vnetSpec.Name, s.Scope.ClusterConfig.ResourceGroup)
	}
	err = future.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot delete, future response")
	}
	_, err = future.Result(s.Client)
	klog.V(2).Infof("successfully deleted vnet %s ", vnetSpec.Name)
	return err
}
