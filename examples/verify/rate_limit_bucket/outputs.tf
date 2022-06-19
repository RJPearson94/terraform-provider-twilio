output "service" {
  description = "The Generated Verify service"
  value       = twilio_verify_service.service
}

output "rate_limit" {
  description = "The Generated Verify Service Rate Limit"
  value       = twilio_verify_service_rate_limit.rate_limit
}

output "rate_limit_bucket" {
  description = "The Generated Verify Service Rate Limit Bucket"
  value       = twilio_verify_service_rate_limit_bucket.rate_limit_bucket
}