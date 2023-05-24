# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

## Gloo Mesh APIs
For each release of Gloo Mesh, we want to freeze the solo-apis dependency by creating a unique tag tied to the dependent release. For example, when Gloo Mesh Enterprise `vX.Y.Z` is released, a corresponding `vX.Y.Z` tag is created in the Gloo Mesh Enterprise repository. We create a `gloo-mesh-vX.Y.Z` tag in solo-apis with a read-only mirror of the API.
