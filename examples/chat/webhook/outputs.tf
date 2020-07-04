output "service" {
  description = "The Generated Chat Service"
  value       = twilio_chat_service.service
}

output "channel" {
  description = "The Generated Chat Channel"
  value       = twilio_chat_channel.channel
}

output "webhook" {
  description = "The Generated Chat Channel Webhook"
  value       = twilio_chat_channel_webhook.webhook
}
