output "service" {
  description = "The Generated Chat Service"
  value       = twilio_chat_service.service
}

output "channel" {
  description = "The Generated Chat Channel"
  value       = twilio_chat_channel.channel
}

output "studio_webhook" {
  description = "The Generated Chat Channel Studio Webhook"
  value       = twilio_chat_channel_studio_webhook.studio_webhook
}
