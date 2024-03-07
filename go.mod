module github.com/solo-io/solo-apis

go 1.22.0

require (
	github.com/cncf/xds/go v0.0.0-20231016030527-8bd2eac9fb4a
	github.com/envoyproxy/go-control-plane v0.11.2-0.20231120184409-3aaf8b9a228b
	github.com/envoyproxy/protoc-gen-validate v1.0.2
	github.com/evanphx/json-patch v5.7.0+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.7.0 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/onsi/ginkgo/v2 v2.15.0 // indirect
	github.com/onsi/gomega v1.31.1 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/rotisserie/eris v0.5.4
	github.com/solo-io/cue v0.4.7
	github.com/solo-io/go-utils v0.24.8 // indirect
	github.com/solo-io/protoc-gen-ext v0.0.18
	github.com/solo-io/skv2 v0.36.1
	github.com/solo-io/solo-kit v0.34.0
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/oauth2 v0.14.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	istio.io/api v1.20.3-0.20240116015448-5563f7225778
	k8s.io/api v0.28.3
	k8s.io/apimachinery v0.28.3
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/controller-runtime v0.16.3
	sigs.k8s.io/yaml v1.4.0 // indirect
)

require google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/component-base v0.28.3 // indirect
)

require (
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231030173426-d783a09b4405 // indirect
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20230926050212-f7f687d19a98 // indirect
	github.com/imdario/mergo v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.4.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.4.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.3.0 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2

	// anyvendor (used in check-code-gen in CI) requires version v5.3.0, but solo-io/go-utils overrides it
	github.com/go-git/go-git/v5 => github.com/go-git/go-git/v5 v5.3.0

	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.5.1

	// needed by Istio 1.20.1:
	go.opentelemetry.io/otel/exporters/prometheus => go.opentelemetry.io/otel/exporters/prometheus v0.39.1-0.20230714155235-03b8c47770f2 // https://github.com/istio/istio/blob/552626bd81f625917bcba6e415cae5aa3d5c61af/go.mod#L86
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v0.39.0 // https://github.com/istio/istio/blob/552626bd81f625917bcba6e415cae5aa3d5c61af/go.mod#L89

	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d

	// k8s dependencies need to be pinned to avoid conflicts with wasm-kit@v0.1.4 transitive dependencies:
	// https://github.com/solo-io/solo-apis/issues/13786
	k8s.io/api => k8s.io/api v0.28.3
	k8s.io/client-go => k8s.io/client-go v0.28.3

	// TODO(samu): pin controller-runtime dependency to a non-published version that includes this commit:
	// https://github.com/kubernetes-sigs/controller-runtime/commit/67d355d101a6a37bae7c3deb0988e112b49c18ad
	// once the v0.16.4 release is cut, remove this replacement.
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.16.4-0.20231208143029-67d355d101a6
)
