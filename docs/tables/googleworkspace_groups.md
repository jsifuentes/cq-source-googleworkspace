# Table: googleworkspace_groups

This table shows data for Google Workspace Groups.

Google Workspace Groups

The primary key for this table is **id**.

## Relations

The following tables depend on googleworkspace_groups:
  - [googleworkspace_group_aliases](googleworkspace_group_aliases.md)
  - [googleworkspace_group_members](googleworkspace_group_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id|String|
|admin_created|Bool|
|aliases|StringArray|
|description|String|
|direct_members_count|Int|
|email|String|
|etag|String|
|id (PK)|String|
|kind|String|
|name|String|
|non_editable_aliases|StringArray|