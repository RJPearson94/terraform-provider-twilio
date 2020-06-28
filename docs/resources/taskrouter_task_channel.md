---
page_title: "Twilio TaskRouter Task Channel"
subcategory: "TaskRouter"
---

# twilio_taskrouter_task_channel Resource

Manages a TaskRouter Task Channel

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_channel" "task_channel" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Task Channel"
  unique_name   = "Unique Task Channel"
}
```

## Argument Reference

The following arguments are supported:

* `friendly_name` - (Mandatory) The name of the task channel
* `unique_name` - (Mandatory) The unique name of the task channel. Changing this forces a new resource to be created
* `channel_optimized_routing` - (Optional) Whether the task channel should prioritise idle workers. Default is false

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the task channel (Same as the SID)
* `sid` - The SID of the task channel (Same as the ID)
* `account_sid` - The Account SID of the task channel is deployed into
* `workspaceSid` - The Workspace SID to create the task channel under
* `friendly_name` - The name of the task channel
* `unique_name` - The unique name of the task channel
* `channel_optimized_routing` - Whether the task channel should prioritise idle workers
* `date_created` - The date that the task channel was created
* `date_updated` - The date that the task channel was updated
* `url` - The url of the task channel
