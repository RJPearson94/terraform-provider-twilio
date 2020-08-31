output "service" {
  description = "The Generated Chat Service"
  value       = twilio_chat_service.service
}

output "user" {
  description = "The Generated Channel User"
  value       = twilio_chat_user.user
}
