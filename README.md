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

### A Gloo Edge release occurred, but the Gloo GHA didn't run, how do I manually sync the APIs?
As noted above, the automated process is kicked off by a Gloo Edge release. In the case where the appropriate GitHub action doesn't run or fails, we need to manually complete what that GitHub action does.

*These steps assume you have a local copy of `gloo` and `solo-apis` repositories located under the same parent directory.*

1. In gloo, checkout the release tag to sync (`git checkout v1.8.4`)
1. In solo-apis, checkout corresponding LTS branch (`git checkout gloo-v1.8.x`) and pull down the latest changes
1. In solo-apis, checkout a new branch, following the format `sync-apis/{TARGET_LTS_BRANCH}/{TARGET_TAG}` (ie `sync-apis/gloo-v1.8.x/gloo-v1.8.4`). This is a special format that is used by the solo-apis GHA's to create the appropriate PR. More information can be found [here](.github/workflows/README.md#create-pr-for-lts-branch)
1. Execute Sync script (`./hack/sync-gloo-apis.sh; make generate -B`)
1. Commit those changes with a useful commit message, and push to a remote branch with the same name as the local branch.
1. Wait a minute or two, and notice that a Pull Request is automatically opened in the solo-apis repo.
1. At this point, we have reached parity with the automated process. All that remains is to have the PR reviewed, and once it is merged, it will be tagged automatically.

### How do I release the Gloo Edge API sync results?
This repo does not require an ordinary "release" and instead publishes tags that represent the API at a point in time.

If you followed the previous steps and either the automated or manual sync completed, once that PR is merged, it should automatically be tagged. [A GitHub Action](.github/workflows/README.md/#tag-commit-on-lts-branch) is responsible for executing this.

If this GitHub Action did not complete for some reason, you'll need to manually publish a tag. Below we outline the steps taken to release a tag for a Gloo `v1.8.4` release:

1. Checkout the branch that contains the updated API. (`git checkout gloo-v1.8.x`)
1. Pull down the latest changes (`git pull`)
1. Create an annotated tag (`git tag -a gloo-v1.8.4 -m "Sync Gloo APIs v1.8.4"`)
1. Push the tag (`git push origin gloo-v1.8.4`)