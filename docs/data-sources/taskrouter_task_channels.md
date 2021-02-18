---
page_title: "Twilio TaskRouter Task Channels"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the task channels are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `workspace_sid`)
- `workspace_sid` - The SID of the workspace the task channels are associated with (Same as the `id`)
- `account_sid` - The account SID associated with the task channels
- `task_channels` - A list of `task_channel` blocks as documented below

---

A `task_channel` block supports the following:

- `sid` - The SID of the task channel
- `friendly_name` - The name of the task channel
- `unique_name` - The unique name of the task channel
- `channel_optimized_routing` - Whether the task channel should prioritise idle workers
- `date_created` - The date in RFC3339 format that the task channel was created
- `date_updated` - The date in RFC3339 format that the task channel was updated
- `url` - The URL of the task channel

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving task channels
