output "service" {
  description = "The Generated Serverless Service"
  value       = twilio_serverless_service.service
}

output "staging_environment" {
  description = "The Generated Staging Serverless Environment"
  value       = twilio_serverless_environment.staging
}

output "prod_environment" {
  description = "The Generated Prod Serverless Environment"
  value       = twilio_serverless_environment.prod
}
