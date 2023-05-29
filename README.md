# Gloo Mesh APIs
This repository hosts the Solo.io API definitions for **Gloo Mesh Enterprise**. It is intended as a read-only mirror;
the source of truth for the APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API without having to depend on the project that owns it.

## Gloo Mesh APIs
For each release of Gloo Mesh, we want to freeze the solo-apis dependency by creating a unique tag tied to the dependent release. For example, when Gloo Mesh Enterprise `vX.Y.Z` is released, a corresponding `vX.Y.Z` tag is created in the Gloo Mesh Enterprise repository. We create a `gloo-mesh-vX.Y.Z` tag in solo-apis with a read-only mirror of the API.

## Examples
The [examples directory](./examples) contains examples that cover various use cases and functionality for client-go for
interacting with the Kubernetes API and Gloo Mesh objects.

> **Note**
> Before running the examples, make sure your Go version matches the Go version in the [go.mod](./go.mod) file.
