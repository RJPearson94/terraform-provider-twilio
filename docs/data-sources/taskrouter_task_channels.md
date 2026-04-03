---
page_title: "twilio_taskrouter_task_channels Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_task_channels Data Source

Use this data source to access information about the task channels associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-channel) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_task_channels" "task_channels" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_channels" {
  value = data.twilio_taskrouter_task_channels.task_channels
}
```

## Schema

### Required

- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the task channels
- `id` (String) The ID of this resource.
- `task_channels` (List of Object) A list of task channels in the workspace (see [below for nested schema](#nestedatt--task_channels))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--task_channels"></a>
### Nested Schema for `task_channels`

Read-Only:

- `channel_optimized_routing` (Boolean)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
- `unique_name` (String)
- `url` (String)
