module github.com/solo-io/solo-apis

go 1.26.1

require (
	github.com/cncf/xds/go v0.0.0-20251022180443-0feb69152e9f
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/evanphx/json-patch/v5 v5.9.11 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.7.0-rc.1 // indirect
	github.com/golang/protobuf v1.5.4
	github.com/onsi/ginkgo/v2 v2.26.0 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.1 // indirect
	github.com/rotisserie/eris v0.5.4
	github.com/solo-io/go-utils v0.29.0 // indirect
	github.com/solo-io/protoc-gen-ext v0.1.0
	github.com/solo-io/skv2 v0.45.0
	github.com/solo-io/solo-kit v0.40.0
	github.com/spf13/pflag v1.0.10 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.1 // indirect
	golang.org/x/exp v0.0.0-20251209150349-8475f28825e9 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/oauth2 v0.34.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	google.golang.org/grpc v1.78.0
	google.golang.org/protobuf v1.36.11
	istio.io/api v1.28.0-beta.1
	k8s.io/api v0.34.2
	k8s.io/apimachinery v0.34.2
	k8s.io/client-go v0.34.2
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20250820121507-0af2bda4dd1d // indirect
	sigs.k8s.io/controller-runtime v0.22.4
	sigs.k8s.io/yaml v1.6.0 // indirect
)

require (
	go.uber.org/mock v0.6.0
	google.golang.org/genproto/googleapis/api v0.0.0-20251202230838-ff82c1b0f217
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/gateway-api v1.4.0
)

require (
	github.com/envoyproxy/go-control-plane/envoy v1.35.1-0.20251010000747-50f0b197aa55
	github.com/kgateway-dev/kgateway/v2 v2.1.1
	github.com/solo-io/gloo v1.17.15
	github.com/solo-io/gloo-operator v0.1.0-rc.2
)

require (
	cel.dev/expr v0.24.0 // indirect
	github.com/fxamacker/cbor/v2 v2.9.0 // indirect
	github.com/google/gnostic-models v0.7.0 // indirect
	github.com/google/gofuzz v1.2.1-0.20210504230335-f78f29fc09ea // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240409071808-615f978279ca // indirect
	github.com/x448/float16 v0.8.4 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.39.0 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.13.0 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v6 v6.3.0 // indirect
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.13.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-openapi/jsonpointer v0.21.2 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/procfs v0.17.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/term v0.40.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	golang.org/x/time v0.14.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.5.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250814151709-d7b6acb124c3 // indirect
	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
)

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2

	github.com/distribution/reference => github.com/distribution/reference v0.5.0

	// We currently rely on NetworkInspectOptions which is deprecated and has been removed from more recent versions
	github.com/docker/docker => github.com/docker/docker v27.5.0+incompatible

	// anyvendor (used in check-code-gen in CI) requires version v5.3.0, but solo-io/go-utils overrides it
	github.com/go-git/go-git/v5 => github.com/go-git/go-git/v5 v5.3.0

	// Newer versions of strcase change the field names for generated code.
	github.com/iancoleman/strcase => github.com/iancoleman/strcase v0.2.0

	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

	// skv2 relies on v0.7.0 and will break on v0.13.0
	// jsonschema v0.13.0 is an indirect dependency of kgateway, via the goreleaser tool and can safely be replaced
	github.com/invopop/jsonschema => github.com/invopop/jsonschema v0.7.0

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.5.1

	// Resolves ambigious import error.
	github.com/ugorji/go => github.com/ugorji/go v1.2.12
	go.opentelemetry.io/collector/component => go.opentelemetry.io/collector/component v0.118.0
	go.opentelemetry.io/collector/exporter => go.opentelemetry.io/collector/exporter v0.116.0

	// Once otel is bumped off v0.116, this can be removed
	go.opentelemetry.io/collector/pdata => go.opentelemetry.io/collector/pdata v1.22.0
	go.opentelemetry.io/collector/pdata/pprofile => go.opentelemetry.io/collector/pdata/pprofile v0.116.0

	// Rate-limiter v0.16.4 expects otelgrpc v0.58.0 API
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.58.0

	// Use prometheus exporter v0.60.0 which properly handles the new otlptranslator API
	go.opentelemetry.io/otel/exporters/prometheus => go.opentelemetry.io/otel/exporters/prometheus v0.60.0

	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d

	// This is necessary while we still depend on the IstioOperator and related apis, which were removed in 1.24
	// This can be removed once we drop ILM v2 support in GME.
	istio.io/api/123 => istio.io/api v1.23.4

	// Use solo's Istio fork, which has actual properly-tagged versions unlike the parent repo
	istio.io/istio => github.com/solo-io/istio v1.28.0-beta.1-solo
)

exclude github.com/knadh/koanf/maps v0.1.2

exclude github.com/knadh/koanf/providers/confmap v1.0.0
