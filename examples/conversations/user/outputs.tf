output "service" {
  description = "The Generated Conversations Service"
  value       = twilio_conversations_service.service
}

output "user" {
  description = "The Generated Conversations User"
  value       = twilio_conversations_user.user
}

