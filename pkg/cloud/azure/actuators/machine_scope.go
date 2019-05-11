package actuators

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	machineclient "github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/typed/machine/v1beta1"
	"github.com/pkg/errors"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	controllerclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	AzureCredsSubscriptionIDKey	= "azure_subscription_id"
	AzureCredsClientIDKey		= "azure_client_id"
	AzureCredsClientSecretKey	= "azure_client_secret"
	AzureCredsTenantIDKey		= "azure_tenant_id"
	AzureCredsResourceGroupKey	= "azure_resourcegroup"
	AzureCredsRegionKey			= "azure_region"
	AzureResourcePrefix			= "azure_resource_prefix"
)

type MachineScopeParams struct {
	AzureClients
	Cluster		*clusterv1.Cluster
	Machine		*machinev1.Machine
	Client		machineclient.MachineV1beta1Interface
	CoreClient	controllerclient.Client
}

func NewMachineScope(params MachineScopeParams) (*MachineScope, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope, err := NewScope(ScopeParams{AzureClients: params.AzureClients, Client: nil, Cluster: params.Cluster})
	if err != nil {
		return nil, err
	}
	machineConfig, err := MachineConfigFromProviderSpec(params.Client, params.Machine.Spec.ProviderSpec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get machine config")
	}
	machineStatus, err := v1alpha1.MachineStatusFromProviderStatus(params.Machine.Status.ProviderStatus)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get machine provider status")
	}
	machineClient := params.Client.Machines(params.Machine.Namespace)
	if machineConfig.CredentialsSecret != nil {
		if err = updateScope(params.CoreClient, machineConfig.CredentialsSecret, scope); err != nil {
			return nil, errors.Wrap(err, "failed to update cluster")
		}
	}
	return &MachineScope{Scope: scope, Machine: params.Machine, MachineClient: machineClient, CoreClient: params.CoreClient, MachineConfig: machineConfig, MachineStatus: machineStatus}, nil
}

type MachineScope struct {
	*Scope
	Machine			*machinev1.Machine
	MachineClient	machineclient.MachineInterface
	CoreClient		controllerclient.Client
	MachineConfig	*v1alpha1.AzureMachineProviderSpec
	MachineStatus	*v1alpha1.AzureMachineProviderStatus
}

func (m *MachineScope) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Machine.Name
}
func (m *MachineScope) Namespace() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Machine.Namespace
}
func (m *MachineScope) Role() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Machine.Labels[v1alpha1.MachineRoleLabel]
}
func (m *MachineScope) Location() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Scope.Location()
}
func (m *MachineScope) storeMachineSpec(machine *machinev1.Machine) (*machinev1.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ext, err := v1alpha1.EncodeMachineSpec(m.MachineConfig)
	if err != nil {
		return nil, err
	}
	machine.Spec.ProviderSpec.Value = ext
	return m.MachineClient.Update(machine)
}
func (m *MachineScope) storeMachineStatus(machine *machinev1.Machine) (*machinev1.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ext, err := v1alpha1.EncodeMachineStatus(m.MachineStatus)
	if err != nil {
		return nil, err
	}
	m.Machine.Status.DeepCopyInto(&machine.Status)
	machine.Status.ProviderStatus = ext
	return m.MachineClient.UpdateStatus(machine)
}
func (m *MachineScope) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.MachineClient == nil {
		return
	}
	latestMachine, err := m.storeMachineSpec(m.Machine)
	if err != nil {
		klog.Errorf("[machinescope] failed to update machine %q in namespace %q: %v", m.Machine.Name, m.Machine.Namespace, err)
		return
	}
	_, err = m.storeMachineStatus(latestMachine)
	if err != nil {
		klog.Errorf("[machinescope] failed to store provider status for machine %q in namespace %q: %v", m.Machine.Name, m.Machine.Namespace, err)
	}
}
func MachineConfigFromProviderSpec(clusterClient machineclient.MachineClassesGetter, providerConfig machinev1.ProviderSpec) (*v1alpha1.AzureMachineProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var config v1alpha1.AzureMachineProviderSpec
	if providerConfig.Value != nil {
		klog.V(4).Info("Decoding ProviderConfig from Value")
		return unmarshalProviderSpec(providerConfig.Value)
	}
	if providerConfig.ValueFrom != nil && providerConfig.ValueFrom.MachineClass != nil {
		ref := providerConfig.ValueFrom.MachineClass
		klog.V(4).Info("Decoding ProviderConfig from MachineClass")
		klog.V(6).Infof("ref: %v", ref)
		if ref.Provider != "" && ref.Provider != "azure" {
			return nil, errors.Errorf("Unsupported provider: %q", ref.Provider)
		}
		if len(ref.Namespace) > 0 && len(ref.Name) > 0 {
			klog.V(4).Infof("Getting MachineClass: %s/%s", ref.Namespace, ref.Name)
			mc, err := clusterClient.MachineClasses(ref.Namespace).Get(ref.Name, metav1.GetOptions{})
			klog.V(6).Infof("Retrieved MachineClass: %+v", mc)
			if err != nil {
				return nil, err
			}
			providerConfig.Value = &mc.ProviderSpec
			return unmarshalProviderSpec(&mc.ProviderSpec)
		}
	}
	return &config, nil
}
func unmarshalProviderSpec(spec *runtime.RawExtension) (*v1alpha1.AzureMachineProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var config v1alpha1.AzureMachineProviderSpec
	if spec != nil {
		if err := yaml.Unmarshal(spec.Raw, &config); err != nil {
			return nil, err
		}
	}
	klog.V(6).Infof("Found ProviderSpec: %+v", config)
	return &config, nil
}
func updateScope(coreClient controllerclient.Client, credentialsSecret *apicorev1.SecretReference, scope *Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if credentialsSecret == nil {
		return errors.New("provided empty credentials secret")
	}
	secretType := types.NamespacedName{Namespace: credentialsSecret.Namespace, Name: credentialsSecret.Name}
	var secret apicorev1.Secret
	if err := coreClient.Get(context.Background(), secretType, &secret); err != nil {
		return err
	}
	subscriptionID, ok := secret.Data[AzureCredsSubscriptionIDKey]
	if !ok {
		return errors.Errorf("Azure subscription id %v did not contain key %v", secretType.String(), AzureCredsSubscriptionIDKey)
	}
	clientID, ok := secret.Data[AzureCredsClientIDKey]
	if !ok {
		return errors.Errorf("Azure client id %v did not contain key %v", secretType.String(), AzureCredsClientIDKey)
	}
	clientSecret, ok := secret.Data[AzureCredsClientSecretKey]
	if !ok {
		return errors.Errorf("Azure client secret %v did not contain key %v", secretType.String(), AzureCredsClientSecretKey)
	}
	tenantID, ok := secret.Data[AzureCredsTenantIDKey]
	if !ok {
		return errors.Errorf("Azure tenant id %v did not contain key %v", secretType.String(), AzureCredsTenantIDKey)
	}
	resourceGroup, ok := secret.Data[AzureCredsResourceGroupKey]
	if !ok {
		return errors.Errorf("Azure resource group %v did not contain key %v", secretType.String(), AzureCredsResourceGroupKey)
	}
	region, ok := secret.Data[AzureCredsRegionKey]
	if !ok {
		return errors.Errorf("Azure region %v did not contain key %v", secretType.String(), AzureCredsRegionKey)
	}
	clusterName, ok := secret.Data[AzureResourcePrefix]
	if !ok {
		return errors.Errorf("Azure resource prefix %v did not contain key %v", secretType.String(), AzureResourcePrefix)
	}
	env, err := azure.EnvironmentFromName("AzurePublicCloud")
	if err != nil {
		return err
	}
	oauthConfig, err := adal.NewOAuthConfig(env.ActiveDirectoryEndpoint, string(tenantID))
	if err != nil {
		return err
	}
	token, err := adal.NewServicePrincipalToken(*oauthConfig, string(clientID), string(clientSecret), env.ResourceManagerEndpoint)
	if err != nil {
		return err
	}
	authorizer, err := autorest.NewBearerAuthorizer(token), nil
	if err != nil {
		return errors.Errorf("failed to create azure session: %v", err)
	}
	scope.Cluster.ObjectMeta.Name = string(clusterName)
	scope.Authorizer = authorizer
	scope.SubscriptionID = string(subscriptionID)
	scope.ClusterConfig.ResourceGroup = string(resourceGroup)
	scope.ClusterConfig.Location = string(region)
	return nil
}
