---
page_title: "Twilio TaskRouter Resource"
subcategory: "TaskRouter"
---

# twilio_taskrouter_worker Resource

Manages a TaskRouter worker

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Worker"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The name of the worker
- `attributes` - (Optional) JSON string of worker attributes
- `activity_sid` - (Optional) Activity SID to be assigned to the worker

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the worker (Same as the SID)
- `sid` - The SID of the worker (Same as the ID)
- `account_sid` - The Account SID of the worker is deployed into
- `workspaceSid` - The Workspace SID to create the worker under
- `friendly_name` - The name of the worker
- `attributes` - JSON string of worker attributes
- `activity_sid` - Activity SID to be assigned to the worker
- `activity_sid_default` - Workspace default activity SID
- `activity_name` - Friendly name of activity
- `available` - Is the worker available to receive tasks
- `date_created` - The date in RFC3339 format that the worker was created
- `date_updated` - The date in RFC3339 format that the worker was updated
- `date_status_changed` - The date in RFC3339 format that the worker status was changed
- `url` - The url of the worker
