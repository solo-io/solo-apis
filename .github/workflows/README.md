# GitHub Workflows

# Syncing LTS Branches

Solo-Apis is a public, read-only mirror for API definitions. In order to keep API definitions in sync, we automate the process of syncing long term support (LTS) branches. 

To keep the LTS branches unique, we prefix them with the name of the owner repo. For example, in gloo-edge the `1.8.x` LTS branch is tracked by the `gloo-1.8.x` LTS branch in solo-apis.

### How do Gloo Edge APIs get synced?

1. A Gloo release occurs (ie v1.8.1 is released from LTS branch v1.8.x)
1. A Gloo GHA generates the new APIs and pushes those to a branch (ie sync-apis/gloo-v1.8.x/gloo-v1.8.1)
1. A Solo-Apis GHA (defined on main branch) notices that a commit was pushed to a `sync-apis` prefixed branch, and creates a PR.
1. That PR is manually reviewed by a developer, and eventually approved
1. A [Solo-Apis GHA](#tag-commit-on-lts-branch) notices that a commit is pushed to a LTS branch (ie gloo-v1.8.x) and tags the commit. This tag is referenced by other repositories.

The GitHub actions on this branch that enable this process are defined below:

## Tag Commit on LTS Branch

After PRs are merged into LTS branches, we:
1. Extract the latest commit message (which the above workflow provided)
1. Determine the tag name from the commit message
1. Push a tag name for that commit
