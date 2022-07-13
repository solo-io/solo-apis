module github.com/solo-io/solo-apis

go 1.16

replace (
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

	// skv2 uses a newer version than the imported solo-kit version which causes issues. Replaces the version with the solo-kit version
	github.com/pseudomuto/protoc-gen-doc => github.com/pseudomuto/protoc-gen-doc v1.0.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.9
	k8s.io/client-go => k8s.io/client-go v0.20.9
)

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/mitchellh/hashstructure v1.0.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.16.0
	github.com/pkg/errors v0.9.1
	github.com/rotisserie/eris v0.4.0
	github.com/solo-io/go-utils v0.21.24
	github.com/solo-io/protoc-gen-ext v0.0.16
	github.com/solo-io/skv2 v0.21.6
	github.com/solo-io/solo-kit v0.30.0
	golang.org/x/tools v0.1.5
	google.golang.org/genproto v0.0.0-20210416161957-9910b6c460de
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	istio.io/tools v0.0.0-20200918020118-6d0a0e49b5d3
	k8s.io/apimachinery v0.20.9
	k8s.io/client-go v0.21.0
	k8s.io/code-generator v0.20.9
	sigs.k8s.io/controller-runtime v0.7.0
)

exclude (
	// Exclude pre-go-mod kubernetes tags, because they are older
	// than v0.x releases but are picked when updating dependencies.
	k8s.io/client-go v1.4.0
	k8s.io/client-go v1.5.0
	k8s.io/client-go v1.5.1
	k8s.io/client-go v1.5.2
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/client-go v2.0.0+incompatible
	k8s.io/client-go v2.0.0-alpha.1+incompatible
	k8s.io/client-go v3.0.0+incompatible
	k8s.io/client-go v3.0.0-beta.0+incompatible
	k8s.io/client-go v4.0.0+incompatible
	k8s.io/client-go v4.0.0-beta.0+incompatible
	k8s.io/client-go v5.0.0+incompatible
	k8s.io/client-go v5.0.1+incompatible
	k8s.io/client-go v6.0.0+incompatible
	k8s.io/client-go v7.0.0+incompatible
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/client-go v9.0.0+incompatible
	k8s.io/client-go v9.0.0-invalid+incompatible
)
