---
page_title: "twilio_taskrouter_task_queue Resource - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_task_queue Resource

Manages a TaskRouter task queue. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-queue) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Task Queue"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the task queue
- `workspace_sid` (String) The SID of the TaskRouter workspace. Changing this forces a new resource

### Optional

- `assignment_activity_sid` (String) The SID of the activity to assign workers when a task is assigned
- `max_reserved_workers` (Number) The maximum number of workers to reserve for a task in this queue (1-50). Defaults to `1`
- `reservation_activity_sid` (String) The SID of the activity to assign workers when a task is reserved
- `target_workers` (String) A worker selection expression for this task queue. Defaults to `1==1`
- `task_order` (String) The order in which tasks are assigned to workers. Valid values are `LIFO` or `FIFO`. Defaults to `FIFO`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this task queue
- `assignment_activity_name` (String) The name of the activity to assign workers when a task is assigned
- `date_created` (String) The date and time the task queue was created, in RFC 3339 format
- `date_updated` (String) The date and time the task queue was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `reservation_activity_name` (String) The name of the activity to assign workers when a task is reserved
- `sid` (String) The unique SID assigned to this task queue by Twilio
- `url` (String) The absolute URL of the task queue resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A task queue can be imported using the `/Workspaces/{workspaceSid}/TaskQueues/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_task_queue.task_queue /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
