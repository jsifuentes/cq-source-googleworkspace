# CloudQuery Google Workspace Source Plugin

[![test](https://github.com/jsifuentes/cq-source-googleworkspace/actions/workflows/test.yaml/badge.svg)](https://github.com/jsifuentes/cq-source-googleworkspace/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsifuentes/cq-source-googleworkspace)](https://goreportcard.com/report/github.com/jsifuentes/cq-source-googleworkspace)

[View this plugin on CloudQuery Hub](https://hub.cloudquery.io/plugins/source/jsifuentes/googleworkspace/)

A Google Workspace source plugin for CloudQuery that loads data from Google
Workspace to any database, data warehouse or data lake supported by
[CloudQuery](https://www.cloudquery.io/), such as sqlite, PostgreSQL, BigQuery, Athena,
and many more.

## Links

- [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
- [Supported Tables](docs/tables/README.md)

## Configuration

The following source configuration file will sync to a sqlite database. See
[the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more
information on how to configure the source and destination.

How to find your Google Workspace Customer ID:
https://support.google.com/a/answer/10070793?hl=en

To authenticate to Google Workspace, you can use either OAuth or a service account.
See the section below for how to configure each.

```yaml
kind: source
spec:
  name: "googleworkspace"
  path: "jsifuentes/googleworkspace"
  registry: "cloudquery"
  version: "v1.2.1"
  destinations:
    - "sqlite"
  spec:
    customer_id: your Google Workspace Customer ID
    # either `oauth` or `service_account` must be provided.
    oauth:
      client_id: your OAuth client ID
      client_secret: your OAuth client secret
      # token_file: ./token.json

    # or
    service_account:
      json_string: '{"type": "service_account","project_id": "...", ...}'
      impersonate_email: email@yourdomain.com
```

### OAuth

You can get your own OAuth credentials using
[this guide](https://developers.google.com/identity/protocols/oauth2#1.-obtain-oauth-2.0-credentials-from-the-dynamic_data.setvar.console_name-.).
When creating your OAuth Client ID, you should select "Desktop app". You also need to enable the Admin SDK API for your Cloud Project.

If you provide `token_file`, the plugin will write to the file your OAuth access token and refresh token. It can help avoid the need to re-authenticate every time the plugin runs.
If you run the plugin in an automated environment, you should probably authenticate with a service account.

### Service Account

To authenticate with a service account, you need to provide a JSON key file. You can create a service account key file in the Google Cloud Console.
You also need to enable the Admin SDK API for your Cloud Project.

Because you are accessing the Admin SDK via a service account, you need to impersonate a user with the necessary permissions to access the data you want to query.
In Google Workspace, your service account can only impersonate a user if the Client ID of the service account is granted domain-wide delegation.

You can follow this guide to grant your service account domain-wide delegation: [link](https://developers.google.com/cloud-search/docs/guides/delegation#delegate_domain-wide_authority_to_your_service_account)

When granting domain wide delegation, you need to provide a list of OAuth scopes. Here is the list you provide: (they are all read-only scopes)

```
https://www.googleapis.com/auth/admin.directory.customer.readonly,https://www.googleapis.com/auth/admin.directory.domain.readonly,https://www.googleapis.com/auth/admin.directory.group.member.readonly,https://www.googleapis.com/auth/admin.directory.group.readonly,https://www.googleapis.com/auth/admin.directory.orgunit.readonly,https://www.googleapis.com/auth/admin.directory.user.alias.readonly,https://www.googleapis.com/auth/admin.directory.user.readonly,https://www.googleapis.com/auth/admin.directory.userschema.readonly,https://www.googleapis.com/auth/admin.directory.resource.calendar.readonly,https://www.googleapis.com/auth/admin.directory.device.chromeos.readonly,https://www.googleapis.com/auth/admin.chrome.printers.readonly
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
