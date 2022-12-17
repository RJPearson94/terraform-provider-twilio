output "service_sid" {
  description = "The SID of the Verify Service"
  value       = twilio_verify_service.service.sid
}