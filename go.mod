module github.com/solo-io/solo-apis

go 1.14

replace (
	// pinned to solo-io's fork of cue version 308aee4ff0928a8e0ec25b9cbbdc445264038463
	// note(ilackarms): this replace must be shared in any skv2-based go module due to incompatibility with upstream versions of cue
	cuelang.org/go => github.com/solo-io/cue v0.4.1-0.20210623143425-308aee4ff092

	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	k8s.io/apimachinery => k8s.io/apimachinery v0.26.4
	k8s.io/client-go => k8s.io/client-go v0.26.4
)

require (
	github.com/envoyproxy/go-control-plane v0.10.2-0.20220325020618-49ff273808a1
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/mitchellh/hashstructure v1.0.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.26.0
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.1.1
	github.com/solo-io/go-utils v0.21.6
	github.com/solo-io/protoc-gen-ext v0.0.18
	github.com/solo-io/skv2 v0.29.3-0.20230322181416-98c3ab6b87b6
	github.com/solo-io/solo-kit v0.21.0
	golang.org/x/tools v0.6.0
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	istio.io/tools v0.0.0-20200918020118-6d0a0e49b5d3
	k8s.io/apimachinery v0.26.4
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/code-generator v0.26.1
	sigs.k8s.io/controller-runtime v0.14.6
)
