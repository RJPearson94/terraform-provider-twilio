output "aws_credential" {
  description = "The Generated AWS Credential Resource"
  value       = twilio_credentials_aws.aws
  sensitive   = true
}
