# Gloo Mesh APIs
This repository hosts the Solo.io API definitions for **Gloo Mesh Enterprise**. It is intended as a read-only mirror;
the source of truth for the APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API without having to depend on the project that owns it.

## Gloo Mesh APIs
For each release of Gloo Mesh, we want to freeze the solo-apis dependency by creating a unique tag tied to the dependent release. For example, when Gloo Mesh Enterprise `vX.Y.Z` is released, a corresponding `vX.Y.Z` tag is created in the Gloo Mesh Enterprise repository. We create a `gloo-mesh-vX.Y.Z` tag in solo-apis with a read-only mirror of the API.

## Consuming Gloo Mesh APIs
In order to use this repo in your own code, you'll need to pin the following dependencies to specific versions:

- github.com/solo-io/solo-apis
- k8s.io/api
- k8s.io/apimachinery
- k8s.io/client-go

Get the appropriate version for the `github.com/solo-io/solo-apis` module by running the following command:

```
$ go get github.com/solo-io/solo-apis@gloo-mesh-v2.4.x
$ go list -m -json github.com/solo-io/solo-apis@gloo-mesh-v2.4.x | jq -r '.Version'
v1.6.32-0.20230530105146-faaf4db0b31d
```

A sample go.mod file includes the following:

```
module github.com/my-org/my-repository

go 1.19

require (
	github.com/solo-io/solo-apis v1.6.32-0.20230530105146-faaf4db0b31d // replace with go list output above
)

replace (
	k8s.io/api => k8s.io/api v0.24.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.24.8
	k8s.io/client-go => k8s.io/client-go v0.24.8
)
```

## Examples
The [examples directory](./examples) contains examples that cover various use cases and functionality for client-go for
interacting with the Kubernetes API and Gloo Mesh objects.

> **Note**
> Before running the examples, make sure your Go version matches the Go version in the [go.mod](./go.mod) file.
