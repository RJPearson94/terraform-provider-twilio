output "workspace" {
  description = "The Generated TaskRouter Workspace"
  value       = twilio_taskrouter_workspace.workspace
}

output "task_queue" {
  description = "The Generated task queue"
  value       = twilio_taskrouter_task_queue.task_queue
}

output "workflow" {
  description = "The Generated workflow"
  value       = twilio_taskrouter_workflow.workflow
}
