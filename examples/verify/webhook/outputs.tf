output "service" {
  description = "The Generated Verify service"
  value       = twilio_verify_service.service
}

output "webhook" {
  description = "The Generated Verify Webhook"
  value       = twilio_verify_webhook.webhook
}
