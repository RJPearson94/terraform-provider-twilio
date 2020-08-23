output "service" {
  description = "The Generated Proxy Service"
  value       = twilio_proxy_service.service
}

output "phone_number" {
  description = "The Phone Number associated with the Proxy Service"
  value       = twilio_proxy_phone_number.phone_number
}
