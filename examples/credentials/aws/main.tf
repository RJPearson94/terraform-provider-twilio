resource "aws_iam_user" "iam_user" {
  name = "test-iam-user"
}

resource "aws_iam_access_key" "access_key" {
  user = aws_iam_user.iam_user.name
}

resource "twilio_credentials_aws" "aws" {
  friendly_name         = "Test AWS Credential Resource"
  aws_access_key_id     = aws_iam_access_key.access_key.id
  aws_secret_access_key = aws_iam_access_key.access_key.secret
}
