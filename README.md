# Solo.io APIs
`gloo-repo-branch` serves to provide the APIs required for Gloo OSS, such as Rate Limiter APIs. Instead of using Gloo LTS
releases to update APIs, this branch is used by Gloo to get updated API information. Therefore we do not need to sync the APIs.
It is not intended to be updated as much as the regular LTS Solo APIs.

This branch stemmed off of the [gloo-edge-safe-hasher](https://github.com/solo-io/solo-apis/tree/gloo-edge-safe-hasher) tag which is used
in Gloo OSS due to its minimized hashing which prevents data races in gloo. [Original Issue](https://github.com/solo-io/gloo/issues/7209).

Any changes to APIs in [rate-limiter](https://github.com/solo-io/rate-limiter/) should propagate to this branch and code-gen should be re-run. There should be [existing automation](https://github.com/solo-io/rate-limiter/actions/workflows/update-solo-apis.yaml) around this but it is always a good practice to manually sync the protos when bumping k8s dependencies or cutting a major release in case the automation was unsuccessful. Any CRDs generated / modified can be ignored since we only care about the protos.

An alternative that had been considered is using `gloo-main` instead of this branch in Gloo OSS. However this poses an issue when building enterprise - if `gloo-main` is ahead of the latest `gloo-v1.n.x tag` used in enterprise, enterprise will utilize solo-apis@gloo-main instead of the specified tag due to the way go bundles dependencies.

More details on other options considered and their tradeoffs can be found in this [slack tread](https://solo-io-corp.slack.com/archives/G01EERAK3KJ/p1701798077193259)
