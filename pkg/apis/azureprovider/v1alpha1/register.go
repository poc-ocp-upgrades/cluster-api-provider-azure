package v1alpha1

import (
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
	"sigs.k8s.io/yaml"
)

var (
	SchemeGroupVersion	= schema.GroupVersion{Group: "azureprovider.k8s.io", Version: "v1alpha1"}
	SchemeBuilder		= &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

func ClusterConfigFromProviderSpec(providerConfig clusterv1.ProviderSpec) (*AzureClusterProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var config AzureClusterProviderSpec
	if providerConfig.Value == nil {
		return &config, nil
	}
	if err := yaml.Unmarshal(providerConfig.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
func ClusterStatusFromProviderStatus(extension *runtime.RawExtension) (*AzureClusterProviderStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if extension == nil {
		return &AzureClusterProviderStatus{}, nil
	}
	status := new(AzureClusterProviderStatus)
	if err := yaml.Unmarshal(extension.Raw, status); err != nil {
		return nil, err
	}
	return status, nil
}
func MachineStatusFromProviderStatus(extension *runtime.RawExtension) (*AzureMachineProviderStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if extension == nil {
		return &AzureMachineProviderStatus{}, nil
	}
	status := new(AzureMachineProviderStatus)
	if err := yaml.Unmarshal(extension.Raw, status); err != nil {
		return nil, err
	}
	return status, nil
}
func EncodeMachineStatus(status *AzureMachineProviderStatus) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if status == nil {
		return &runtime.RawExtension{}, nil
	}
	var rawBytes []byte
	var err error
	if rawBytes, err = json.Marshal(status); err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: rawBytes}, nil
}
func EncodeMachineSpec(spec *AzureMachineProviderSpec) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}
	var rawBytes []byte
	var err error
	if rawBytes, err = json.Marshal(spec); err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: rawBytes}, nil
}
func EncodeClusterStatus(status *AzureClusterProviderStatus) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if status == nil {
		return &runtime.RawExtension{}, nil
	}
	var rawBytes []byte
	var err error
	if rawBytes, err = json.Marshal(status); err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: rawBytes}, nil
}
func EncodeClusterSpec(spec *AzureClusterProviderSpec) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec == nil {
		return &runtime.RawExtension{}, nil
	}
	var rawBytes []byte
	var err error
	if rawBytes, err = json.Marshal(spec); err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: rawBytes}, nil
}
