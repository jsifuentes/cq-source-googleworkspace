# Table: googleworkspace_users

This table shows data for Google Workspace Users.

Google Workspace Users

The primary key for this table is **id**.

## Relations

The following tables depend on googleworkspace_users:
  - [googleworkspace_user_aliases](googleworkspace_user_aliases.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|customer_id|String|
|first_name|String|
|last_name|String|
|agreed_to_terms|Bool|
|aliases|StringArray|
|archived|Bool|
|change_password_at_next_login|Bool|
|creation_time|String|
|custom_schemas|JSON|
|deletion_time|String|
|etag|String|
|hash_function|String|
|id (PK)|String|
|include_in_global_address_list|Bool|
|ip_whitelisted|Bool|
|is_admin|Bool|
|is_delegated_admin|Bool|
|is_enforced_in2_sv|Bool|
|is_enrolled_in2_sv|Bool|
|is_mailbox_setup|Bool|
|kind|String|
|last_login_time|String|
|name|JSON|
|non_editable_aliases|StringArray|
|org_unit_path|String|
|password|String|
|primary_email|String|
|recovery_email|String|
|recovery_phone|String|
|suspended|Bool|
|suspension_reason|String|
|thumbnail_photo_etag|String|
|thumbnail_photo_url|String|