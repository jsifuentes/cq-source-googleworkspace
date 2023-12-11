# Table: googleworkspace_groups

This table shows data for Googleworkspace Groups.

Google Workspace Groups

The primary key for this table is **id**.

## Relations

The following tables depend on googleworkspace_groups:
  - [googleworkspace_group_aliases](googleworkspace_group_aliases.md)
  - [googleworkspace_group_members](googleworkspace_group_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|customer_id|`utf8`|
|admin_created|`bool`|
|aliases|`list<item: utf8, nullable>`|
|description|`utf8`|
|direct_members_count|`int64`|
|email|`utf8`|
|etag|`utf8`|
|id (PK)|`utf8`|
|kind|`utf8`|
|name|`utf8`|
|non_editable_aliases|`list<item: utf8, nullable>`|