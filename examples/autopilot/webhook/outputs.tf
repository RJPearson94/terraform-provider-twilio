output "assistant" {
  description = "The Generated Autopilot Assistant"
  value       = twilio_autopilot_assistant.assistant
}

output "webhook" {
  description = "The Generated Autopilot Webhook"
  value       = twilio_autopilot_webhook.webhook
}