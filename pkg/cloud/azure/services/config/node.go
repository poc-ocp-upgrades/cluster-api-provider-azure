package config

import (
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
)

const (
	nodeBashScript = `{{.Header}}

mkdir -p /etc/kubernetes/pki

echo -n '{{.CloudProviderConfig}}' > /etc/kubernetes/azure.json

cat >/tmp/kubeadm-node.yaml <<EOF
---
apiVersion: kubeadm.k8s.io/v1beta1
kind: JoinConfiguration
discovery:
  bootstrapToken:
    token: "{{.BootstrapToken}}"
    apiServerEndpoint: "{{.InternalLBAddress}}:6443"
    caCertHashes:
      - "{{.CACertHash}}"
nodeRegistration:
  criSocket: /var/run/containerd/containerd.sock
  kubeletExtraArgs:
    cloud-provider: azure
    cloud-config: /etc/kubernetes/azure.json
EOF

# Configure containerd prerequisites
modprobe overlay
modprobe br_netfilter

# Setup required sysctl params, these persist across reboots.
cat > /etc/sysctl.d/99-kubernetes-cri.conf <<EOF
net.bridge.bridge-nf-call-iptables  = 1
net.ipv4.ip_forward                 = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF

sysctl --system

apt-get install -y libseccomp2

# Install containerd
# Export required environment variables.
export CONTAINERD_VERSION="1.2.4"
export CONTAINERD_SHA256="3391758c62d17a56807ddac98b05487d9e78e5beb614a0602caab747b0eda9e0"

# Download containerd tar.
wget https://storage.googleapis.com/cri-containerd-release/cri-containerd-${CONTAINERD_VERSION}.linux-amd64.tar.gz

# Check hash.
echo "${CONTAINERD_SHA256} cri-containerd-${CONTAINERD_VERSION}.linux-amd64.tar.gz" | sha256sum --check -

# Unpack.
tar --no-overwrite-dir -C / -xzf cri-containerd-${CONTAINERD_VERSION}.linux-amd64.tar.gz

# Start containerd.
systemctl start containerd

# Install kubeadm (https://kubernetes.io/docs/setup/independent/install-kubeadm/)
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet="{{.KubernetesVersion}}-00" kubeadm="{{.KubernetesVersion}}-00" kubectl="{{.KubernetesVersion}}-00"
apt-mark hold kubelet kubeadm kubectl

kubeadm join --config /tmp/kubeadm-node.yaml || true
`
)

type NodeInput struct {
	baseConfig
	CACertHash			string
	BootstrapToken		string
	InternalLBAddress	string
	KubernetesVersion	string
	CloudProviderConfig	string
}

func NewNode(input *NodeInput) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	input.Header = defaultHeader
	return generate(v1alpha1.Node, nodeBashScript, input)
}
