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
          "${aws_s3_bucket.test_bucket.arn}/recordings/*"
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

// Update the default recording settings to enable storage of data in your S3 bucket 
resource "twilio_video_recording_settings" "recording_settings" {
  friendly_name       = "Composition Settings"
  aws_credentials_sid = twilio_credentials_aws.aws.sid
  aws_storage_enabled = true
  aws_s3_url          = "https://${aws_s3_bucket.test_bucket.bucket_domain_name}/recordings/"
}