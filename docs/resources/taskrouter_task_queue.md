---
page_title: "Twilio TaskRouter Task Queue"
subcategory: "TaskRouter"
---

# twilio_taskrouter_task_queue Resource

Manages a TaskRouter Task Queue

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

* `friendly_name` - (Mandatory) The name of the task queue
* `workspaceSid` - (Mandatory) The Workspace SID to create the task queue under. Changing this forces a new resource to be created
* `assignment_activity_sid` - (Optional) The Assignment Activity SID for the task queue
* `max_reserved_workers` - (Optional) The max number of workers to create a reservation for. Default is 1
* `target_workers` - (Optional) Worker selection criteria for any tasks that enter the task queue
* `task_order` - (Optional) How TaskRouter will assign workers tasks on the queue. Default is FIFO
* `reservation_activity_sid` - (Optional) The Reservation Activity SID for the task queue

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the task queue (Same as the SID)
* `sid` - The SID of the task queue (Same as the ID)
* `account_sid` - The Account SID of the task queue is deployed into
* `workspaceSid` - The Workspace SID to create the task queue under
* `friendly_name` - The name of the task queue
* `event_callback_url` - The callback URL for task queue Events
* `task_order` - How TaskRouter will assign workers tasks on the queue
* `assignment_activity_name` - The Assignment Activity Name for the task queue
* `assignment_activity_sid` - The Assignment Activity SID for the task queue
* `reservation_activity_name` - The Reservation Activity Name for the task queue
* `reservation_activity_sid` - The Reservation Activity SID for the task queue
* `target_workers` - Worker selection criteria for any tasks that enter the task queue
* `max_reserved_workers` - The max number of workers to create a reservation for
* `date_created` - The date that the task queue was created
* `date_updated` - The date that the task queue was updated
* `url` - The url of the task queue
