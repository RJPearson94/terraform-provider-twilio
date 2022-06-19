---
page_title: "Twilio Video Composition Settings"
subcategory: "Video"
---

# twilio_video_composition_settings Resource

Manages the default Programmable Video composition settings. See the [encrypted composition docs](https://www.twilio.com/docs/video/api/encrypted-compositions) and [external S3 composition docs](https://www.twilio.com/docs/video/api/external-s3-compositions) for more information

!> This resource modifies the default video composition settings for the account. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

!> This feature is only available as part of the [Twilio Enterprise Edition and Security Edition](https://www.twilio.com/editions)

## Example Usage

### With Encryption

```hcl
resource "twilio_credentials_public_key" "public_key" {
  friendly_name = "Test Public Key Resource"
  public_key    = "-----BEGIN PUBLIC KEY-----....-----END PUBLIC KEY-----"
}

resource "twilio_video_composition_settings" "composition_settings" {
  friendly_name      = "Composition Settings"
  encryption_enabled = true
  encryption_key_sid = twilio_credentials_public_key.public_key.sid
}
```

### With External S3 Bucket - Basic

```hcl
resource "twilio_credentials_aws" "aws" {
  friendly_name         = "Test AWS Credential Resource"
  aws_access_key_id     = "aws_access_key_id"
  aws_secret_access_key = "aws_secret_access_key"
}

resource "twilio_video_composition_settings" "composition_settings" {
  friendly_name       = "Composition Settings"
  aws_credentials_sid = twilio_credentials_aws.aws.sid
  aws_storage_enabled = true
  aws_s3_url          = "https://test-bucket.s3.amazonaws.com/compositions/"
}
```

### With External S3 Bucket - With Twilio & AWS Providers

```hcl
// Create AWS IAM User in your AWS account
resource "aws_iam_user" "iam_user" {
  name = "test-iam-user"
}

resource "aws_iam_access_key" "access_key" {
  user = aws_iam_user.iam_user.name
}

// Supply AWS IAM User credentials to your Twilio account
resource "twilio_credentials_aws" "aws" {
  friendly_name         = "Test AWS Credential Resource"
  aws_access_key_id     = aws_iam_access_key.access_key.id
  aws_secret_access_key = aws_iam_access_key.access_key.secret
}

// Create encrypted private bucket in your AWS account
resource "aws_kms_key" "kms_key" {
  description             = "This key is used to encrypt bucket objects"
  deletion_window_in_days = 7
}

resource "aws_s3_bucket" "test_bucket" {
  bucket = "my-test-bucket"
  acl    = "private"

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = aws_kms_key.kms_key.arn
        sse_algorithm     = "aws:kms"
      }
    }
  }
}

// Grant the IAM user permissions to put objects into the encrypted S3 bucket
// The policy is adapted from the Twilio documentation https://www.twilio.com/docs/video/tutorials/storing-aws-s3
resource "aws_iam_policy" "policy" {
  name        = "test-policy"
  description = "A test policy"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "UploadUserDenyEverything",
        "Effect" : "Deny",
        "NotAction" : "*",
        "Resource" : "*"
      },
      {
        "Sid" : "UploadUserAllowPutObject",
        "Effect" : "Allow",
        "Action" : [
          "s3:PutObject"
        ],
        "Resource" : [
          "${aws_s3_bucket.test_bucket.arn}/compositions/*"
        ]
      },
      {
        "Sid" : "AccessToKmsForEncryption",
        "Effect" : "Allow",
        "Action" : [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey"
        ],
        "Resource" : [
          aws_kms_key.kms_key.arn
        ]
      }
    ]
  })
}

resource "aws_iam_user_policy_attachment" "policy_attachment" {
  user       = aws_iam_user.iam_user.name
  policy_arn = aws_iam_policy.policy.arn
}

// Update the default composition settings to enable storage of data in your S3 bucket
resource "twilio_video_composition_settings" "composition_settings" {
  friendly_name       = "Composition Settings"
  aws_credentials_sid = twilio_credentials_aws.aws.sid
  aws_storage_enabled = true
  aws_s3_url          = "https://${aws_s3_bucket.test_bucket.bucket_domain_name}/compositions/"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The friendly name of the composition settings
- `aws_credentials_sid` - (Optional) The SID of the AWS credentials supplied to Twilio to use to store compositions in your S3 bucket
- `aws_s3_url` - (Optional) The URL of the S3 bucket to store compositions in
- `aws_storage_enabled` - (Optional) Whether to store compositions in your S3 bucket. The default value is `false`
- `encryption_enabled` - (Optional) Whether to encrypt the compositions. The default value is `false`
- `encryption_key_sid` - (Optional) The SID of the credential supplied to Twilio to use to encrypt the compositions

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

- `create` - (Defaults to 10 minutes) Used when updating the composition settings
- `update` - (Defaults to 10 minutes) Used when updating the composition settings
- `read` - (Defaults to 5 minutes) Used when retrieving the composition settings
