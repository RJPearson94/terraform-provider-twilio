output "service" {
  description = "The Generated Chat Service"
  value       = twilio_chat_service.service
}

output "role" {
  description = "The Generated Channel role"
  value       = twilio_chat_role.role
}
