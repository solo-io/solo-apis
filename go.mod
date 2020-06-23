module github.com/solo-io/solo-apis

go 1.14

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309

require (
	github.com/solo-io/skv2 v0.4.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v8.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.5.1
)
