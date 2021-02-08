output "service" {
  description = "The Generated Conversations Service"
  value       = twilio_conversations_service.service
}

output "configuration" {
  description = "The Generated Conversations Service Configuration"
  value       = twilio_conversations_service_configuration.configuration
}

output "notifications" {
  description = "The Generated Conversations Service Notifications"
  value       = twilio_conversations_service_notification.service_notification
}
