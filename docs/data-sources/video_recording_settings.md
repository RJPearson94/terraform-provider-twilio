---
page_title: "twilio_video_recording_settings Data Source - twilio"
subcategory: "Video"
description: |-
  
---

# twilio_video_recording_settings Data Source

Use this data source to access information about the default recording settings. See the [encrypted recording docs](https://www.twilio.com/docs/video/api/encrypted-recordings) and [external S3 recording docs](https://www.twilio.com/docs/video/api/external-s3-recordings) for more information

!> This feature is only available as part of the [Twilio Enterprise Edition and Security Edition](https://www.twilio.com/editions)

## Example Usage

```hcl
data "twilio_video_recording_settings" "recording_settings" {}
```

## Schema

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this video recording settings
- `aws_credentials_sid` (String) The SID of the stored AWS credentials for external S3 recording storage
- `aws_s3_url` (String) The URL of the AWS S3 bucket where recordings are stored
- `aws_storage_enabled` (Boolean) Whether recordings are stored in an external AWS S3 bucket
- `encryption_enabled` (Boolean) Whether recordings are encrypted at rest
- `encryption_key_sid` (String) The SID of the stored encryption key used for at-rest encryption of recordings
- `friendly_name` (String) A human-readable label for the video recording settings
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the video recording settings resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
