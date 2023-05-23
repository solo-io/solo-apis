# GitHub Workflows

## Tag Commit on LTS Branch

After PRs are merged into LTS branches, we:
1. Extract the latest commit message (which the above workflow provided)
1. Determine the tag name from the commit message
1. Push a tag name for that commit

When wanting to release a new tag for this branch, the PR name has to follow the following format: `Solo APIs. @tag-name=<tag>`.
This will be released as the tag: `sa-<tag>`. We do not follow any strict tagging for this as it is intended for sole use by
Gloo OSS, but make sure to check for existing tags with `sa-` to avoid collision.