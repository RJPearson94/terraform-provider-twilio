output "workspace" {
  description = "The Generated TaskRouter Workspace"
  value       = twilio_taskrouter_workspace.workspace
}

output "worker" {
  description = "The generated Worker"
  value       = twilio_taskrouter_worker.worker
}
