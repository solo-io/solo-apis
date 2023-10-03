module github.com/solo-io/solo-apis

go 1.21.1

require (
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4
	github.com/envoyproxy/go-control-plane v0.11.1-0.20230524094728-9239064ad72f
	github.com/envoyproxy/protoc-gen-validate v0.10.1
	github.com/evanphx/json-patch/v5 v5.6.0 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/onsi/ginkgo/v2 v2.11.0 // indirect
	github.com/onsi/gomega v1.27.10 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/rotisserie/eris v0.5.4
	github.com/solo-io/cue v0.4.7
	github.com/solo-io/go-utils v0.24.6 // indirect
	github.com/solo-io/protoc-gen-ext v0.0.18
	github.com/solo-io/skv2 v0.34.5
	github.com/solo-io/solo-kit v0.33.0
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.25.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/oauth2 v0.9.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	istio.io/api v0.0.0-20230803063743-0b7b8bc2e0bb
	k8s.io/api v0.28.1
	k8s.io/apiextensions-apiserver v0.28.1 // indirect
	k8s.io/apimachinery v0.28.1
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/utils v0.0.0-20230406110748-d93618cff8a2 // indirect
	sigs.k8s.io/controller-runtime v0.15.2
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require google.golang.org/genproto/googleapis/api v0.0.0-20230726155614-23370e0ffb3e

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/component-base v0.28.1 // indirect
)

require (
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20230228050547-1710fef4ab10 // indirect
	github.com/imdario/mergo v0.3.15 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/procfs v0.11.0 // indirect
	golang.org/x/term v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.3.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/docker/docker => github.com/docker/docker v23.0.1+incompatible
	github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.11.1-0.20230416233444-7f2a3030ef40

	// anyvendor (used in check-code-gen in CI) requires version v5.3.0, but solo-io/go-utils overrides it
	github.com/go-git/go-git/v5 => github.com/go-git/go-git/v5 v5.3.0
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3

	github.com/prometheus/common => github.com/prometheus/common v0.42.0

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.5.1

	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d

	// Using grpc v1.54.0 since pkg/meshctl/commands/generate/graphql/grpc/grpc.go imports
	// hjump/protoreflect which imports a package test/grpc_testing which has been moved to interop/grpc_testing
	// in the grpc repo - https://github.com/grpc/grpc-go/pull/6164 since grpc v1.55.0.
	// This causes v1.55.0+ to not be backwards compatible with v1.54.0 and below if the grpc_testing package is imported.
	// Pinned to v1.54.0 until either hjump/protoreflect is updated to grpc v1.55.0+
	// or this is fixed in a newer version of grpc.
	google.golang.org/grpc => google.golang.org/grpc v1.54.0

	// 1.18
	istio.io/istio => istio.io/istio v0.0.0-20230824050558-a835e9cca9eb

	k8s.io/client-go => k8s.io/client-go v0.28.1

	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.15.0
)
