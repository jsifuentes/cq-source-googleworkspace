The following source configuration file will sync to a PostgreSQL database. See
[the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more
information on how to configure the source and destination.

If `token_file` is set, after you successfully authenticate via OAuth, your
access token will be written to the `token_file`. The `token_file` will be used
on subsequent syncs.

How to find your Google Workspace Customer ID:
https://support.google.com/a/answer/10070793?hl=en

You can get your own OAuth credentials using
[this guide](https://developers.google.com/identity/protocols/oauth2#1.-obtain-oauth-2.0-credentials-from-the-dynamic_data.setvar.console_name-.).
You need to enable the Admin SDK API for your Cloud Project.

```yaml
kind: source
spec:
  name: "googleworkspace"
  path: "jsifuentes/googleworkspace"
  registry: "github"
  version: "v1.0.1"
  destinations:
    - "postgresql"
  spec:
    customer_id: your Google Workspace Customer ID
    oauth:
      # token_file: ./token.json
      client_id: your Google Cloud Project OAuth Client ID
      client_secret: your Google Cloud Project OAuth Client Secret
```
