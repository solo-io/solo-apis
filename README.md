# Solo.io APIs
This repository hosts the Solo.io API definitions. It is intended as a read-only mirror; the source of truth for the 
APIs are the projects that own them. The goal of this repository is to allow for projects to consume an API 
without having to depend on the project that owns it.

## Gloo Edge APIs
For each release of Gloo Edge, we want to freeze the solo-apis dependency by creating a unique tag tied to the dependent release. For example, when Gloo Edge `v1.8.4` is released, a corresponding `v1.8.4` tag is created in the Gloo Edge repository. We want to create a `gloo-v1.8.4` tag in solo-apis with a read-only mirror of the API.

In order to keep API definitions in sync, we automate the process of syncing long term support (LTS) branches. To keep the LTS branches unique, we prefix them with the dependent repo name (ie `gloo`). For example, in gloo the `v1.8.x` LTS branch is tracked by the `gloo-v1.8.x` LTS branch in solo-apis.

### How do Gloo Edge APIs get synced?

1. A Gloo release occurs (ie v1.8.1 is released from LTS branch v1.8.x)
1. A Gloo GHA generates the new APIs and pushes those to a branch (ie sync-apis/gloo-v1.8.x/gloo-v1.8.1)
1. A [Solo-Apis GHA](.github/workflows/README.md#create-pr-for-lts-branch) notices that a commit was pushed to a `sync-apis` prefixed branch, and creates a PR.
1. **That PR is manually reviewed by a developer, and eventually approved**
1. A [Solo-Apis GHA](.github/workflows/README.md/#tag-commit-on-lts-branch) notices that a commit is pushed to a LTS branch (ie `gloo-v1.8.x`) and tags the commit. This tag is referenced by other repositories.

### A Gloo Edge release occurred, but the Gloo GHA didn't run, how do I manually sync the APIs?
As noted above, the automated process is kicked of by a Gloo Edge release. In the case where the appropriate GitHub action doesn't run or fails, we need to manually complete what that GitHub action does.

1. In gloo, checkout the release tag to sync (`git checkout v1.8.4`)
1. In solo-apis, checkout corresponding LTS branch (`git checkout gloo-v1.8.x`) and pull down the latest changes
1. In solo-apis, checkout a new branch, following the format `sync-apis/{TARGET_LTS_BRANCH}/{TARGET_TAG}`. This is a special format that is used by the solo-apis GHA's to create the appropriate PR. More information can be found [here](.github/workflows/README.md#create-pr-for-lts-branch)
1. Execute Sync script (`./hack/sync-gloo-apis.sh; make generate -B`)
1. Commit those changes with a useful commit message, and push to a remote branch with the same name as the local branch.
1. Wait a minute or two, and notice that a Pull Request is automatically opened in the solo-apis repo.
1. At this point, we have reached parity with the automated process. All that remains is to have the PR reviewed, and once it is merged, it will be tagged automatically.
