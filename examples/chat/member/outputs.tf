output "service" {
  description = "The Generated Chat Service"
  value       = twilio_chat_service.service
}

output "channel" {
  description = "The Generated Chat Channel"
  value       = twilio_chat_channel.channel
}

output "user" {
  description = "The Generated Chat User"
  value       = twilio_chat_user.user
}

output "member" {
  description = "The Generated Chat Channel Member"
  value       = twilio_chat_channel_member.member
}
