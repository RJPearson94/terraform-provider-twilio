---
page_title: "twilio_video_composition_settings Data Source - twilio"
subcategory: "Video"
description: |-
  
---

# twilio_video_composition_settings Data Source

Use this data source to access information about the default composition settings. See the [encrypted composition docs](https://www.twilio.com/docs/video/api/encrypted-compositions) and [external S3 composition docs](https://www.twilio.com/docs/video/api/external-s3-compositions) for more information

!> This feature is only available as part of the [Twilio Enterprise Edition and Security Edition](https://www.twilio.com/editions)

## Example Usage

```hcl
data "twilio_video_composition_settings" "composition_settings" {}
```

## Schema

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this video composition settings
- `aws_credentials_sid` (String) The SID of the stored AWS credentials for external S3 composition storage
- `aws_s3_url` (String) The URL of the AWS S3 bucket where compositions are stored
- `aws_storage_enabled` (Boolean) Whether compositions are stored in an external AWS S3 bucket
- `encryption_enabled` (Boolean) Whether compositions are encrypted at rest
- `encryption_key_sid` (String) The SID of the stored encryption key used for at-rest encryption of compositions
- `friendly_name` (String) A human-readable label for the video composition settings
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the video composition settings resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
