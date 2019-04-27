package certificates

import (
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

type Service struct{ scope *actuators.Scope }

func NewService(scope *actuators.Scope) *Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Service{scope: scope}
}
