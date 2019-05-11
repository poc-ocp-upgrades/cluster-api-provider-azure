package cluster

import (
	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/certificates"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/groups"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/internalloadbalancers"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/publicips"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/publicloadbalancers"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/routetables"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/securitygroups"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/subnets"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/virtualnetworks"
)

type Reconciler struct {
	scope				*actuators.Scope
	certificatesSvc		azure.Service
	groupsSvc			azure.Service
	vnetSvc				azure.Service
	securityGroupSvc	azure.Service
	routeTableSvc		azure.Service
	subnetsSvc			azure.Service
	internalLBSvc		azure.Service
	publicIPSvc			azure.Service
	publicLBSvc			azure.Service
}

func NewReconciler(scope *actuators.Scope) *Reconciler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Reconciler{scope: scope, certificatesSvc: certificates.NewService(scope), groupsSvc: groups.NewService(scope), vnetSvc: virtualnetworks.NewService(scope), securityGroupSvc: securitygroups.NewService(scope), routeTableSvc: routetables.NewService(scope), subnetsSvc: subnets.NewService(scope), internalLBSvc: internalloadbalancers.NewService(scope), publicIPSvc: publicips.NewService(scope), publicLBSvc: publicloadbalancers.NewService(scope)}
}
func (s *Reconciler) Reconcile() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	actuators.CreateOrUpdateNetworkAPIServerIP(s.scope)
	if err := s.certificatesSvc.CreateOrUpdate(s.scope.Context, nil); err != nil {
		return errors.Wrapf(err, "failed to createorupdate certificates for cluster %s", s.scope.Cluster.Name)
	}
	if err := s.groupsSvc.CreateOrUpdate(s.scope.Context, nil); err != nil {
		return errors.Wrapf(err, "failed to createorupdate resource group for cluster %s", s.scope.Cluster.Name)
	}
	vnetSpec := &virtualnetworks.Spec{Name: azure.GenerateVnetName(s.scope.Cluster.Name), CIDR: azure.DefaultVnetCIDR}
	if err := s.vnetSvc.CreateOrUpdate(s.scope.Context, vnetSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate virtual network for cluster %s", s.scope.Cluster.Name)
	}
	sgSpec := &securitygroups.Spec{Name: azure.GenerateControlPlaneSecurityGroupName(s.scope.Cluster.Name), IsControlPlane: true}
	if err := s.securityGroupSvc.CreateOrUpdate(s.scope.Context, sgSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate control plane network security group for cluster %s", s.scope.Cluster.Name)
	}
	sgSpec = &securitygroups.Spec{Name: azure.GenerateNodeSecurityGroupName(s.scope.Cluster.Name), IsControlPlane: false}
	if err := s.securityGroupSvc.CreateOrUpdate(s.scope.Context, sgSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate node network security group for cluster %s", s.scope.Cluster.Name)
	}
	rtSpec := &routetables.Spec{Name: azure.GenerateNodeRouteTableName(s.scope.Cluster.Name)}
	if err := s.routeTableSvc.CreateOrUpdate(s.scope.Context, rtSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate node route table for cluster %s", s.scope.Cluster.Name)
	}
	subnetSpec := &subnets.Spec{Name: azure.GenerateControlPlaneSubnetName(s.scope.Cluster.Name), CIDR: azure.DefaultControlPlaneSubnetCIDR, VnetName: azure.GenerateVnetName(s.scope.Cluster.Name), SecurityGroupName: azure.GenerateControlPlaneSecurityGroupName(s.scope.Cluster.Name)}
	if err := s.subnetsSvc.CreateOrUpdate(s.scope.Context, subnetSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate control plane subnet for cluster %s", s.scope.Cluster.Name)
	}
	subnetSpec = &subnets.Spec{Name: azure.GenerateNodeSubnetName(s.scope.Cluster.Name), CIDR: azure.DefaultNodeSubnetCIDR, VnetName: azure.GenerateVnetName(s.scope.Cluster.Name), SecurityGroupName: azure.GenerateControlPlaneSecurityGroupName(s.scope.Cluster.Name), RouteTableName: azure.GenerateNodeRouteTableName(s.scope.Cluster.Name)}
	if err := s.subnetsSvc.CreateOrUpdate(s.scope.Context, subnetSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate node subnet for cluster %s", s.scope.Cluster.Name)
	}
	internalLBSpec := &internalloadbalancers.Spec{Name: azure.GenerateInternalLBName(s.scope.Cluster.Name), SubnetName: azure.GenerateControlPlaneSubnetName(s.scope.Cluster.Name), VnetName: azure.GenerateVnetName(s.scope.Cluster.Name), IPAddress: azure.DefaultInternalLBIPAddress}
	if err := s.internalLBSvc.CreateOrUpdate(s.scope.Context, internalLBSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate control plane internal load balancer for cluster %s", s.scope.Cluster.Name)
	}
	publicIPSpec := &publicips.Spec{Name: s.scope.Network().APIServerIP.Name}
	if err := s.publicIPSvc.CreateOrUpdate(s.scope.Context, publicIPSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate control plane public ip for cluster %s", s.scope.Cluster.Name)
	}
	publicLBSpec := &publicloadbalancers.Spec{Name: azure.GeneratePublicLBName(s.scope.Cluster.Name), PublicIPName: s.scope.Network().APIServerIP.Name}
	if err := s.publicLBSvc.CreateOrUpdate(s.scope.Context, publicLBSpec); err != nil {
		return errors.Wrapf(err, "failed to createorupdate control plane public load balancer for cluster %s", s.scope.Cluster.Name)
	}
	return nil
}
func (s *Reconciler) Delete() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.deleteLB(); err != nil {
		return errors.Wrap(err, "failed to delete load balancer")
	}
	if err := s.deleteSubnets(); err != nil {
		return errors.Wrap(err, "failed to delete subnets")
	}
	rtSpec := &routetables.Spec{Name: azure.GenerateNodeRouteTableName(s.scope.Cluster.Name)}
	if err := s.routeTableSvc.Delete(s.scope.Context, rtSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete route table %s for cluster %s", azure.GenerateNodeRouteTableName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	if err := s.deleteNSG(); err != nil {
		return errors.Wrap(err, "failed to delete network security group")
	}
	vnetSpec := &virtualnetworks.Spec{Name: azure.GenerateVnetName(s.scope.Cluster.Name)}
	if err := s.vnetSvc.Delete(s.scope.Context, vnetSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete virtual network %s for cluster %s", azure.GenerateVnetName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	if err := s.groupsSvc.Delete(s.scope.Context, nil); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete resource group for cluster %s", s.scope.Cluster.Name)
		}
	}
	return nil
}
func (s *Reconciler) deleteLB() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	publicLBSpec := &publicloadbalancers.Spec{Name: azure.GeneratePublicLBName(s.scope.Cluster.Name)}
	if err := s.publicLBSvc.Delete(s.scope.Context, publicLBSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete lb %s for cluster %s", azure.GeneratePublicLBName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	publicIPSpec := &publicips.Spec{Name: s.scope.Network().APIServerIP.Name}
	if err := s.publicIPSvc.Delete(s.scope.Context, publicIPSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete public ip %s for cluster %s", s.scope.Network().APIServerIP.Name, s.scope.Cluster.Name)
		}
	}
	internalLBSpec := &internalloadbalancers.Spec{Name: azure.GenerateNodeSubnetName(s.scope.Cluster.Name)}
	if err := s.internalLBSvc.Delete(s.scope.Context, internalLBSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to internal load balancer %s for cluster %s", azure.GenerateNodeSubnetName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	return nil
}
func (s *Reconciler) deleteSubnets() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	subnetSpec := &subnets.Spec{Name: azure.GenerateNodeSubnetName(s.scope.Cluster.Name), VnetName: azure.GenerateVnetName(s.scope.Cluster.Name)}
	if err := s.subnetsSvc.Delete(s.scope.Context, subnetSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete %s subnet for cluster %s", azure.GenerateNodeSubnetName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	subnetSpec = &subnets.Spec{Name: azure.GenerateControlPlaneSubnetName(s.scope.Cluster.Name), VnetName: azure.GenerateVnetName(s.scope.Cluster.Name)}
	if err := s.subnetsSvc.Delete(s.scope.Context, subnetSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete %s subnet for cluster %s", azure.GenerateControlPlaneSubnetName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	return nil
}
func (s *Reconciler) deleteNSG() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sgSpec := &securitygroups.Spec{Name: azure.GenerateNodeSecurityGroupName(s.scope.Cluster.Name)}
	if err := s.securityGroupSvc.Delete(s.scope.Context, sgSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete security group %s for cluster %s", azure.GenerateNodeSecurityGroupName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	sgSpec = &securitygroups.Spec{Name: azure.GenerateControlPlaneSecurityGroupName(s.scope.Cluster.Name)}
	if err := s.securityGroupSvc.Delete(s.scope.Context, sgSpec); err != nil {
		if !azure.ResourceNotFound(err) {
			return errors.Wrapf(err, "failed to delete security group %s for cluster %s", azure.GenerateControlPlaneSecurityGroupName(s.scope.Cluster.Name), s.scope.Cluster.Name)
		}
	}
	return nil
}
