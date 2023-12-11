# CloudQuery Google Workspace Source Plugin

[![test](https://github.com/jsifuentes/cq-source-googleworkspace/actions/workflows/test.yaml/badge.svg)](https://github.com/jsifuentes/cq-source-googleworkspace/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsifuentes/cq-source-googleworkspace)](https://goreportcard.com/report/github.com/jsifuentes/cq-source-googleworkspace)

A Google Workspace source plugin for CloudQuery that loads data from Google
Workspace to any database, data warehouse or data lake supported by
[CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena,
and many more.

## Links

- [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
- [Supported Tables](docs/tables/README.md)

## Configuration

The following source configuration file will sync to a PostgreSQL database. See
[the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more
information on how to configure the source and destination.

If `token_file` is set, after you successfully authenticate via OAuth, your
access token will be written to the `token_file`.

How to find your Google Workspace Customer ID:
https://support.google.com/a/answer/10070793?hl=en

You can get your own OAuth credentials using
[this guide](https://developers.google.com/identity/protocols/oauth2#1.-obtain-oauth-2.0-credentials-from-the-dynamic_data.setvar.console_name-.).

```yaml
kind: source
spec:
  name: "googleworkspace"
  path: "jsifuentes/googleworkspace"
  registry: "github"
  version: "v1.0.0"
  destinations:
    - "postgresql"
  spec:
    customer_id: your Google Workspace Customer ID
    oauth:
      # token_file: ./token.js
      client_id: your Google Cloud Project OAuth Client ID
      client_secret: your Google Cloud Project OAuth Client Secret
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0`
   with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build
the release binaries and create the new release on GitHub. To customize the
release notes, see the Go releaser
[changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
