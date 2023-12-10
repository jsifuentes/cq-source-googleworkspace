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

### Setting up your Google OAuth Project

1. Create a new project in the
   [Google Cloud Console](https://console.cloud.google.com/)
1. Go to APIs and Services and enable the "Admin SDK API"
1. Go to the "Credentials" page
1. Create a new OAuth Client ID for a Desktop app
1. Make note of your client ID and client secret. Use this in your spec
   configuration.

### Source Configuration

The following source configuration file will sync to a PostgreSQL database. See
[the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more
information on how to configure the source and destination.

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
