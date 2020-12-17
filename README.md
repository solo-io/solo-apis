# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

## Gloo Apis
Syncing the Gloo APIs is currently a manual process and can be completed with 2 extra steps 
before running `make generate`
1. In the gloo repo, checkout the desired gloo tag with `git checkout tags/[GLOO_TAG]`
2. Ensure that the gloo repo is available at `../`
3. Run `./hack/sync-gloo-apis.sh; make generate -B`. You may need to run `make mod-download update-deps` if this is your first time syncing the api.

After running make generate a PR should be created against solo-apis with the tag of Gloo being
used. Once the PR has merged a tag should be created off of master correspnding to the Gloo tag 
which was used.
