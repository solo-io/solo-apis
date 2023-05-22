# Solo.io APIs
`gloo-repo-branch` serves to provide the APIs required for Gloo OSS, such as Rate Limiter APIs. Instead of using Gloo LTS
releases to update APIs, this branch is used by Gloo to get updated API information. Therefore we do not need to sync the APIs. 
It is not intended to be updated as much as the regular LTS Solo APIs.

This branch stemmed off of the [gloo-edge-safe-hasher](https://github.com/solo-io/solo-apis/tree/gloo-edge-safe-hasher) tag which is used
in Gloo OSS due to its minimized hashing which prevents data races in gloo. [Original Issue](https://github.com/solo-io/gloo/issues/7209).