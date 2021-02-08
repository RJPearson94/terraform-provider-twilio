output "service" {
  description = "The Generated Conversations Service"
  value       = twilio_conversations_service.service
}

output "role" {
  description = "The Generated Conversations Role"
  value       = twilio_conversations_role.role
}

