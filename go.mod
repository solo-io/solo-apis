module github.com/solo-io/solo-apis

go 1.24.5

require (
	github.com/cncf/xds/go v0.0.0-20250121191232-2f005788dc42
	github.com/envoyproxy/go-control-plane v0.13.4
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/evanphx/json-patch v5.9.11+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.9.11 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/golang/protobuf v1.5.4
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.22.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.62.0 // indirect
	github.com/rotisserie/eris v0.5.4
	github.com/solo-io/go-utils v0.28.4 // indirect
	github.com/solo-io/protoc-gen-ext v0.1.0
	github.com/solo-io/skv2 v0.43.3
	github.com/solo-io/solo-kit v0.35.0
	github.com/spf13/pflag v1.0.6 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/oauth2 v0.28.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
	istio.io/api v1.22.7-0.20241205190107-5d7b98128323
	k8s.io/api v0.33.2
	k8s.io/apiextensions-apiserver v0.33.2 // indirect
	k8s.io/apimachinery v0.33.2
	k8s.io/client-go v0.33.2
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
	sigs.k8s.io/controller-runtime v0.20.1
	sigs.k8s.io/yaml v1.4.0 // indirect
)

require (
	go.uber.org/mock v0.4.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250218202821-56aae31c358a
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	cel.dev/expr v0.19.1 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/google/gnostic-models v0.6.9 // indirect
	github.com/google/gofuzz v1.2.1-0.20210504230335-f78f29fc09ea // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.12.1 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff // indirect
	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.6.0 // indirect
)

replace (
	cloud.google.com/go/storage => cloud.google.com/go/storage v1.42.0
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2

	github.com/distribution/reference => github.com/distribution/reference v0.5.0

	// Bumping helm to v3.18.4 caused this dependency to be bumped to v0.13.0,
	// which is not compatible with Istio 1.22, so we need to pin this.
	github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.12.1-0.20240415211714-57c85e1829e6

	// anyvendor (used in check-code-gen in CI) requires version v5.3.0, but solo-io/go-utils overrides it
	github.com/go-git/go-git/v5 => github.com/go-git/go-git/v5 v5.3.0

	// Newer versions of strcase change the field names for generated code.
	github.com/iancoleman/strcase => github.com/iancoleman/strcase v0.2.0

	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.5.1

	// Resolves ambigious import error.
	github.com/ugorji/go => github.com/ugorji/go v1.2.12

	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d

	google.golang.org/grpc => google.golang.org/grpc v1.68.1

	// Use solo's Istio fork, which has actual properly-tagged versions unlike the parent repo
	istio.io/istio => github.com/solo-io/istio v1.22.8-patch1-solo
)
