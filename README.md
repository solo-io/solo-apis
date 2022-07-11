# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

## Gloo Mesh APIs
For each release of Gloo Mesh, we want to freeze the solo-apis dependency by creating a unique tag tied to the dependent release. For example, when Gloo Mesh Enterprise `v2.1.0-beta10` is released, a corresponding `v2.1.0-beta10` tag is created in the Gloo Mesh Enterprise repository. We create a `gloo-mesh-v2.1.0-beta10` tag in solo-apis with a read-only mirror of the API.

## Consuming Gloo Mesh APIs
In order to use this repo in your own code, you'll need to pin the following dependencies to specific versions:
- `github.com/solo-io/cue@v0.4.1-0.20210623143425-308aee4ff092`
- `k8s.io/client-go@v0.19.6`

A sample go.mod file includes the following:
```
replace (
    cuelang.org/go => github.com/solo-io/cue v0.4.1-0.20210623143425-308aee4ff092
    k8s.io/client-go => k8s.io/client-go v0.19.6
)
```