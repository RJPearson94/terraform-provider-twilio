output "workspace" {
  description = "The Generated TaskRouter Workspace"
  value       = twilio_taskrouter_workspace.workspace
}

output "task_channel" {
  description = "The Generated TaskRouter Task Channel"
  value       = twilio_taskrouter_task_channel.task_channel
}
