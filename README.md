# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

## Gloo Edge APIs
For each release of Gloo Edge, we want to freeze the solo-apis dependency by creating a unique tag tied to the dependent release. For example, when Gloo Edge `v1.8.4` is released, a corresponding `v1.8.4` tag is created in the Gloo Edge repository. We want to create a `gloo-v1.8.4` tag in solo-apis with a read-only mirror of the API.

In order to keep API definitions in sync, we automate the process of syncing long term support (LTS) branches. To keep the LTS branches unique, we prefix them with the dependent repo name (ie `gloo`). For example, in gloo the `v1.8.x` LTS branch is tracked by the `gloo-v1.8.x` LTS branch in solo-apis.

### How do Gloo Edge APIs get synced?

1. A Gloo release occurs (ie `v1.8.1` is released from LTS branch `v1.8.x`)
1. A Gloo GHA generates the new APIs and pushes those to a branch (ie `sync-apis/gloo-v1.8.x/gloo-v1.8.1`)
1. A [Solo-Apis GHA](.github/workflows/README.md#create-pr-for-lts-branch) notices that a commit was pushed to a `sync-apis` prefixed branch, and creates a PR.
1. **That PR is manually reviewed by a developer, and eventually approved**
1. A [Solo-Apis GHA](.github/workflows/README.md/#tag-commit-on-lts-branch) notices that a commit is pushed to a LTS branch (ie `gloo-v1.8.x`) and tags the commit. This tag is referenced by other repositories.

### A Gloo Edge release occurred, but the Gloo GHA didn't run, how do I trigger the GHA manually?
As noted above, the automated process is typically triggered by a Gloo Edge release. In the case where the appropriate GitHub action doesn't run or fails, we need to manually trigger the Github action.

In the [Gloo workflow to push API changes to solo-apis](https://github.com/solo-io/gloo/actions/workflows/push-solo-apis-branch.yaml) there is an option to `Run Workflow`.

1. Keep the branch set to main
1. Set `Release Tag Name` to the Gloo release tag (ie `v1.8.4`)
1. Set `Release Branch` to the Gloo branch used for that release (ie `v1.8.x`)
1. Select `Run Workflow`. When the workflow completes, a PR should be opened in solo-apis. This PR will need to be approved by a developer, but once it merges will be tagged automatically.