output "assistant" {
  description = "The Generated Autopilot Assistant"
  value       = twilio_autopilot_assistant.assistant
}

output "field_type" {
  description = "The Generated Autopilot Field Type"
  value       = twilio_autopilot_field_type.field_type
}