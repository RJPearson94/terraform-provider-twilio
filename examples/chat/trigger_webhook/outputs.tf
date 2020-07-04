output "service" {
  description = "The Generated Chat Service"
  value       = twilio_chat_service.service
}

output "channel" {
  description = "The Generated Chat Channel"
  value       = twilio_chat_channel.channel
}

output "trigger_webhook" {
  description = "The Generated Chat Channel Trigger Webhook"
  value       = twilio_chat_channel_trigger_webhook.trigger_webhook
}
