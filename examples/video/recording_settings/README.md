# Video Composition Settings

This example configures the default video recording setting to save the recordings in an external S3 bucket. This example provisions a KMS encrypted S3 bucket, an IAM user with only permissions to put objects in the bucket and to use the KMS key to manage the encryption of the object in your AWS account. The example also configured the necessary resources in Twilio to use the credentials to save the resources in your S3 bucket.

This feature is only available as part of the [Twilio Enterprise Edition and Security Edition](https://www.twilio.com/editions)

## Requirements

| Name      | Version   |
| --------- | --------- |
| terraform | >= 0.13   |
| aws       | >= 3.51.0 |
| twilio    | >= 0.14.0 |

## Providers

| Name   | Version   |
| ------ | --------- |
| aws    | >= 3.51.0 |
| twilio | >= 0.14.0 |

## Modules

No Modules.

## Resources

| Name                                                                                                                                         |
| -------------------------------------------------------------------------------------------------------------------------------------------- |
| [aws_iam_access_key](https://registry.terraform.io/providers/hashicorp/aws/3.51.0/docs/resources/iam_access_key)                             |
| [aws_iam_policy](https://registry.terraform.io/providers/hashicorp/aws/3.51.0/docs/resources/iam_policy)                                     |
| [aws_iam_user](https://registry.terraform.io/providers/hashicorp/aws/3.51.0/docs/resources/iam_user)                                         |
| [aws_iam_user_policy_attachment](https://registry.terraform.io/providers/hashicorp/aws/3.51.0/docs/resources/iam_user_policy_attachment)     |
| [aws_kms_key](https://registry.terraform.io/providers/hashicorp/aws/3.51.0/docs/resources/kms_key)                                           |
| [aws_s3_bucket](https://registry.terraform.io/providers/hashicorp/aws/3.51.0/docs/resources/s3_bucket)                                       |
| [twilio_credentials_aws](https://registry.terraform.io/providers/RJPearson94/twilio/0.14.0/docs/resources/credentials_aws)                   |
| [twilio_video_recording_settings](https://registry.terraform.io/providers/RJPearson94/twilio/0.14.0/docs/resources/video_recording_settings) |

## Inputs

No input.

## Outputs

No output.
