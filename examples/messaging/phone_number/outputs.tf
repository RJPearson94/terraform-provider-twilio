output "service" {
  description = "The Generated Messaging Service"
  value       = twilio_messaging_service.service
}

output "phone_number" {
  description = "The Phone Number associated with the Messaging Service"
  value       = twilio_messaging_phone_number.phone_number
}
