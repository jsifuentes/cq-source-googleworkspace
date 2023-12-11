# Table: googleworkspace_user_aliases

This table shows data for Googleworkspace User Aliases.

Google Workspace User Aliases

The composite primary key for this table is (**user_id**, **alias**).

## Relations

This table depends on [googleworkspace_users](googleworkspace_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|user_id (PK)|`utf8`|
|user_primary_email|`utf8`|
|alias (PK)|`utf8`|