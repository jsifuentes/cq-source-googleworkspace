# Table: googleworkspace_group_aliases

This table shows data for Googleworkspace Group Aliases.

Google Workspace Group Aliases

The composite primary key for this table is (**group_id**, **alias**).

## Relations

This table depends on [googleworkspace_groups](googleworkspace_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|group_id (PK)|`utf8`|
|group_primary_email|`utf8`|
|alias (PK)|`utf8`|