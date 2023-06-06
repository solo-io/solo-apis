module github.com/solo-io/solo-apis

go 1.19

require (
	github.com/cncf/xds/go v0.0.0-20220314180256-7f1daf1720fc
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220420171917-689c2bccf0ec
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/mitchellh/hashstructure v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.5.4
	github.com/solo-io/cue v0.4.6
	github.com/solo-io/protoc-gen-ext v0.0.18
	github.com/solo-io/skv2 v0.24.1
	github.com/solo-io/solo-kit v0.30.2
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1
	istio.io/api v0.0.0-20220817194812-9ed7a41198c8
	k8s.io/api v0.25.2
	k8s.io/apimachinery v0.25.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	sigs.k8s.io/controller-runtime v0.12.1
)

require (
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/go-logr/zapr v1.2.3 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.1 // indirect
	github.com/onsi/gomega v1.26.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.1-0.20210607210712-147c58e9608a // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/solo-io/go-utils v0.22.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/exp v0.0.0-20220921164117-439092de6870 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/oauth2 v0.0.0-20221014153046-6fdb5e3db783 // indirect
	golang.org/x/sys v0.4.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apiextensions-apiserver v0.25.2 // indirect
	k8s.io/klog/v2 v2.70.2-0.20220707122935-0990e81f1a8f // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require (
	github.com/Azure/go-autorest/autorest v0.11.27 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.20 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	k8s.io/component-base v0.24.1 // indirect
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 //indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/term v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.2.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	istio.io/gogo-genproto v0.0.0-20220413184606-e76735307e2d // indirect
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	// github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.10.2-0.20211130161932-f62def555c97

	// anyvendor (used in check-code-gen in CI) requires version v5.3.0, but solo-io/go-utils overrides it
	github.com/go-git/go-git/v5 => github.com/go-git/go-git/v5 v5.3.0
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.5.1

	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d
	istio.io/api => istio.io/api v0.0.0-20220512212136-561ffec82582

	istio.io/client-go => istio.io/client-go v1.13.9
	istio.io/istio => istio.io/istio v0.0.0-20220606160615-c6abaf5e5fd6
	k8s.io/api => k8s.io/api v0.23.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.7
	k8s.io/apiserver => k8s.io/apiserver v0.23.7
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.7
	k8s.io/client-go => k8s.io/client-go v0.23.7
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.7
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.7
	k8s.io/code-generator => k8s.io/code-generator v0.23.7
	k8s.io/component-base => k8s.io/component-base v0.23.7
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.7
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.7
	k8s.io/cri-api => k8s.io/cri-api v0.23.7
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.7
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.7
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.7
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.7
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.7
	k8s.io/kubectl => k8s.io/kubectl v0.23.7
	k8s.io/kubelet => k8s.io/kubelet v0.23.7

	// Pinned Kubernetes dependencies
	k8s.io/kubernetes => k8s.io/kubernetes v1.23.7
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.7
	k8s.io/metrics => k8s.io/metrics v0.23.7
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.7
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.7
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.7

	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.11.2

	//github.com/googleapis/gnostic  => github.com/google/gnostic v0.6.7

	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
	sigs.k8s.io/gateway-api => sigs.k8s.io/gateway-api v0.4.3
)
