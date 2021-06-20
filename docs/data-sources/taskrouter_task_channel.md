---
page_title: "Twilio TaskRouter Task Channel"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the task channel is associated with
- `sid` - (Optional) The SID of the task channel
- `unique_name` - (Optional) The unique name of the task channel

~> Either `sid` or `unique_name` must be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the task channel (Same as the `sid`)
- `sid` - The SID of the task channel (Same as the `id`)
- `account_sid` - The account SID of the task channel is deployed into
- `workspace_sid` - The workspace SID to create the task channel under
- `friendly_name` - The name of the task channel
- `unique_name` - The unique name of the task channel
- `channel_optimized_routing` - Whether the task channel should prioritise idle workers
- `date_created` - The date in RFC3339 format that the task channel was created
- `date_updated` - The date in RFC3339 format that the task channel was updated
- `url` - The URL of the task channel

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the task channel
