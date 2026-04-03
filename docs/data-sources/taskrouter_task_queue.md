---
page_title: "twilio_taskrouter_task_queue Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_task_queue Data Source

Use this data source to access information about an existing TaskRouter task queue. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-queue) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_queue" {
  value = data.twilio_taskrouter_task_queue.task_queue
}
```

## Schema

### Required

- `sid` (String) The SID of the task queue to retrieve
- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this task queue
- `assignment_activity_name` (String) The name of the activity to assign workers when a task is assigned
- `assignment_activity_sid` (String) The SID of the activity to assign workers when a task is assigned
- `date_created` (String) The date and time the task queue was created, in RFC 3339 format
- `date_updated` (String) The date and time the task queue was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the task queue
- `id` (String) The ID of this resource.
- `max_reserved_workers` (Number) The maximum number of workers to reserve for a task in this queue
- `reservation_activity_name` (String) The name of the activity to assign workers when a task is reserved
- `reservation_activity_sid` (String) The SID of the activity to assign workers when a task is reserved
- `target_workers` (String) A worker selection expression for this task queue
- `task_order` (String) The order in which tasks are assigned to workers
- `url` (String) The absolute URL of the task queue resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
