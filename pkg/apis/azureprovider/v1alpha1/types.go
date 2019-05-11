package v1alpha1

import (
	"time"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AzureResourceReference struct {
	ID *string `json:"id,omitempty"`
}
type AzureMachineProviderConditionType string

const (
	MachineCreated AzureMachineProviderConditionType = "MachineCreated"
)

type AzureMachineProviderCondition struct {
	Type				AzureMachineProviderConditionType	`json:"type"`
	Status				corev1.ConditionStatus				`json:"status"`
	LastProbeTime		metav1.Time							`json:"lastProbeTime"`
	LastTransitionTime	metav1.Time							`json:"lastTransitionTime"`
	Reason				string								`json:"reason"`
	Message				string								`json:"message"`
}

const (
	ControlPlane		string	= "master"
	Node				string	= "worker"
	MachineRoleLabel			= "machine.openshift.io/cluster-api-machine-role"
)

type Network struct {
	SecurityGroups	map[SecurityGroupRole]*SecurityGroup	`json:"securityGroups,omitempty"`
	APIServerLB		LoadBalancer							`json:"apiServerLb,omitempty"`
	APIServerIP		PublicIP								`json:"apiServerIp,omitempty"`
}
type Subnets []*SubnetSpec

func (s Subnets) ToMap() map[string]*SubnetSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := make(map[string]*SubnetSpec)
	for _, x := range s {
		res[x.ID] = x
	}
	return res
}

type SecurityGroupRole string

var (
	SecurityGroupBastion		= SecurityGroupRole("bastion")
	SecurityGroupNode			= SecurityGroupRole(Node)
	SecurityGroupControlPlane	= SecurityGroupRole(ControlPlane)
)

type SecurityGroup struct {
	ID				string			`json:"id"`
	Name			string			`json:"name"`
	IngressRules	IngressRules	`json:"ingressRule"`
}
type SecurityGroupProtocol string

var (
	SecurityGroupProtocolAll	= SecurityGroupProtocol("*")
	SecurityGroupProtocolTCP	= SecurityGroupProtocol("Tcp")
	SecurityGroupProtocolUDP	= SecurityGroupProtocol("Udp")
)

type IngressRule struct {
	Description			string					`json:"description"`
	Protocol			SecurityGroupProtocol	`json:"protocol"`
	SourcePorts			*string					`json:"sourcePorts,omitempty"`
	DestinationPorts	*string					`json:"destinationPorts,omitempty"`
	Source				*string					`json:"source,omitempty"`
	Destination			*string					`json:"destination,omitempty"`
}
type IngressRules []*IngressRule
type PublicIP struct {
	ID			string	`json:"id,omitempty"`
	Name		string	`json:"name,omitempty"`
	IPAddress	string	`json:"ipAddress,omitempty"`
	DNSName		string	`json:"dnsName,omitempty"`
}
type LoadBalancer struct {
	ID					string				`json:"id,omitempty"`
	Name				string				`json:"name,omitempty"`
	SKU					SKU					`json:"sku,omitempty"`
	FrontendIPConfig	FrontendIPConfig	`json:"frontendIpConfig,omitempty"`
	BackendPool			BackendPool			`json:"backendPool,omitempty"`
}
type SKU string

var (
	SKUBasic	= SKU("Basic")
	SKUStandard	= SKU("Standard")
)

type FrontendIPConfig struct{}
type BackendPool struct {
	Name	string	`json:"name,omitempty"`
	ID		string	`json:"id,omitempty"`
}
type LoadBalancerProtocol string

var (
	LoadBalancerProtocolTCP		= LoadBalancerProtocol("TCP")
	LoadBalancerProtocolSSL		= LoadBalancerProtocol("SSL")
	LoadBalancerProtocolHTTP	= LoadBalancerProtocol("HTTP")
	LoadBalancerProtocolHTTPS	= LoadBalancerProtocol("HTTPS")
)

type LoadBalancerListener struct {
	Protocol			LoadBalancerProtocol	`json:"protocol"`
	Port				int64					`json:"port"`
	InstanceProtocol	LoadBalancerProtocol	`json:"instanceProtocol"`
	InstancePort		int64					`json:"instancePort"`
}
type LoadBalancerHealthCheck struct {
	Target				string			`json:"target"`
	Interval			time.Duration	`json:"interval"`
	Timeout				time.Duration	`json:"timeout"`
	HealthyThreshold	int64			`json:"healthyThreshold"`
	UnhealthyThreshold	int64			`json:"unhealthyThreshold"`
}
type VMState string

var (
	VMStateCreating		= VMState("Creating")
	VMStateDeleting		= VMState("Deleting")
	VMStateFailed		= VMState("Failed")
	VMStateMigrating	= VMState("Migrating")
	VMStateSucceeded	= VMState("Succeeded")
	VMStateUpdating		= VMState("Updating")
)

type VM struct {
	ID				string		`json:"id,omitempty"`
	Name			string		`json:"name,omitempty"`
	VMSize			string		`json:"vmSize,omitempty"`
	Image			Image		`json:"image,omitempty"`
	OSDisk			OSDisk		`json:"osDisk,omitempty"`
	StartupScript	string		`json:"startupScript,omitempty"`
	State			VMState		`json:"vmState,omitempty"`
	Identity		VMIdentity	`json:"identity,omitempty"`
}
type Image struct {
	Publisher	string	`json:"publisher"`
	Offer		string	`json:"offer"`
	SKU			string	`json:"sku"`
	Version		string	`json:"version"`
	ResourceID	string	`json:"resourceID"`
}
type VMIdentity string
type OSDisk struct {
	OSType		string		`json:"osType"`
	ManagedDisk	ManagedDisk	`json:"managedDisk"`
	DiskSizeGB	int32		`json:"diskSizeGB"`
}
type ManagedDisk struct {
	StorageAccountType string `json:"storageAccountType"`
}
