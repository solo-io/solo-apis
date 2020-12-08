module github.com/solo-io/solo-apis

go 1.14

replace (
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.6
	k8s.io/client-go => k8s.io/client-go v0.18.6
)

require (
	github.com/envoyproxy/go-control-plane v0.9.6-0.20200529035633-fc42e08917e9
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4
	github.com/mitchellh/hashstructure v1.0.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/solo-io/go-utils v0.20.0
	github.com/solo-io/k8s-utils v0.0.2 // indirect
	github.com/solo-io/protoc-gen-ext v0.0.9
	github.com/solo-io/skv2 v0.13.5
	github.com/solo-io/solo-kit v0.13.9
	golang.org/x/tools v0.0.0-20201027213030-631220838841
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.30.0
	istio.io/tools v0.0.0-20200918020118-6d0a0e49b5d3
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/code-generator v0.18.6
	sigs.k8s.io/controller-runtime v0.6.2
)
