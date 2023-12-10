# Table: googleworkspace_group_aliases

This table shows data for Google Workspace Group Aliases.

Google Workspace Group Aliases

The composite primary key for this table is (**group_id**, **alias**).

## Relations

This table depends on [googleworkspace_groups](googleworkspace_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|group_id (PK)|String|
|group_primary_email|String|
|alias (PK)|String|