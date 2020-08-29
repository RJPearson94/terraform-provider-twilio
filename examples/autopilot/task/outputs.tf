output "assistant" {
  description = "The Generated Autopilot Assistant"
  value       = twilio_autopilot_assistant.assistant
}

output "task" {
  description = "The Generated Autopilot Task"
  value       = twilio_autopilot_task.task
}

output "tasks" {
  description = "The Generated Autopilot Tasks"
  value       = data.twilio_autopilot_tasks.tasks
}
