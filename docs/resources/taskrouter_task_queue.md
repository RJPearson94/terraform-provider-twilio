---
page_title: "Twilio TaskRouter Task Queue"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The name of the task queue. The value cannot be an empty string
- `workspace_sid` - (Mandatory) The TaskRouter workspace SID to associate the task queue with. Changing this forces a new resource to be created
- `assignment_activity_sid` - (Optional) The assignment activity SID for the task queue
- `max_reserved_workers` - (Optional) The max number of workers to create a reservation for. The value must be between `1` and `50` (inclusive). The default value is `1`
- `target_workers` - (Optional) Worker selection criteria for any tasks that enter the task queue. The default value is `1==1`
- `task_order` - (Optional) How TaskRouter will assign workers tasks on the queue. Valid values are `LIFO` or `FIFO`. Default value is `FIFO`
- `reservation_activity_sid` - (Optional) The reservation activity SID for the task queue
- `event_callback_url` - (Optional) The callback URL for task queue events

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the task queue (Same as the `sid`)
- `sid` - The SID of the task queue (Same as the `id`)
- `account_sid` - The account SID of the task queue is deployed into
- `workspace_sid` - The workspace SID to create the task queue under
- `friendly_name` - The name of the task queue
- `event_callback_url` - The callback URL for task queue events
- `task_order` - How TaskRouter will assign workers tasks on the queue
- `assignment_activity_name` - The assignment activity name for the task queue
- `assignment_activity_sid` - The assignment activity SID for the task queue
- `reservation_activity_name` - The reservation activity name for the task queue
- `reservation_activity_sid` - The reservation activity SID for the task queue
- `target_workers` - Worker selection criteria for any tasks that enter the task queue
- `max_reserved_workers` - The max number of workers to create a reservation for
- `date_created` - The date in RFC3339 format that the task queue was created
- `date_updated` - The date in RFC3339 format that the task queue was updated
- `url` - The URL of the task queue

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the task queue
- `update` - (Defaults to 10 minutes) Used when updating the task queue
- `read` - (Defaults to 5 minutes) Used when retrieving the task queue
- `delete` - (Defaults to 10 minutes) Used when deleting the task queue

## Import

A task queue can be imported using the `/Workspaces/{workspaceSid}/TaskQueues/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_task_queue.task_queue /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
