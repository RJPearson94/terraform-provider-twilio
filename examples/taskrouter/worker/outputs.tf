output "workspace" {
  description = "The Generated TaskRouter Workspace"
  value       = twilio_taskrouter_workspace.workspace
}

output "worker" {
  description = "The Generated Worker"
  value       = twilio_taskrouter_worker.worker
}
