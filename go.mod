module github.com/solo-io/solo-apis

go 1.14

replace (
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.1
	k8s.io/client-go => k8s.io/client-go v0.17.1
)

require (
	github.com/envoyproxy/go-control-plane v0.9.1
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4-0.20200406172829-6d816de489c1
	github.com/mitchellh/hashstructure v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/solo-io/go-utils v0.16.5
	github.com/solo-io/protoc-gen-ext v0.0.9
	github.com/solo-io/skv2 v0.7.1
	github.com/solo-io/solo-kit v0.13.9
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6
	google.golang.org/grpc v1.26.0
	k8s.io/apimachinery v0.18.4
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/code-generator v0.18.4
	sigs.k8s.io/controller-runtime v0.6.1
)
