package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AzureClusterProviderSpec) DeepCopyInto(out *AzureClusterProviderSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.NetworkSpec.DeepCopyInto(&out.NetworkSpec)
	in.CAKeyPair.DeepCopyInto(&out.CAKeyPair)
	in.EtcdCAKeyPair.DeepCopyInto(&out.EtcdCAKeyPair)
	in.FrontProxyCAKeyPair.DeepCopyInto(&out.FrontProxyCAKeyPair)
	in.SAKeyPair.DeepCopyInto(&out.SAKeyPair)
	if in.DiscoveryHashes != nil {
		in, out := &in.DiscoveryHashes, &out.DiscoveryHashes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ClusterConfiguration.DeepCopyInto(&out.ClusterConfiguration)
	return
}
func (in *AzureClusterProviderSpec) DeepCopy() *AzureClusterProviderSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AzureClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *AzureClusterProviderSpec) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AzureClusterProviderStatus) DeepCopyInto(out *AzureClusterProviderStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Network.DeepCopyInto(&out.Network)
	out.Bastion = in.Bastion
	return
}
func (in *AzureClusterProviderStatus) DeepCopy() *AzureClusterProviderStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AzureClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *AzureClusterProviderStatus) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AzureMachineProviderCondition) DeepCopyInto(out *AzureMachineProviderCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *AzureMachineProviderCondition) DeepCopy() *AzureMachineProviderCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AzureMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *AzureMachineProviderSpec) DeepCopyInto(out *AzureMachineProviderSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	out.Image = in.Image
	out.OSDisk = in.OSDisk
	return
}
func (in *AzureMachineProviderSpec) DeepCopy() *AzureMachineProviderSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AzureMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *AzureMachineProviderSpec) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AzureMachineProviderStatus) DeepCopyInto(out *AzureMachineProviderStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.VMID != nil {
		in, out := &in.VMID, &out.VMID
		*out = new(string)
		**out = **in
	}
	if in.VMState != nil {
		in, out := &in.VMState, &out.VMState
		*out = new(VMState)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AzureMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *AzureMachineProviderStatus) DeepCopy() *AzureMachineProviderStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AzureMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *AzureMachineProviderStatus) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *AzureResourceReference) DeepCopyInto(out *AzureResourceReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	return
}
func (in *AzureResourceReference) DeepCopy() *AzureResourceReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AzureResourceReference)
	in.DeepCopyInto(out)
	return out
}
func (in *BackendPool) DeepCopyInto(out *BackendPool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *BackendPool) DeepCopy() *BackendPool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BackendPool)
	in.DeepCopyInto(out)
	return out
}
func (in *FrontendIPConfig) DeepCopyInto(out *FrontendIPConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *FrontendIPConfig) DeepCopy() *FrontendIPConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(FrontendIPConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *Image) DeepCopyInto(out *Image) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *Image) DeepCopy() *Image {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SourcePorts != nil {
		in, out := &in.SourcePorts, &out.SourcePorts
		*out = new(string)
		**out = **in
	}
	if in.DestinationPorts != nil {
		in, out := &in.DestinationPorts, &out.DestinationPorts
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(string)
		**out = **in
	}
	return
}
func (in *IngressRule) DeepCopy() *IngressRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}
func (in IngressRules) DeepCopyInto(out *IngressRules) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	{
		in := &in
		*out = make(IngressRules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IngressRule)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}
func (in IngressRules) DeepCopy() IngressRules {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(IngressRules)
	in.DeepCopyInto(out)
	return *out
}
func (in *KeyPair) DeepCopyInto(out *KeyPair) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *KeyPair) DeepCopy() *KeyPair {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KeyPair)
	in.DeepCopyInto(out)
	return out
}
func (in *KubeadmConfiguration) DeepCopyInto(out *KubeadmConfiguration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Join.DeepCopyInto(&out.Join)
	in.Init.DeepCopyInto(&out.Init)
	return
}
func (in *KubeadmConfiguration) DeepCopy() *KubeadmConfiguration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(KubeadmConfiguration)
	in.DeepCopyInto(out)
	return out
}
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.FrontendIPConfig = in.FrontendIPConfig
	out.BackendPool = in.BackendPool
	return
}
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}
func (in *LoadBalancerHealthCheck) DeepCopyInto(out *LoadBalancerHealthCheck) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *LoadBalancerHealthCheck) DeepCopy() *LoadBalancerHealthCheck {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LoadBalancerHealthCheck)
	in.DeepCopyInto(out)
	return out
}
func (in *LoadBalancerListener) DeepCopyInto(out *LoadBalancerListener) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *LoadBalancerListener) DeepCopy() *LoadBalancerListener {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListener)
	in.DeepCopyInto(out)
	return out
}
func (in *ManagedDisk) DeepCopyInto(out *ManagedDisk) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ManagedDisk) DeepCopy() *ManagedDisk {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ManagedDisk)
	in.DeepCopyInto(out)
	return out
}
func (in *Network) DeepCopyInto(out *Network) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make(map[SecurityGroupRole]*SecurityGroup, len(*in))
		for key, val := range *in {
			var outVal *SecurityGroup
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(SecurityGroup)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	out.APIServerLB = in.APIServerLB
	out.APIServerIP = in.APIServerIP
	return
}
func (in *Network) DeepCopy() *Network {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Vnet = in.Vnet
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *OSDisk) DeepCopyInto(out *OSDisk) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.ManagedDisk = in.ManagedDisk
	return
}
func (in *OSDisk) DeepCopy() *OSDisk {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OSDisk)
	in.DeepCopyInto(out)
	return out
}
func (in *PublicIP) DeepCopyInto(out *PublicIP) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *PublicIP) DeepCopy() *PublicIP {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PublicIP)
	in.DeepCopyInto(out)
	return out
}
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.IngressRules != nil {
		in, out := &in.IngressRules, &out.IngressRules
		*out = make(IngressRules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(IngressRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.SecurityGroup.DeepCopyInto(&out.SecurityGroup)
	return
}
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}
func (in Subnets) DeepCopyInto(out *Subnets) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	{
		in := &in
		*out = make(Subnets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetSpec)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}
func (in Subnets) DeepCopy() Subnets {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}
func (in *VM) DeepCopyInto(out *VM) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Image = in.Image
	out.OSDisk = in.OSDisk
	return
}
func (in *VM) DeepCopy() *VM {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(VM)
	in.DeepCopyInto(out)
	return out
}
func (in *VnetSpec) DeepCopyInto(out *VnetSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *VnetSpec) DeepCopy() *VnetSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(VnetSpec)
	in.DeepCopyInto(out)
	return out
}
