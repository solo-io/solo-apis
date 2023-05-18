module github.com/solo-io/solo-apis

go 1.16

replace (
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.0.0
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
	// todo (fabian) brought from v0.17.17 => v0.30.0. should we backport bumps to skv2 0.17.X? would that be a big hassle?
	github.com/solo-io/skv2 v0.30.0
	// todo (fabian) update solokit, then bring it here?
	github.com/solo-io/solo-kit v0.23.0
	golang.org/x/tools v0.6.0
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	istio.io/tools v0.0.0-20200918020118-6d0a0e49b5d3
	k8s.io/apimachinery v0.26.4
	k8s.io/client-go v0.26.4
	k8s.io/code-generator v0.26.4
	sigs.k8s.io/controller-runtime v0.14.6
)
