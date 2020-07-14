# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

## Gloo Apis
Syncing the Gloo APIs is currently a manual process and can be completed with 2 extra steps 
before running `make generate`
1. Ensure gloo is checked out to master and available at `../`
2. run `./hack/sync-apis.sh`