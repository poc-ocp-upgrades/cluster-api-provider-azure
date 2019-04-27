package actuators

import (
	"context"
	"fmt"
	"os"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	client "github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/typed/cluster/v1alpha1"
	"github.com/pkg/errors"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

type ScopeParams struct {
	AzureClients
	Cluster	*clusterv1.Cluster
	Client	client.ClusterV1alpha1Interface
}

func NewScope(params ScopeParams) (*Scope, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if params.Cluster == nil {
		return &Scope{AzureClients: params.AzureClients, Cluster: &clusterv1.Cluster{}, ClusterClient: nil, ClusterConfig: &v1alpha1.AzureClusterProviderSpec{}, ClusterStatus: &v1alpha1.AzureClusterProviderStatus{}, Context: context.Background()}, nil
	}
	clusterConfig, err := v1alpha1.ClusterConfigFromProviderSpec(params.Cluster.Spec.ProviderSpec)
	if err != nil {
		return nil, errors.Errorf("failed to load cluster provider config: %v", err)
	}
	clusterStatus, err := v1alpha1.ClusterStatusFromProviderStatus(params.Cluster.Status.ProviderStatus)
	if err != nil {
		return nil, errors.Errorf("failed to load cluster provider status: %v", err)
	}
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return nil, errors.Errorf("failed to create azure session: %v", err)
	}
	params.AzureClients.Authorizer = authorizer
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subscriptionID == "" {
		return nil, fmt.Errorf("error creating azure services. Environment variable AZURE_SUBSCRIPTION_ID is not set")
	}
	params.AzureClients.SubscriptionID = subscriptionID
	var clusterClient client.ClusterInterface
	if params.Client != nil {
		clusterClient = params.Client.Clusters(params.Cluster.Namespace)
	}
	return &Scope{AzureClients: params.AzureClients, Cluster: params.Cluster, ClusterClient: clusterClient, ClusterConfig: clusterConfig, ClusterStatus: clusterStatus, Context: context.Background()}, nil
}

type Scope struct {
	AzureClients
	Cluster		*clusterv1.Cluster
	ClusterClient	client.ClusterInterface
	ClusterConfig	*v1alpha1.AzureClusterProviderSpec
	ClusterStatus	*v1alpha1.AzureClusterProviderStatus
	Context		context.Context
}

func (s *Scope) Network() *v1alpha1.Network {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &s.ClusterStatus.Network
}
func (s *Scope) Vnet() *v1alpha1.VnetSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &s.ClusterConfig.NetworkSpec.Vnet
}
func (s *Scope) Subnets() v1alpha1.Subnets {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.ClusterConfig.NetworkSpec.Subnets
}
func (s *Scope) SecurityGroups() map[v1alpha1.SecurityGroupRole]*v1alpha1.SecurityGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.ClusterStatus.Network.SecurityGroups
}
func (s *Scope) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Cluster.Name
}
func (s *Scope) Namespace() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Cluster.Namespace
}
func (s *Scope) Location() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.ClusterConfig.Location
}
func (s *Scope) storeClusterConfig(cluster *clusterv1.Cluster) (*clusterv1.Cluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ext, err := v1alpha1.EncodeClusterSpec(s.ClusterConfig)
	if err != nil {
		return nil, err
	}
	cluster.Spec.ProviderSpec.Value = ext
	return s.ClusterClient.Update(cluster)
}
func (s *Scope) storeClusterStatus(cluster *clusterv1.Cluster) (*clusterv1.Cluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ext, err := v1alpha1.EncodeClusterStatus(s.ClusterStatus)
	if err != nil {
		return nil, err
	}
	cluster.Status.ProviderStatus = ext
	return s.ClusterClient.UpdateStatus(cluster)
}
func (s *Scope) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.ClusterClient == nil {
		return
	}
	latestCluster, err := s.storeClusterConfig(s.Cluster)
	if err != nil {
		klog.Errorf("[scope] failed to store provider config for cluster %q in namespace %q: %v", s.Cluster.Name, s.Cluster.Namespace, err)
		return
	}
	_, err = s.storeClusterStatus(latestCluster)
	if err != nil {
		klog.Errorf("[scope] failed to store provider status for cluster %q in namespace %q: %v", s.Cluster.Name, s.Cluster.Namespace, err)
	}
}
