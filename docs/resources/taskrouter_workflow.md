---
page_title: "twilio_taskrouter_workflow Resource - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_workflow Resource

Manages a TaskRouter workflow. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workflow) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "twilio-test"
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
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
}
```

## Schema

### Required

- `configuration` (String) A JSON string of the workflow configuration
- `friendly_name` (String) A human-readable label for the workflow
- `workspace_sid` (String) The SID of the TaskRouter workspace. Changing this forces a new resource

### Optional

- `assignment_callback_url` (String) The URL to call when a task is assigned to a worker
- `fallback_assignment_callback_url` (String) The URL to call when a task assignment event is not handled by the primary callback
- `task_reservation_timeout` (Number) The timeout in seconds for a task reservation (1-86400). Defaults to `120`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this workflow
- `date_created` (String) The date and time the workflow was created, in RFC 3339 format
- `date_updated` (String) The date and time the workflow was last updated, in RFC 3339 format
- `document_content_type` (String) The MIME type of the workflow document
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this workflow by Twilio
- `url` (String) The absolute URL of the workflow resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A workflow can be imported using the `/Workspaces/{workspaceSid}/Workflows/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_workflow.workflow /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
