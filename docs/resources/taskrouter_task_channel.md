---
page_title: "twilio_taskrouter_task_channel Resource - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_task_channel Resource

Manages a TaskRouter task channel. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-channel) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

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

## Schema

### Required

- `friendly_name` (String) A human-readable label for the task channel
- `unique_name` (String) A unique identifier for the task channel. Changing this forces a new resource
- `workspace_sid` (String) The SID of the TaskRouter workspace. Changing this forces a new resource

### Optional

- `channel_optimized_routing` (Boolean) Whether the task channel should prioritize optimized routing. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this task channel
- `date_created` (String) The date and time the task channel was created, in RFC 3339 format
- `date_updated` (String) The date and time the task channel was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this task channel by Twilio
- `url` (String) The absolute URL of the task channel resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A task channel can be imported using the `/Workspaces/{workspaceSid}/TaskChannels/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_task_channel.task_channel /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
