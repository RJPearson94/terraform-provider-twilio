---
page_title: "twilio_taskrouter_workspace Resource - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_workspace Resource

Manages a TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workspace) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name      = "Test Workspace"
  multi_task_enabled = true
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the workspace

### Optional

- `event_callback_url` (String) The URL to call when an event is fired in the workspace
- `event_filters` (List of String) A list of event types to subscribe to. Valid values are `task.created`, `task.completed`, `task.canceled`, `task.deleted`, `task.updated`, `task.wrapup`, `task-queue.entered`, `task-queue.moved`, `task-queue.timeout`, `reservation.created`, `reservation.accepted`, `reservation.rejected`, `reservation.timeout`, `reservation.canceled`, `reservation.rescinded`, `reservation.wrapup`, `reservation.completed`, `reservation.failed`, `workflow.entered`, `workflow.timeout`, `workflow.target-matched`, `worker.activity.update`, `worker.attributes.update`, `worker.capacity.update`, `worker.channel.availability.update`
- `multi_task_enabled` (Boolean) Whether multi-tasking is enabled for the workspace. Defaults to `true`
- `prioritize_queue_order` (String) The order in which task queues are prioritized. Valid values are `LIFO` or `FIFO`. Defaults to `FIFO`
- `template` (String) The template to use when creating the workspace. Valid values are `NONE` or `FIFO`. Defaults to `NONE`. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this workspace
- `date_created` (String) The date and time the workspace was created, in RFC 3339 format
- `date_updated` (String) The date and time the workspace was last updated, in RFC 3339 format
- `default_activity_name` (String) The name of the default activity for the workspace
- `default_activity_sid` (String) The SID of the default activity for the workspace
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this workspace by Twilio
- `timeout_activity_name` (String) The name of the timeout activity for the workspace
- `timeout_activity_sid` (String) The SID of the timeout activity for the workspace
- `url` (String) The absolute URL of the workspace resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A workspace can be imported using the `/Workspaces/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_workspace.workspace /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> `template` cannot be imported
