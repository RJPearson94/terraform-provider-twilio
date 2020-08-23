output "service" {
  description = "The Generated Messaging Service"
  value       = twilio_messaging_service.service
}

output "alpha_sender" {
  description = "The Alpha Sender associated with the Messaging Service"
  value       = twilio_messaging_alpha_sender.alpha_sender
}
