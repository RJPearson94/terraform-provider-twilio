output "service" {
  description = "The Generated Verify service"
  value       = twilio_verify_service.service
}

output "messaging_configuration" {
  description = "The Generated Verify Messaging Configuration"
  value       = twilio_verify_messaging_configuration.messaging_configuration
}
