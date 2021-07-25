---
page_title: "Twilio Video Composition Settings"
subcategory: "Video"
---

# twilio_video_composition_settings Data Source

Use this data source to access information about the default composition settings. See the [encrypted composition docs](https://www.twilio.com/docs/video/api/encrypted-compositions) and [external S3 composition docs](https://www.twilio.com/docs/video/api/external-s3-compositions) for more information

!> This feature is only available as part of the [Twilio Enterprise Edition and Security Edition](https://www.twilio.com/editions)

## Example Usage

```hcl
data "twilio_video_composition_settings" "composition_settings" {}
```

## Argument Reference

The following arguments are supported:

N/A - This data source has no arguments

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the composition settings (Same as the `account_sid`)
- `account_sid` - The account SID the composition settings is associated with
- `aws_credentials_sid` - The SID of the AWS credentials supplied to Twilio which are used to store compositions in your S3 bucket
- `aws_s3_url` - The URL of the S3 bucket where compositions are stored
- `aws_storage_enabled` - (Optional) Whether compositions are stored in your S3 bucket
- `encryption_enabled` - Whether encrypted compositions is enabled
- `encryption_key_sid` - The SID of the credential supplied to Twilio which is used to encrypt the compositions
- `url` - The URL of the composition settings

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the default composition settings
