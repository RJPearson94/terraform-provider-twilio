# twilio_taskrouter_workflow

Manages a TaskRouter workflow

## Example Usage

```hcl
resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Worker"
}

resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Queue"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Workflow"
  configuration = <<EOF
{
  "task_routing": {
    "filters": [],
    "default_filter": {
      "queue": "${twilio_taskrouter_task_queue.task_queue.sid}"
    }
  }
}
EOF
}
```

## Argument Reference

The following arguments are supported:

* `friendly_name` - (Mandatory) The name of the workflow
* `configuration` - (Mandatory) JSON string of workflow configuration
* `assignment_callback_url` - (Optional) Assignment Callback URL
* `fallback_assignment_callback_url` - (Optional) Fallback Assignment Callback URL
* `task_reservation_timeout` - (Optional) Maximum time the task can be unassigned for before it times out. Default is 120s

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the workflow (Same as the SID)
* `sid` - The SID of the workflow (Same as the ID)
* `account_sid` - The Account SID of the worker is deployed into
* `workspaceSid` - The Workspace SID to create the worker under
* `friendly_name` - The name of the worker
* `configuration` - JSON string of workflow configuration
* `assignment_callback_url` - Assignment Callback URL
* `fallback_assignment_callback_url` - Fallback Assignment Callback URL
* `task_reservation_timeout` - Maximum time the task can be unassigned for before it times out
* `document_content_type` - The MIME type of the document
* `date_created` - The date that the worker was created
* `date_updated` - The date that the worker was updated
* `url` - The url of the worker
