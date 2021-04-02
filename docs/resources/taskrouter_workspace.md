---
page_title: "Twilio TaskRouter Workspace"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The name of the workspace
- `event_callback_url` - (Optional) The event callback URL
- `event_filters` - (Optional) list of event callback filters. Valid values are `task.created`,`task.completed`,`task.canceled`,`task.deleted`,`task.updated`,`task.wrapup`,`task-queue.entered`,`task-queue.moved`,`task-queue.timeout`,`reservation.created`,`reservation.accepted`,`reservation.rejected`,`reservation.timeout`,`reservation.canceled`,`reservation.rescinded`,`reservation.completed`,`workflow.entered`,`workflow.timeout`,`workflow.target-matched`,`worker.activity.update`,`worker.attributes.update`,`worker.capacity.update` or `worker.channel.availability.update`
- `multi_task_enabled` - (Optional) Whether or not multitasking is enabled
- `template` - (Optional) TaskRouter template to use. Valid values are `NONE` or `FIFO`. The default value is `NONE`
- `prioritize_queue_order` - (Optional) Determine how TaskRouter prioritizes incoming tasks. Valid values are `LIFO` or `FIFO`. The default value is `FIFO`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the workspace (Same as the `sid`)
- `sid` - The SID of the workspace (Same as the `id`)
- `account_sid` - The account SID of the workspace is deployed into
- `friendly_name` - The name of the workspace
- `event_callback_url` - The event callback URL
- `event_filters` - The event callback filter
- `multi_task_enabled` - Whether or not multitasking is enabled
- `prioritize_queue_order` - Determine how TaskRouter prioritizes incoming
- `default_activity_name` - Name of default activity
- `default_activity_sid` - SID of default activity
- `timeout_activity_name` - Name of timeout activity
- `timeout_activity_sid` - SID of timeout activity
- `date_created` - The date in RFC3339 format that the workspace was created
- `date_updated` - The date in RFC3339 format that the workspace was updated
- `url` - The URL of the workspace

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the workspace
- `update` - (Defaults to 10 minutes) Used when updating the workspace
- `read` - (Defaults to 5 minutes) Used when retrieving the workspace
- `delete` - (Defaults to 10 minutes) Used when deleting the workspace

## Import

A workspace can be imported using the `/Workspaces/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_workspace.workspace /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> "template" cannot be imported
