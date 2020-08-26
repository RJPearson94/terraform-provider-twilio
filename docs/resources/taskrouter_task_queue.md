---
page_title: "Twilio TaskRouter Task Queue"
subcategory: "TaskRouter"
---

# twilio_taskrouter_task_queue Resource

Manages a TaskRouter task queue

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

- `friendly_name` - (Mandatory) The name of the task queue
- `workspaceSid` - (Mandatory) The workspace SID to create the task queue under. Changing this forces a new resource to be created
- `assignment_activity_sid` - (Optional) The assignment activity SID for the task queue
- `max_reserved_workers` - (Optional) The max number of workers to create a reservation for
- `target_workers` - (Optional) Worker selection criteria for any tasks that enter the task queue
- `task_order` - (Optional) How TaskRouter will assign workers tasks on the queue
- `reservation_activity_sid` - (Optional) The reservation activity SID for the task queue

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the task queue (Same as the SID)
- `sid` - The SID of the task queue (Same as the ID)
- `account_sid` - The account SID of the task queue is deployed into
- `workspaceSid` - The workspace SID to create the task queue under
- `friendly_name` - The name of the task queue
- `event_callback_url` - The callback URL for task queue Events
- `task_order` - How TaskRouter will assign workers tasks on the queue
- `assignment_activity_name` - The assignment activity name for the task queue
- `assignment_activity_sid` - The assignment activity SID for the task queue
- `reservation_activity_name` - The reservation activity name for the task queue
- `reservation_activity_sid` - The reservation activity SID for the task queue
- `target_workers` - Worker selection criteria for any tasks that enter the task queue
- `max_reserved_workers` - The max number of workers to create a reservation for
- `date_created` - The date in RFC3339 format that the task queue was created
- `date_updated` - The date in RFC3339 format that the task queue was updated
- `url` - The url of the task queue

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the task queue
- `update` - (Defaults to 10 minutes) Used when updating the task queue
- `read` - (Defaults to 5 minutes) Used when retrieving the task queue
- `delete` - (Defaults to 10 minutes) Used when deleting the task queue

## Import

A task queue can be imported using the `"/Workspaces/{workspaceSid}/TaskQueues/{sid}"` format, e.g.

```shell
terraform import twilio_taskrouter_task_queue.task_queue /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
