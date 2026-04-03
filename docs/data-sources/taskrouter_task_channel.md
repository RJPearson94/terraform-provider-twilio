---
page_title: "twilio_taskrouter_task_channel Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_task_channel Data Source

Use this data source to access information about an existing TaskRouter task channel. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-channel) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

### SID

```hcl
data "twilio_taskrouter_task_channel" "task_channel" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_channel" {
  value = data.twilio_taskrouter_task_channel.task_channel
}
```

### Unique Name

```hcl
data "twilio_taskrouter_task_channel" "task_channel" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  unique_name   = "UniqueName"
}

output "task_channel" {
  value = data.twilio_taskrouter_task_channel.task_channel
}
```

## Schema

### Required

- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `sid` (String) The SID of the task channel to retrieve. Exactly one of `sid` or `unique_name` must be specified
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `unique_name` (String) The unique name of the task channel to retrieve. Exactly one of `sid` or `unique_name` must be specified

### Read-Only

- `account_sid` (String) The SID of the account that owns this task channel
- `channel_optimized_routing` (Boolean) Whether the task channel prioritizes optimized routing
- `date_created` (String) The date and time the task channel was created, in RFC 3339 format
- `date_updated` (String) The date and time the task channel was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the task channel
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the task channel resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
