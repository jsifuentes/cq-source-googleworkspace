# Table: googleworkspace_user_aliases

This table shows data for Google Workspace User Aliases.

Google Workspace User Aliases

The composite primary key for this table is (**user_id**, **alias**).

## Relations

This table depends on [googleworkspace_users](googleworkspace_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|user_id (PK)|String|
|user_primary_email|String|
|alias (PK)|String|