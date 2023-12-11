# Table: googleworkspace_users

This table shows data for Googleworkspace Users.

Google Workspace Users

The primary key for this table is **id**.

## Relations

The following tables depend on googleworkspace_users:
  - [googleworkspace_user_aliases](googleworkspace_user_aliases.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|customer_id|`utf8`|
|first_name|`utf8`|
|last_name|`utf8`|
|agreed_to_terms|`bool`|
|aliases|`list<item: utf8, nullable>`|
|archived|`bool`|
|change_password_at_next_login|`bool`|
|creation_time|`utf8`|
|custom_schemas|`json`|
|deletion_time|`utf8`|
|etag|`utf8`|
|hash_function|`utf8`|
|id (PK)|`utf8`|
|include_in_global_address_list|`bool`|
|ip_whitelisted|`bool`|
|is_admin|`bool`|
|is_delegated_admin|`bool`|
|is_enforced_in2_sv|`bool`|
|is_enrolled_in2_sv|`bool`|
|is_mailbox_setup|`bool`|
|kind|`utf8`|
|last_login_time|`utf8`|
|name|`json`|
|non_editable_aliases|`list<item: utf8, nullable>`|
|org_unit_path|`utf8`|
|password|`utf8`|
|primary_email|`utf8`|
|recovery_email|`utf8`|
|recovery_phone|`utf8`|
|suspended|`bool`|
|suspension_reason|`utf8`|
|thumbnail_photo_etag|`utf8`|
|thumbnail_photo_url|`utf8`|