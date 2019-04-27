package virtualmachineextensions

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

type Spec struct {
	Name		string
	VMName		string
	ScriptData	string
}

func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmExtSpec, ok := spec.(*Spec)
	if !ok {
		return compute.VirtualMachineExtension{}, errors.New("invalid vm specification")
	}
	vmExt, err := s.Client.Get(ctx, s.Scope.ClusterConfig.ResourceGroup, vmExtSpec.VMName, vmExtSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, errors.Wrapf(err, "vm extension %s not found", vmExtSpec.Name)
	} else if err != nil {
		return vmExt, err
	}
	return vmExt, nil
}
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmExtSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid vm specification")
	}
	klog.V(2).Infof("creating vm extension %s ", vmExtSpec.Name)
	future, err := s.Client.CreateOrUpdate(ctx, s.Scope.ClusterConfig.ResourceGroup, vmExtSpec.VMName, vmExtSpec.Name, compute.VirtualMachineExtension{Name: to.StringPtr(vmExtSpec.Name), Location: to.StringPtr(s.Scope.ClusterConfig.Location), VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{Type: to.StringPtr("CustomScript"), TypeHandlerVersion: to.StringPtr("2.0"), AutoUpgradeMinorVersion: to.BoolPtr(true), Settings: map[string]bool{"skipDos2Unix": true}, Publisher: to.StringPtr("Microsoft.Azure.Extensions"), ProtectedSettings: map[string]string{"script": vmExtSpec.ScriptData}}})
	if err != nil {
		return errors.Wrapf(err, "cannot create vm extension")
	}
	err = future.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrapf(err, "cannot get the extension create or update future response")
	}
	_, err = future.Result(s.Client)
	if err != nil {
		return errors.Wrapf(err, "cannot create vm")
	}
	klog.V(2).Infof("successfully created vm extension %s ", vmExtSpec.Name)
	return err
}
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmExtSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("Invalid VNET Specification")
	}
	klog.V(2).Infof("deleting vm extension %s ", vmExtSpec.Name)
	future, err := s.Client.Delete(ctx, s.Scope.ClusterConfig.ResourceGroup, vmExtSpec.VMName, vmExtSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "failed to delete vm extension %s in resource group %s", vmExtSpec.Name, s.Scope.ClusterConfig.ResourceGroup)
	}
	err = future.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot delete, future response")
	}
	_, err = future.Result(s.Client)
	klog.V(2).Infof("successfully deleted vm %s ", vmExtSpec.Name)
	return err
}
