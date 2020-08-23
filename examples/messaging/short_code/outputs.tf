output "service" {
  description = "The Generated Messaging Service"
  value       = twilio_messaging_service.service
}

output "short_code" {
  description = "The Short Code associated with the Messaging Service"
  value       = twilio_messaging_short_code.short_code
}
