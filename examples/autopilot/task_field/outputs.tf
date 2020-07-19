output "assistant" {
  description = "The Generated Autopilot Assistant"
  value       = twilio_autopilot_assistant.assistant
}

output "task" {
  description = "The Generated Autopilot Task"
  value       = twilio_autopilot_task.task
}

output "task_field" {
  description = "The Generated Autopilot Task Field"
  value       = twilio_autopilot_task_field.task_field
}