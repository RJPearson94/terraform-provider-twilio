output "workspace" {
  description = "The Generated TaskRouter Workspace"
  value       = twilio_taskrouter_workspace.workspace
}

output "workspace_activity" {
  description = "The Generated TaskRouter Workspace Activity"
  value       = twilio_taskrouter_workspace_activity.activity
}
