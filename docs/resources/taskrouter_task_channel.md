---
page_title: "Twilio TaskRouter Task Channel"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The TaskRouter workspace SID to associate the task channel with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the task channel
- `unique_name` - (Mandatory) The unique name of the task channel. Changing this forces a new resource to be created
- `channel_optimized_routing` - (Optional) Whether the task channel should prioritise idle workers

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

- `create` - (Defaults to 10 minutes) Used when creating the task channel
- `update` - (Defaults to 10 minutes) Used when updating the task channel
- `read` - (Defaults to 5 minutes) Used when retrieving the task channel
- `delete` - (Defaults to 10 minutes) Used when deleting the task channel

## Import

A task channel can be imported using the `/Workspaces/{workspaceSid}/TaskChannels/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_task_channel.task_channel /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskChannels/TCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
