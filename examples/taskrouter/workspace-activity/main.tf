resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_workspace_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Workspace Activity"
  available     = true
}
