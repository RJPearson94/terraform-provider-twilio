output "workspace" {
  description = "The Generated TaskRouter Workspace"
  value       = twilio_taskrouter_workspace.workspace
}

output "task_queue" {
  description = "The Generated TaskRouter Task Queue"
  value       = twilio_taskrouter_task_queue.task_queue
}
