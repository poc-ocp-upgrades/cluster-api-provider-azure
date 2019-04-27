package virtualmachines

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/networkinterfaces"
)

type Spec struct {
	Name		string
	NICName		string
	SSHKeyData	string
	Size		string
	Zone		string
	Image		v1alpha1.Image
	OSDisk		v1alpha1.OSDisk
	CustomData	string
}

func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmSpec, ok := spec.(*Spec)
	if !ok {
		return compute.VirtualMachine{}, errors.New("invalid vm specification")
	}
	vm, err := s.Client.Get(ctx, s.Scope.ClusterConfig.ResourceGroup, vmSpec.Name, "")
	if err != nil && azure.ResourceNotFound(err) {
		return nil, errors.Wrapf(err, "vm %s not found", vmSpec.Name)
	} else if err != nil {
		return vm, err
	}
	return vm, nil
}
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid vm specification")
	}
	klog.V(2).Infof("getting nic %s", vmSpec.NICName)
	nicInterface, err := networkinterfaces.NewService(s.Scope).Get(ctx, &networkinterfaces.Spec{Name: vmSpec.NICName})
	if err != nil {
		return err
	}
	nic, ok := nicInterface.(network.Interface)
	if !ok {
		return errors.New("error getting network security group")
	}
	klog.V(2).Infof("got nic %s", vmSpec.NICName)
	klog.V(2).Infof("creating vm %s ", vmSpec.Name)
	sshKeyData := vmSpec.SSHKeyData
	if sshKeyData == "" {
		privateKey, perr := rsa.GenerateKey(rand.Reader, 2048)
		if perr != nil {
			return errors.Wrap(perr, "Failed to generate private key")
		}
		publicRsaKey, perr := ssh.NewPublicKey(&privateKey.PublicKey)
		if perr != nil {
			return errors.Wrap(perr, "Failed to generate public key")
		}
		sshKeyData = string(ssh.MarshalAuthorizedKey(publicRsaKey))
	}
	randomPassword, err := GenerateRandomString(32)
	if err != nil {
		return errors.Wrapf(err, "failed to generate random string")
	}
	imageReference := &compute.ImageReference{Publisher: to.StringPtr(vmSpec.Image.Publisher), Offer: to.StringPtr(vmSpec.Image.Offer), Sku: to.StringPtr(vmSpec.Image.SKU), Version: to.StringPtr(vmSpec.Image.Version)}
	if vmSpec.Image.ResourceID != "" {
		imageReference = &compute.ImageReference{ID: to.StringPtr(fmt.Sprintf("/subscriptions/%s%s", s.Scope.SubscriptionID, vmSpec.Image.ResourceID))}
	}
	osProfile := &compute.OSProfile{ComputerName: to.StringPtr(vmSpec.Name), AdminUsername: to.StringPtr(azure.DefaultUserName), AdminPassword: to.StringPtr(randomPassword)}
	if sshKeyData != "" {
		osProfile.LinuxConfiguration = &compute.LinuxConfiguration{SSH: &compute.SSHConfiguration{PublicKeys: &[]compute.SSHPublicKey{{Path: to.StringPtr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", azure.DefaultUserName)), KeyData: to.StringPtr(sshKeyData)}}}}
	}
	if vmSpec.CustomData != "" {
		osProfile.CustomData = to.StringPtr(vmSpec.CustomData)
	}
	virtualMachine := compute.VirtualMachine{Location: to.StringPtr(s.Scope.ClusterConfig.Location), VirtualMachineProperties: &compute.VirtualMachineProperties{HardwareProfile: &compute.HardwareProfile{VMSize: compute.VirtualMachineSizeTypes(vmSpec.Size)}, StorageProfile: &compute.StorageProfile{ImageReference: imageReference, OsDisk: &compute.OSDisk{Name: to.StringPtr(fmt.Sprintf("%s_OSDisk", vmSpec.Name)), OsType: compute.OperatingSystemTypes(vmSpec.OSDisk.OSType), CreateOption: compute.DiskCreateOptionTypesFromImage, DiskSizeGB: to.Int32Ptr(vmSpec.OSDisk.DiskSizeGB), ManagedDisk: &compute.ManagedDiskParameters{StorageAccountType: compute.StorageAccountTypes(vmSpec.OSDisk.ManagedDisk.StorageAccountType)}}}, OsProfile: osProfile, NetworkProfile: &compute.NetworkProfile{NetworkInterfaces: &[]compute.NetworkInterfaceReference{{ID: nic.ID, NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{Primary: to.BoolPtr(true)}}}}}}
	if vmSpec.Zone != "" {
		zones := []string{vmSpec.Zone}
		virtualMachine.Zones = &zones
	}
	future, err := s.Client.CreateOrUpdate(ctx, s.Scope.ClusterConfig.ResourceGroup, vmSpec.Name, virtualMachine)
	if err != nil {
		return errors.Wrapf(err, "cannot create vm")
	}
	err = future.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrapf(err, "cannot get the vm create or update future response")
	}
	_, err = future.Result(s.Client)
	if err != nil {
		return err
	}
	klog.V(2).Infof("successfully created vm %s ", vmSpec.Name)
	return err
}
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vmSpec, ok := spec.(*Spec)
	if !ok {
		return errors.New("invalid vm Specification")
	}
	klog.V(2).Infof("deleting vm %s ", vmSpec.Name)
	future, err := s.Client.Delete(ctx, s.Scope.ClusterConfig.ResourceGroup, vmSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "failed to delete vm %s in resource group %s", vmSpec.Name, s.Scope.ClusterConfig.ResourceGroup)
	}
	err = future.WaitForCompletionRef(ctx, s.Client.Client)
	if err != nil {
		return errors.Wrap(err, "cannot delete, future response")
	}
	_, err = future.Result(s.Client)
	klog.V(2).Infof("successfully deleted vm %s ", vmSpec.Name)
	return err
}
func GenerateRandomString(n int) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
