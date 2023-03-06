# Table: gcp_storage_bucket_policies

https://cloud.google.com/iam/docs/reference/rest/v1/Policy

The primary key for this table is **_cq_id**.

## Relations

This table depends on [gcp_storage_buckets](gcp_storage_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|bucket_name|String|
|bindings|JSON|