---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_instance_storage_config Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::InstanceStorageConfig
---

# awscc_connect_instance_storage_config (Data Source)

Data Source schema for AWS::Connect::InstanceStorageConfig



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `association_id` (String) An associationID is automatically generated when a storage config is associated with an instance
- `instance_arn` (String) Connect Instance ID with which the storage config will be associated
- `kinesis_firehose_config` (Attributes) (see [below for nested schema](#nestedatt--kinesis_firehose_config))
- `kinesis_stream_config` (Attributes) (see [below for nested schema](#nestedatt--kinesis_stream_config))
- `kinesis_video_stream_config` (Attributes) (see [below for nested schema](#nestedatt--kinesis_video_stream_config))
- `resource_type` (String) Specifies the type of storage resource available for the instance
- `s3_config` (Attributes) (see [below for nested schema](#nestedatt--s3_config))
- `storage_type` (String) Specifies the storage type to be associated with the instance

<a id="nestedatt--kinesis_firehose_config"></a>
### Nested Schema for `kinesis_firehose_config`

Read-Only:

- `firehose_arn` (String) An ARN is a unique AWS resource identifier.


<a id="nestedatt--kinesis_stream_config"></a>
### Nested Schema for `kinesis_stream_config`

Read-Only:

- `stream_arn` (String) An ARN is a unique AWS resource identifier.


<a id="nestedatt--kinesis_video_stream_config"></a>
### Nested Schema for `kinesis_video_stream_config`

Read-Only:

- `encryption_config` (Attributes) (see [below for nested schema](#nestedatt--kinesis_video_stream_config--encryption_config))
- `prefix` (String) Prefixes are used to infer logical hierarchy
- `retention_period_hours` (Number) Number of hours

<a id="nestedatt--kinesis_video_stream_config--encryption_config"></a>
### Nested Schema for `kinesis_video_stream_config.encryption_config`

Read-Only:

- `encryption_type` (String) Specifies default encryption using AWS KMS-Managed Keys
- `key_id` (String) Specifies the encryption key id



<a id="nestedatt--s3_config"></a>
### Nested Schema for `s3_config`

Read-Only:

- `bucket_name` (String) A name for the S3 Bucket
- `bucket_prefix` (String) Prefixes are used to infer logical hierarchy
- `encryption_config` (Attributes) (see [below for nested schema](#nestedatt--s3_config--encryption_config))

<a id="nestedatt--s3_config--encryption_config"></a>
### Nested Schema for `s3_config.encryption_config`

Read-Only:

- `encryption_type` (String) Specifies default encryption using AWS KMS-Managed Keys
- `key_id` (String) Specifies the encryption key id


