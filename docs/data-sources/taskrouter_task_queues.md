---
page_title: "twilio_taskrouter_task_queues Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_task_queues Data Source

Use this data source to access information about the task queues associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-queue) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_task_queues" "task_queues" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_queues" {
  value = data.twilio_taskrouter_task_queues.task_queues
}
```

## Schema

### Required

- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the task queues
- `id` (String) The ID of this resource.
- `task_queues` (List of Object) A list of task queues in the workspace (see [below for nested schema](#nestedatt--task_queues))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--task_queues"></a>
### Nested Schema for `task_queues`

Read-Only:

- `assignment_activity_name` (String)
- `assignment_activity_sid` (String)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `max_reserved_workers` (Number)
- `reservation_activity_name` (String)
- `reservation_activity_sid` (String)
- `sid` (String)
- `target_workers` (String)
- `task_order` (String)
- `url` (String)
