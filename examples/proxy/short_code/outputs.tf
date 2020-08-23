output "service" {
  description = "The Generated Proxy Service"
  value       = twilio_proxy_service.service
}

output "short_code" {
  description = "The Short Code associated with the Proxy Service"
  value       = twilio_proxy_short_code.short_code
}
