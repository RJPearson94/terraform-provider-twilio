output "assistant" {
  description = "The Generated Autopilot Assistant"
  value       = twilio_autopilot_assistant.assistant
}

output "task" {
  description = "The Generated Autopilot Task"
  value       = twilio_autopilot_task.task
}

output "task_sample" {
  description = "The Generated Autopilot Task Sample"
  value       = twilio_autopilot_task_sample.task_sample
}
