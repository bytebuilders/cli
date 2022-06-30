/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package installer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"

	api "go.bytebuilders.dev/installer/apis/installer/v1alpha1"

	"github.com/pkg/errors"
	passgen "gomodules.xyz/password-generator"
	"gomodules.xyz/pointer"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"kubepack.dev/kubepack/apis/kubepack/v1alpha1"
	pkglib "kubepack.dev/kubepack/pkg/lib"
	"sigs.k8s.io/yaml"
)

// Steps to do manually
// buckets to create
// DNS record configure

func generateInstaller(in *api.AceOptionsSpec) (string, error) {
	chrtRef := v1alpha1.ChartRef{
		URL:  "https://charts.appscode.com/stable/",
		Name: "ace",
	}

	// chartDir := filepath.Join(homedir.HomeDir(), "go/src/go.bytebuilders.dev/installer")

	if err := os.RemoveAll(confDir()); err != nil {
		return "", errors.Wrapf(err, "failed to delete dir: %s", confDir())
	}
	if err := os.MkdirAll(confDir(), 0o755); err != nil {
		return "", errors.Wrapf(err, "failed to create dir: %s", confDir())
	}

	//{
	//	opts := NewOptions()
	//	if data, err := yaml.Marshal(*opts); err != nil {
	//		return err
	//	} else {
	//		_ = ioutil.WriteFile(filepath.Join(confDir(), "options-initial.yaml"), data, 0o644)
	//	}
	//}

	// in := NewSampleOptions()
	if data, err := yaml.Marshal(*in); err != nil {
		return "", err
	} else {
		_ = ioutil.WriteFile(filepath.Join(confDir(), "options.yaml"), data, 0o644)
	}

	chrt, err := pkglib.DefaultRegistry.GetChart(chrtRef.URL, chrtRef.Name, "")
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(chrt.Values)
	if err != nil {
		return "", err
	}
	outOrig := new(api.AceSpec)
	err = yaml.Unmarshal(data, outOrig)
	if err != nil {
		return "", err
	}
	err = InitComponentsOut(in, outOrig)
	if err != nil {
		return "", err
	}

	outMod := outOrig.DeepCopy()
	err = Convert(in, outOrig, outMod)
	if err != nil {
		return "", err
	}

	//{
	//	cmd, _, err := pkglib.PrintHelm3CommandFromStructValues(v1alpha1.InstallOptions{
	//		ChartRef: v1alpha1.ChartRef{
	//			URL:  "https://charts.appscode.com/stable/",
	//			Name: "ace",
	//		},
	//		Version:     "",
	//		ReleaseName: "ace",
	//		Namespace:   "ace",
	//	}, outOrig, outMod, false)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println(cmd)
	//}

	{
		cmd, vb, err := pkglib.PrintHelm3CommandFromStructValues(v1alpha1.InstallOptions{
			ChartRef: v1alpha1.ChartRef{
				URL:  "https://charts.appscode.com/stable/",
				Name: "ace",
			},
			Version:     "",
			ReleaseName: "ace",
			Namespace:   "ace",
		}, outOrig, outMod, true)
		if err != nil {
			return "", err
		}
		// fmt.Println(cmd)
		_ = ioutil.WriteFile(filepath.Join(confDir(), "values.yaml"), vb, 0o644)

		delim := "```"
		md := fmt.Sprintf(`# Install ACE

## Install Prerequisites

- Create New Google Cloud Project appscode-ace
- Add Billing account
- Gave eng@appscode.com owner access
- create buckets gs://ace-avatars, gs://ace-invoices
- TODO: create bucket for kubepack
- created new service account
- Add "Storage Object Creator" permission to the buckets
- Get token from Cloudflare for appscode.cloud Domain
- not using KMS

%s
helm upgrade -i kubedb appscode/kubedb \
  --version v2022.05.24 \
  --namespace kubedb --create-namespace \
  --set kubedb-provisioner.enabled=true \
  --set kubedb-ops-manager.enabled=false \
  --set kubedb-autoscaler.enabled=false \
  --set kubedb-dashboard.enabled=false \
  --set kubedb-schema-manager.enabled=false \
  --set-file global.license=/Users/tamal/Downloads/kubedb-enterprise-license-20aae10d-67db-4041-bcdf-fe46f58d9231.txt

helm upgrade -i stash appscode/stash \
  --version v2022.05.18 \
  --namespace stash --create-namespace \
  --set features.enterprise=true \
  --set-file global.license=/Users/tamal/Downloads/kubedb-enterprise-license-20aae10d-67db-4041-bcdf-fe46f58d9231.txt

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.8.0 \
  --set installCRDs=true

helm upgrade -i kube-prometheus-stack prometheus-community/kube-prometheus-stack \
  --namespace monitoring --create-namespace
%s

## Deply ACE


%s
%s
%s`, delim, delim, delim, cmd, delim)
		_ = ioutil.WriteFile(filepath.Join(confDir(), "README.md"), []byte(md), 0o644)

		return md, nil
	}

	//sh := shell.NewSession()
	//sh.SetDir(chartDir)
	//sh.ShowCMD = true
	//
	//if data, err := sh.Command("helm", "template", "charts/ace", "--values", filepath.Join(confDir(), "values.yaml")).Output(); err != nil {
	//	return err
	//} else {
	//	_ = ioutil.WriteFile(filepath.Join(confDir(), "ace.yaml"), data, 0o644)
	//}
}

func Convert(in *api.AceOptionsSpec, base, out *api.AceSpec) error {
	if err := InitComponents(in, out); err != nil {
		return err
	}
	if err := GeneratePlatformValues(in, base, out); err != nil {
		return err
	}
	if err := GenerateIngress(in, out); err != nil {
		return err
	}
	if err := GenerateNats(in, out); err != nil {
		return err
	}
	return nil
}

func NewOptions() *api.AceOptionsSpec {
	hosted := false
	return &api.AceOptionsSpec{
		Release: api.ObjectReference{
			Name:      "ace",
			Namespace: "ace",
		},
		License:          "",
		Registry:         "",
		RegistryFQDN:     "",
		ImagePullSecrets: nil,
		Monitoring:       api.GlobalMonitoring{},
		Infra: api.AceOptionsPlatformInfra{
			StorageClass: api.LocalObjectReference{
				Name: "standard",
			},
			// TLS:      api.AceOptionsInfraTLS{},
			// DNS:      api.InfraDns{},
			CloudServices: api.AceOptionsInfraCloudServices{
				// Provider: "",
				// Auth:     api.ObjstoreAuth{},
				Objstore: api.AceOptionsInfraObjstore{
					Bucket: "gs://appscode",
				},
				// Kms:      nil,
			},
		},
		Settings: api.AceOptionsSettings{
			DB: api.AceOptionsDBSettings{
				Persistence: api.PersistenceSpec{
					Size: resource.MustParse("20Gi"),
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			Cache: api.AceOptionsCacheSettings{
				Persistence: api.PersistenceSpec{
					Size: resource.MustParse("10Gi"),
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			SMTP: api.AceOptionsSMTPSettings{},
			Platform: api.AceOptionsPlatformSettings{
				Domain: "",
				Hosted: hosted,
			},
			//Security: api.SecuritySettings{
			//	Oauth2JWTSecret: "",
			//	CsrfSecretKey:   "",
			//},
		},
		Billing: api.AceOptionsComponentSpec{
			Enabled: hosted,
		},
		PlatformUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		AccountsUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		ClusterUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		DeployUi: api.AceOptionsComponentSpec{
			Enabled: hosted,
		},
		Grafana: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		KubedbUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		MarketplaceUi: api.AceOptionsComponentSpec{
			Enabled: hosted,
		},
		PlatformApi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		PromProxy: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		Ingress: api.AceOptionsIngressNginx{
			ExposeVia: api.ServiceTypeLoadBalancer,
			// Resources:    core.ResourceRequirements{},
			// NodeSelector: nil,
		},
		Nats: api.AceOptionsNatsSettings{
			ExposeVia: api.ServiceTypeLoadBalancer,
			Replics:   1,
			//Resources:    core.ResourceRequirements{
			//	Limits:   nil,
			//	Requests: nil,
			//},
			//NodeSelector: nil,
		},
	}
}

func SampleResource() core.ResourceRequirements {
	return core.ResourceRequirements{
		Limits: core.ResourceList{
			core.ResourceMemory: resource.MustParse("128Mi"),
		},
		Requests: core.ResourceList{
			core.ResourceMemory: resource.MustParse("128Mi"),
		},
	}
}

func InitComponentsOut(in *api.AceOptionsSpec, out *api.AceSpec) error {
	out.Reloader = api.AceReloader{
		Enabled: true,
	}

	if in.Billing.Enabled {
		out.Billing = api.AceBilling{
			Enabled:     true,
			BillingSpec: &api.BillingSpec{},
		}
	}
	if in.PlatformUi.Enabled {
		out.PlatformUi = api.AcePlatformUi{
			Enabled:        true,
			PlatformUiSpec: &api.PlatformUiSpec{},
		}
	}
	if in.AccountsUi.Enabled {
		out.AccountsUi = api.AceAccountsUi{
			Enabled:        true,
			AccountsUiSpec: &api.AccountsUiSpec{},
		}
	}
	if in.ClusterUi.Enabled {
		out.ClusterUi = api.AceClusterUi{
			Enabled:       true,
			ClusterUiSpec: &api.ClusterUiSpec{},
		}
	}
	if in.DeployUi.Enabled {
		out.DeployUi = api.AceDeployUi{
			Enabled:      true,
			DeployUiSpec: &api.DeployUiSpec{},
		}
	}
	if in.Grafana.Enabled {
		out.Grafana = api.AceGrafana{
			Enabled:     true,
			GrafanaSpec: &api.GrafanaSpec{},
		}
	}
	if in.KubedbUi.Enabled {
		out.KubedbUi = api.AceKubedbUi{
			Enabled:      true,
			KubedbUiSpec: &api.KubedbUiSpec{},
		}
	}
	if in.MarketplaceUi.Enabled {
		out.MarketplaceUi = api.AceMarketplaceUi{
			Enabled:           true,
			MarketplaceUiSpec: &api.MarketplaceUiSpec{},
		}
	}
	if in.PlatformApi.Enabled {
		out.PlatformApi = api.AcePlatformApi{
			Enabled:         true,
			PlatformApiSpec: &api.PlatformApiSpec{},
		}
	}
	if in.PromProxy.Enabled {
		out.PromProxy = api.AcePromProxy{
			Enabled:       true,
			PromProxySpec: &api.PromProxySpec{},
		}
	}

	out.IngressNginx = api.AceIngressNginx{
		Enabled:          true,
		IngressNginxSpec: &api.IngressNginxSpec{},
	}
	out.IngressDns = api.AceIngressDns{
		Enabled:         false,
		ExternalDnsSpec: &api.ExternalDnsSpec{},
	}

	out.Nats = api.AceNats{
		Enabled:  true,
		NatsSpec: &api.NatsSpec{},
	}
	if in.Nats.ExposeVia != api.ServiceTypeLoadBalancer {
		out.NatsDns = api.AceNatsDns{
			Enabled:         false,
			ExternalDnsSpec: &api.ExternalDnsSpec{},
		}
	}

	return nil
}

func NewSampleOptions() *api.AceOptionsSpec {
	hosted := false
	return &api.AceOptionsSpec{
		Release: api.ObjectReference{
			Name:      "ace",
			Namespace: "ace",
		},
		License:          "",
		Registry:         "",
		RegistryFQDN:     "",
		ImagePullSecrets: nil,
		Monitoring: api.GlobalMonitoring{
			Agent: "prometheus.io/operator",
			ServiceMonitor: api.GlobalServiceMonitor{
				Labels: map[string]string{
					"release": "kube-prometheus-stack",
				},
			},
			Exporter: api.GlobalPrometheusExporter{
				Resources: SampleResource(),
			},
		},
		Infra: api.AceOptionsPlatformInfra{
			StorageClass: api.LocalObjectReference{
				Name: "standard",
			},
			TLS: api.AceOptionsInfraTLS{
				Email: "ops@appscode.cloud",
			},
			DNS: api.InfraDns{
				Provider: "cloudflare",
				Auth: api.DNSProviderAuth{
					Email: "---",
					Token: "XYZ",
				},
			},
			CloudServices: api.AceOptionsInfraCloudServices{
				Provider: "Google",
				Auth: api.ObjstoreAuth{
					ServiceAccountJson: `{"secret": "json"}`,
				},
				Objstore: api.AceOptionsInfraObjstore{
					Bucket: "gs://ace",
				},
				//Kms: &api.AceOptionsInfraKms{
				//	MasterKeyURL: "",
				//},
			},
		},
		Settings: api.AceOptionsSettings{
			DB: api.AceOptionsDBSettings{
				Persistence: api.PersistenceSpec{
					Size: resource.MustParse("20Gi"),
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			Cache: api.AceOptionsCacheSettings{
				Persistence: api.PersistenceSpec{
					Size: resource.MustParse("10Gi"),
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("512Mi"),
					},
				},
			},
			SMTP: api.AceOptionsSMTPSettings{},
			Platform: api.AceOptionsPlatformSettings{
				Domain: "appscode.cloud",
				Hosted: hosted,
			},
			//Security: api.SecuritySettings{
			//	Oauth2JWTSecret: "",
			//	CsrfSecretKey:   "",
			//},
		},
		Billing: api.AceOptionsComponentSpec{
			Enabled: hosted,
		},
		PlatformUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		AccountsUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		ClusterUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		DeployUi: api.AceOptionsComponentSpec{
			Enabled: hosted,
		},
		Grafana: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		KubedbUi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		MarketplaceUi: api.AceOptionsComponentSpec{
			Enabled: hosted,
		},
		PlatformApi: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		PromProxy: api.AceOptionsComponentSpec{
			Enabled: true,
		},
		Ingress: api.AceOptionsIngressNginx{
			ExposeVia: api.ServiceTypeHostPort,
			// Resources:    core.ResourceRequirements{},
			// NodeSelector: nil,
		},
		Nats: api.AceOptionsNatsSettings{
			ExposeVia: api.ServiceTypeHostPort,
			Replics:   1,
			//Resources:    core.ResourceRequirements{
			//	Limits:   nil,
			//	Requests: nil,
			//},
			//NodeSelector: nil,
		},
	}
}

func InitComponents(in *api.AceOptionsSpec, out *api.AceSpec) error {
	out.Reloader = api.AceReloader{
		Enabled: true,
	}

	if in.Billing.Enabled {
		out.Billing = api.AceBilling{
			Enabled: true,
			BillingSpec: &api.BillingSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.Billing.Resources,
				NodeSelector: in.Billing.NodeSelector,
			},
		}
	}
	if in.PlatformUi.Enabled {
		out.PlatformUi = api.AcePlatformUi{
			Enabled: true,
			PlatformUiSpec: &api.PlatformUiSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.PlatformUi.Resources,
				NodeSelector: in.PlatformUi.NodeSelector,
			},
		}
	}
	if in.AccountsUi.Enabled {
		out.AccountsUi = api.AceAccountsUi{
			Enabled: true,
			AccountsUiSpec: &api.AccountsUiSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.AccountsUi.Resources,
				NodeSelector: in.AccountsUi.NodeSelector,
			},
		}
	}
	if in.ClusterUi.Enabled {
		out.ClusterUi = api.AceClusterUi{
			Enabled: true,
			ClusterUiSpec: &api.ClusterUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.ClusterUi.Resources,
				NodeSelector: in.ClusterUi.NodeSelector,
			},
		}
	}
	if in.DeployUi.Enabled {
		out.DeployUi = api.AceDeployUi{
			Enabled: true,
			DeployUiSpec: &api.DeployUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.DeployUi.Resources,
				NodeSelector: in.DeployUi.NodeSelector,
			},
		}
	}
	if in.Grafana.Enabled {
		out.Grafana = api.AceGrafana{
			Enabled: true,
			GrafanaSpec: &api.GrafanaSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.Grafana.Resources,
				NodeSelector: in.Grafana.NodeSelector,
			},
		}
	}
	if in.KubedbUi.Enabled {
		out.KubedbUi = api.AceKubedbUi{
			Enabled: true,
			KubedbUiSpec: &api.KubedbUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.KubedbUi.Resources,
				NodeSelector: in.KubedbUi.NodeSelector,
			},
		}
	}
	if in.MarketplaceUi.Enabled {
		out.MarketplaceUi = api.AceMarketplaceUi{
			Enabled: true,
			MarketplaceUiSpec: &api.MarketplaceUiSpec{
				//PodAnnotations: map[string]string{
				//	"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				//},
				Resources:    in.MarketplaceUi.Resources,
				NodeSelector: in.MarketplaceUi.NodeSelector,
			},
		}
	}
	if in.PlatformApi.Enabled {
		out.PlatformApi = api.AcePlatformApi{
			Enabled: true,
			PlatformApiSpec: &api.PlatformApiSpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.PlatformApi.Resources,
				NodeSelector: in.PlatformApi.NodeSelector,
			},
		}
	}
	if in.PromProxy.Enabled {
		out.PromProxy = api.AcePromProxy{
			Enabled: true,
			PromProxySpec: &api.PromProxySpec{
				PodAnnotations: map[string]string{
					"secret.reloader.stakater.com/reload": tplPlatformConfig(in),
				},
				Resources:    in.PromProxy.Resources,
				NodeSelector: in.PromProxy.NodeSelector,
			},
		}
	}
	return nil
}

func GenerateIngress(in *api.AceOptionsSpec, out *api.AceSpec) error {
	if in.Ingress.ExposeVia == api.ServiceTypeLoadBalancer {
		out.IngressNginx = api.AceIngressNginx{
			Enabled: true,
			IngressNginxSpec: &api.IngressNginxSpec{
				Controller: api.IngressNginxController{
					IngressClassByName: true,
					IngressClassResource: api.IngressNginxControllerIngressClassResource{
						Enabled:         true,
						ControllerValue: fmt.Sprintf("k8s.io/ingress-nginx-%s", in.Release.Name),
						Name:            fmt.Sprintf("nginx-%s", in.Release.Name),
					},
					//HostPort: &api.IngressNginxControllerHostPort{
					//	Enabled: false,
					//},
					//Kind:         "",
					NodeSelector: in.Ingress.NodeSelector,
					Service: &api.IngressNginxControllerService{
						External: api.IngressNginxControllerServiceExternal{
							Enabled: true,
						},
					},
				},
			},
		}
	} else {
		out.IngressNginx = api.AceIngressNginx{
			Enabled: true,
			IngressNginxSpec: &api.IngressNginxSpec{
				Controller: api.IngressNginxController{
					IngressClassByName: true,
					IngressClassResource: api.IngressNginxControllerIngressClassResource{
						Enabled:         true,
						ControllerValue: fmt.Sprintf("k8s.io/ingress-nginx-%s", in.Release.Name),
						Name:            fmt.Sprintf("nginx-%s", in.Release.Name),
					},
					HostPort: &api.IngressNginxControllerHostPort{
						Enabled: true,
					},
					Kind:         "DaemonSet",
					NodeSelector: in.Ingress.NodeSelector,
					Service: &api.IngressNginxControllerService{
						External: api.IngressNginxControllerServiceExternal{
							Enabled: false,
						},
					},
				},
			},
		}
	}

	out.IngressDns = api.AceIngressDns{
		Enabled: true,
		ExternalDnsSpec: &api.ExternalDnsSpec{
			DomainFilters: []string{
				in.Settings.Platform.Domain,
			},
			// ref: https://github.com/kubernetes-sigs/external-dns/pull/2718
			Image: api.ExternalDnsImageReference{
				Repository: "appscode/external-dns",
				Tag:        "external-dns-helm-chart-1.9.0-1-gbd1bb40c",
				PullPolicy: "IfNotPresent",
			},
			LogLevel:   "debug",
			Sources:    []string{"ingress"},
			ExtraArgs:  []string{"--ignore-ingress-tls-spec"},
			Policy:     "sync",
			Registry:   "txt",
			TxtOwnerID: "ingress-dns",
		},
	}

	// TODO: Add additional DNS providers
	if in.Infra.DNS.Provider == "cloudflare" {
		out.IngressDns.Provider = "cloudflare"
		out.IngressDns.Env = []core.EnvVar{
			{
				Name:  "CF_API_TOKEN",
				Value: in.Infra.DNS.Auth.Token,
			},
		}
	}

	return nil
}

func GenerateNats(in *api.AceOptionsSpec, out *api.AceSpec) error {
	if in.Nats.Replics != 1 && in.Nats.Replics != 3 {
		return errors.Errorf("nats replicas can be 1 or 3, found %d", in.Nats.Replics)
	}

	nc, err := genNatsCredentials()
	if err != nil {
		return err
	}

	out.Settings.Nats = api.NatsSettings{
		ShardCount:      128, // reduce to 32
		Replics:         in.Nats.Replics,
		MountPath:       "/nats/creds",
		OperatorCreds:   nc["Operator.creds"],
		OperatorJwt:     nc["Operator.jwt"],
		SystemCreds:     nc["SYS.creds"],
		SystemJwt:       nc["SYS.jwt"],
		SystemPubKey:    nc["SYS.pub"],
		SystemUserCreds: nc["sys.creds"],
		AdminCreds:      nc["Admin.creds"],
		AdminUserCreds:  nc["admin.creds"],
	}

	out.Nats = api.AceNats{
		Enabled: true,
		NatsSpec: &api.NatsSpec{
			NodeSelector: in.Nats.NodeSelector,
			StatefulSetPodLabels: map[string]string{
				"secret.reloader.stakater.com/reload": tplNATSTLSSecret(in),
			},
			Nats: api.NatsServerSpec{
				Advertise: false,
				// ExternalAccess: true, // true means HostPost
				Limits: api.NatsServerLimitsSpec{
					MaxPayload: pointer.StringP("4Mb"),
				},
				Logging: api.NatsLoggingSpec{
					Debug: pointer.FalseP(),
					Trace: pointer.FalseP(),
				},
				Jetstream: api.JetstreamSpec{
					Enabled: true,
					FileStorage: api.JetstreamFileStorage{
						Enabled:          true,
						StorageDirectory: "/nats/jetstream",
						Size:             resource.MustParse("10Gi"), // TODO: high?
						StorageClassName: in.Infra.StorageClass.Name,
					},
				},
				Resources: core.ResourceRequirements{
					Limits: core.ResourceList{
						core.ResourceMemory: resource.MustParse("2Gi"),
					},
					Requests: core.ResourceList{
						core.ResourceMemory: resource.MustParse("2Gi"),
					},
				},
				TLS: &api.NatsServerTLSSpec{
					AllowNonTLS: false,
					Secret: api.LocalObjectReference{
						Name: tplNATSTLSSecret(in),
					},
					// Ca:          "",
					Cert: core.TLSCertKey,
					Key:  core.TLSPrivateKeyKey,
				},
			},
			Natsbox: api.NatsboxSpec{
				Enabled: false,
			},
			Exporter: api.NatsExporterSpec{
				Enabled: true,
				ServiceMonitor: api.NatsExporterServiceMonitorSpec{
					Enabled:   true,
					Namespace: "", // use nats namespace
					Labels:    in.Monitoring.ServiceMonitor.Labels,
					Path:      "/metrics",
				},
			},
			// Affinity:  nil,
			// Cluster:   api.NatsClusterSpec{},
			Auth: api.NatsAuthSpec{
				Enabled: true,
				Operatorjwt: &api.NatsOperatorJWTSpec{
					ConfigMap: api.ConfigMapKeySelector{
						Name: tplNATSCredSecret(in),
						Key:  "Operator.jwt",
					},
				},
				SystemAccount: pointer.StringP(nc["SYS.pub"]), // account or user?
				Resolver: api.NatsResolverSpec{
					Type:          "full",
					Operator:      pointer.StringP(nc["Operator.jwt"]),
					SystemAccount: pointer.StringP(nc["SYS.pub"]), // account or user
					Store: api.NatsResolverStoreSpec{
						Dir:              "/etc/nats-config/accounts/jwt",
						Size:             resource.MustParse("10Gi"),
						StorageClassName: in.Infra.StorageClass.Name,
					},
					ResolverPreload: map[string]string{
						nc["SYS.pub"]:   nc["SYS.jwt"],
						nc["Admin.pub"]: nc["Admin.jwt"], // TODO: skip?
					},
				},
			},
			Websocket: api.NatsWebsocketSpec{
				Enabled: true,
				Port:    443,
				AllowedOrigins: []string{
					fmt.Sprintf("https://%s", in.Settings.Platform.Domain),
					fmt.Sprintf("https://console.%s", in.Settings.Platform.Domain),
					fmt.Sprintf("https://kubedb.%s", in.Settings.Platform.Domain),
					fmt.Sprintf("https://grafana.%s", in.Settings.Platform.Domain),
				},
				TLS: &api.TLSSpec{
					Secret: api.LocalObjectReference{
						Name: tplNATSTLSSecret(in),
					},
					// Ca:          "",
					Cert: core.TLSCertKey,
					Key:  core.TLSPrivateKeyKey,
				},
			},
			UseFQDN: false,
		},
	}
	if in.Nats.Replics > 1 {
		natsPodSelector := &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app.kubernetes.io/instance": "nats-server",
				"app.kubernetes.io/name":     "nats",
			},
		}
		out.Nats.Affinity = &core.Affinity{
			PodAntiAffinity: &core.PodAntiAffinity{
				PreferredDuringSchedulingIgnoredDuringExecution: []core.WeightedPodAffinityTerm{
					// Prefer to not schedule multiple pods on the same node
					{
						Weight: 100,
						PodAffinityTerm: core.PodAffinityTerm{
							Namespaces:    []string{in.Release.Namespace},
							LabelSelector: natsPodSelector,
							TopologyKey:   core.LabelHostname,
						},
					},
					// Prefer to not schedule multiple pods on the node with same zone
					{
						Weight: 50,
						PodAffinityTerm: core.PodAffinityTerm{
							Namespaces:    []string{in.Release.Namespace},
							LabelSelector: natsPodSelector,
							TopologyKey:   core.LabelTopologyZone,
						},
					},
				},
			},
		}

		out.Nats.Cluster = api.NatsClusterSpec{
			Enabled:  true,
			Replicas: in.Nats.Replics,
			TLS: &api.TLSSpec{
				Secret: api.LocalObjectReference{
					Name: tplNATSTLSSecret(in),
				},
				// Ca:          "",
				Cert: core.TLSCertKey,
				Key:  core.TLSPrivateKeyKey,
			},
		}
	}

	if in.Nats.ExposeVia == api.ServiceTypeLoadBalancer {
		out.Nats.Nats.ExternalAccess = false
		// out.Nats.Websocket.Port = 9222

		// ingress TCP
		// expose NATS client port via TCP
		out.IngressNginx.IngressNginxSpec.TCP = map[string]string{
			"4222": fmt.Sprintf("%s/nats-server:4222", in.Release.Namespace),
		}
	} else {
		// out.Nats.Websocket.Port = 443
		out.Nats.Nats.ExternalAccess = true

		out.NatsDns = api.AceNatsDns{
			Enabled: true,
			ExternalDnsSpec: &api.ExternalDnsSpec{
				Sources: []string{"node"},
				Image: api.ExternalDnsImageReference{
					Repository: "appscode/external-dns",
					Tag:        "external-dns-helm-chart-1.9.0-1-gbd1bb40c",
					PullPolicy: "IfNotPresent",
				},
				DomainFilters: []string{in.Settings.Platform.Domain},
				LogLevel:      "debug",
				ExtraArgs: []string{
					fmt.Sprintf("--fqdn-template=nats.%s", in.Settings.Platform.Domain),
					fmt.Sprintf("--label-filter=%s", labels.Set(in.Nats.NodeSelector).String()),
				},
				Policy:     "sync",
				Registry:   "txt",
				TxtOwnerID: "nats-dns",
			},
		}

		// TODO: Add additional DNS providers
		if in.Infra.DNS.Provider == "cloudflare" {
			out.NatsDns.Provider = "cloudflare"
			out.NatsDns.Env = []core.EnvVar{
				{
					Name:  "CF_API_TOKEN",
					Value: in.Infra.DNS.Auth.Token,
				},
			}
		}
	}

	return nil
}

func GeneratePlatformValues(in *api.AceOptionsSpec, base, out *api.AceSpec) error {
	out.Global = api.AceGlobalValues{
		NameOverride:     in.Release.Name,
		FullnameOverride: base.Global.FullnameOverride,
		Platform: api.AceOptionsPlatformSettings{
			Domain: in.Settings.Platform.Domain,
			Hosted: in.Settings.Platform.Hosted,
		},
		License:          base.Global.License,
		Registry:         base.Global.Registry,
		RegistryFQDN:     base.Global.RegistryFQDN,
		ImagePullSecrets: base.Global.ImagePullSecrets,
		ServiceAccount:   base.Global.ServiceAccount,
		Monitoring:       in.Monitoring,
		Infra: api.PlatformInfra{
			StorageClass: in.Infra.StorageClass,
			TLS: api.InfraTLS{
				// TODO: prod URL: https://acme-v02.api.letsencrypt.org/directory
				AcmeServer: "https://acme-staging-v02.api.letsencrypt.org/directory",
				Email:      in.Infra.TLS.Email,
			},
			DNS: in.Infra.DNS,
			Objstore: api.InfraObjstore{
				Provider:  in.Infra.CloudServices.Provider,
				MountPath: "/data/credentials",
				Auth:      in.Infra.CloudServices.Auth,
			},
			//Kms: api.InfraKms{
			//	Provider:     in.Infra.Objstore.Provider,
			//	MasterKeyURL: fmt.Sprintf("base64key://%s", passgen.GenerateForCharset(64, passgen.AlphaNum)),
			//},
			Avatars: api.InfraAvatars{
				Bucket: mustBucketName(in.Infra.CloudServices.Objstore.Bucket, "avatars"),
			},
			// TODO: bucket proxy
			//Kubepack: api.InfraKubepack{
			//	Host:   "",
			//	Bucket: "",
			//},
			// TODO: skip Customer install vs appscode install
			//Badger: api.InfraBadger{
			//	MountPath: "/badger",
			//	Levels:    7,
			//},
			//Invoice: api.InfraInvoice{
			//	MountPath:    "/billing",
			//	Bucket:       mustBucketName(in.Infra.Objstore.Bucket, "invoices"),
			//	TrackerEmail: "",
			//},
		},
	}
	if in.Infra.CloudServices.Kms == nil || in.Infra.CloudServices.Kms.MasterKeyURL == "" {
		out.Global.Infra.Kms = api.InfraKms{
			MasterKeyURL: fmt.Sprintf("base64key://%s", passgen.GenerateForCharset(64, passgen.AlphaNum)),
		}
	} else {
		out.Global.Infra.Kms = api.InfraKms{
			MasterKeyURL: in.Infra.CloudServices.Kms.MasterKeyURL,
		}
	}

	if in.Settings.Platform.Hosted {
		// TODO: bucket proxy
		//out.Global.Infra.Kubepack = api.InfraKubepack{
		//	Host:   "",
		//	Bucket: "gs://",
		//}
		out.Global.Infra.Badger = api.InfraBadger{
			MountPath: "/badger",
			Levels:    7,
		}
		out.Global.Infra.Invoice = api.InfraInvoice{
			MountPath:    "/billing",
			Bucket:       mustBucketName(in.Infra.CloudServices.Objstore.Bucket, "invoices"),
			TrackerEmail: "",
		}
	}

	out.Settings = api.Settings{
		DB: api.DBSettings{
			Version:           "13.2",
			DatabaseName:      in.Release.Name,
			TerminationPolicy: "Delete", // TODO: change for prod mode
			Persistence:       in.Settings.DB.Persistence,
			Resources:         in.Settings.DB.Resources,
			Auth: api.BasicAuth{
				Username: "postgres",
				Password: passgen.Generate(api.DefaultPasswordLength),
			},
		},
		Cache: api.CacheSettings{
			Version:           "6.0.6",
			TerminationPolicy: "Delete",
			Persistence:       in.Settings.Cache.Persistence,
			Resources:         in.Settings.Cache.Resources,
			Auth: api.BasicAuth{
				Username: "root",
				Password: passgen.Generate(api.DefaultPasswordLength),
			},
			CacheInterval: 60,
		},
		Smtp: api.SmtpSettings{
			Host:       in.Settings.SMTP.Host,
			TlsEnabled: in.Settings.SMTP.TlsEnabled,
			From:       in.Settings.SMTP.From, // fmt.Sprintf("no-reply@%s", in.AceOptionsSettings.Platform.Domain), // TODO: configure?
			Username:   in.Settings.SMTP.Username,
			Password:   in.Settings.SMTP.Password,
			SubjectPrefix: func() string {
				if in.Settings.Platform.Hosted {
					return "ByteBuilders |"
				}
				return "ACE |"
			}(),
			SendAsPlainText: in.Settings.SMTP.SendAsPlainText,
		},
		// Nats:        api.AceOptionsNatsSettings{},
		Platform: api.PlatformSettings{
			AppName: func() string {
				if in.Settings.Platform.Hosted {
					return "ByteBuilders: Kubernetes Native Data Platform"
				}
				return "ACE: Kubernetes Native Data Platform"
			}(),
			RunMode:                         "prod",
			ExperimentalFeatures:            false,
			ForcePrivate:                    false,
			DisableHttpGit:                  false,
			InstallLock:                     true, // TODO: why?
			RepositoryUploadEnabled:         true,
			RepositoryUploadAllowedTypes:    nil,
			RepositoryUploadMaxFileSize:     3,
			RepositoryUploadMaxFiles:        5,
			ServiceEnableCaptcha:            true,
			ServiceRegisterEmailConfirm:     false,
			ServiceDisableRegistration:      false,
			ServiceRequireSignInView:        false,
			ServiceEnableNotifyMail:         true,
			CookieName:                      "i_like_bytebuilders",
			ServerLandingPage:               "home",
			LogMode:                         "console",
			LogLevel:                        "Info", // Trace
			OtherShowFooterBranding:         false,
			OtherShowFooterVersion:          true,
			OtherShowFooterTemplateLoadTime: true,
			EnableCSRFCookieHttpOnly:        false,
		},
		// Stripe:   api.StripeSettings{},
		Security: api.SecuritySettings{
			Oauth2JWTSecret: passgen.GenerateForCharset(43, passgen.AlphaNum),
			CsrfSecretKey:   passgen.GenerateForCharset(64, passgen.AlphaNum),
		},
		// Searchlight: api.SearchlightSettings{},
		Grafana: api.GrafanaSettings{
			AppMode: "production",
		},
	}
	if in.Settings.Platform.Hosted {
		// out.AceOptionsSettings.Stripe = api.StripeSettings{}
		// out.AceOptionsSettings.Searchlight: api.SearchlightSettings{},
	}

	return nil
}

func tplPlatformTLSSecret(in *api.AceOptionsSpec) string {
	return fmt.Sprintf("%s-cert", in.Release.Name)
}

func tplNATSCredSecret(in *api.AceOptionsSpec) string {
	return fmt.Sprintf("%s-nats-cred", in.Release.Name)
}

func tplNATSTLSSecret(in *api.AceOptionsSpec) string {
	if in.Nats.ExposeVia == api.ServiceTypeLoadBalancer {
		return fmt.Sprintf("%s-nats-cert", in.Release.Name)
	}
	return fmt.Sprintf("%s-cert", in.Release.Name)
}

func tplPlatformConfig(in *api.AceOptionsSpec) string {
	return fmt.Sprintf("%s-config", in.Release.Name)
}

func getBucket(bucket string, elem ...string) (string, error) {
	u, err := url.Parse(bucket)
	if err != nil {
		return "", err
	}
	u.Path = path.Join(append([]string{u.Path}, elem...)...)
	return u.String(), nil
}

func mustBucketName(bucket string, elem ...string) string {
	if name, err := getBucket(bucket, elem...); err != nil {
		panic(errors.Wrap(err, "failed to generate bucket name"))
	} else {
		return name
	}
}
