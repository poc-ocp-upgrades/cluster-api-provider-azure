package machine

import (
	"encoding/json"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
)

func (a *Actuator) updateMachineAnnotationJSON(machine *clusterv1.Machine, annotation string, content map[string]interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(content)
	if err != nil {
		return err
	}
	a.updateMachineAnnotation(machine, annotation, string(b))
	return nil
}
func (a *Actuator) updateMachineAnnotation(machine *clusterv1.Machine, annotation string, content string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	annotations := machine.GetAnnotations()
	annotations[annotation] = content
	machine.SetAnnotations(annotations)
}
func (a *Actuator) machineAnnotationJSON(machine *clusterv1.Machine, annotation string) (map[string]interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := map[string]interface{}{}
	jsonAnnotation := a.machineAnnotation(machine, annotation)
	if len(jsonAnnotation) == 0 {
		return out, nil
	}
	err := json.Unmarshal([]byte(jsonAnnotation), &out)
	if err != nil {
		return out, err
	}
	return out, nil
}
func (a *Actuator) machineAnnotation(machine *clusterv1.Machine, annotation string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return machine.GetAnnotations()[annotation]
}
