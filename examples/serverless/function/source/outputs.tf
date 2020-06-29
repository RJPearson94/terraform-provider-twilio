output "service" {
  description = "The Generated Serverless Service"
  value       = twilio_serverless_service.service
}

output "function" {
  description = "The Generated Serverless Function"
  value       = twilio_serverless_function.function
}

output "function_version" {
  description = "The Generated Serverless Function Version"
  value       = twilio_serverless_function_version.function_version
}
