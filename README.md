# gh-cmcm

gh-cmcm (commit comment) is a [gh](https://github.com/cli/cli) extension which comments to a commit by [GitHub API](https://docs.github.com/en/rest).

## Installation

```
gh extension install johnmanjiro13/gh-cmcm
```

## Usage

### Authorization

You need to set your personal access token of github to `GITHUB_TOKEN` environment variable.

If you use github enterprise, you need to set your api base url to `GITHUB_BASE_URL` environment variable.

### Create a comment
```
gh cmcm create <commit_sha> --body 'Comment by cli'
```

### Get a comment
```
gh cmcm get <comment_id>
```
You can use the `--json` flag if you want to get result as json.

### List comments
```
gh cmcm list <comment_id>
```
You can use the `--json` flag if you want to get result as json.

### Update a comment
```
gh cmcm update <comment_id> --body 'Updated comment by cli'
```

### Delete a comment
```
gh cmcm delete <comment_id>
```
