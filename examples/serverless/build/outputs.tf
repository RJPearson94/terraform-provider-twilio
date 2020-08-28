output "service" {
  description = "The Generated Serverless Service"
  value       = twilio_serverless_service.service
}

output "function" {
  description = "The Generated Serverless Function"
  value       = twilio_serverless_function.function
}

output "build" {
  description = "The Generated Serverless Build"
  value       = twilio_serverless_build.build
}
