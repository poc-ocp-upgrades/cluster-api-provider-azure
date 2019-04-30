package actuators

var (
	DefaultScopeGetter		ScopeGetter		= ScopeGetterFunc(NewScope)
	DefaultMachineScopeGetter	MachineScopeGetter	= MachineScopeGetterFunc(NewMachineScope)
)

type ScopeGetter interface {
	GetScope(params ScopeParams) (*Scope, error)
}
type ScopeGetterFunc func(params ScopeParams) (*Scope, error)

func (f ScopeGetterFunc) GetScope(params ScopeParams) (*Scope, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f(params)
}

type MachineScopeGetter interface {
	GetMachineScope(params MachineScopeParams) (*MachineScope, error)
}
type MachineScopeGetterFunc func(params MachineScopeParams) (*MachineScope, error)

func (f MachineScopeGetterFunc) GetMachineScope(params MachineScopeParams) (*MachineScope, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f(params)
}
