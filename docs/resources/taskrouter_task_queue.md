# twilio_taskrouter_task_queue

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

* `friendly_name` - (Mandatory) The name of the Activity
* `workspaceSid` - (Mandatory) The Workspace SID to create the Activity under. Changing this forces a new resource to be created
* `assignment_activity_sid` - (Optional) The Assignment Activity SID for the Task Queue
* `max_reserved_workers` - (Optional) The max number of workers to create a reservation for. Default is 1
* `target_workers` - (Optional) Worker selection criteria for any tasks that enter the Task Queue
* `task_order` - (Optional) How TaskRouter will assign workers tasks on the queue. Default is FIFO
* `reservation_activity_sid` - (Optional) The Reservation Activity SID for the Task Queue

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Task Queue (Same as the SID)
* `sid` - The SID of the Task Queue (Same as the ID)
* `account_sid` - The Account SID of the Task Queue is deployed into
* `workspaceSid` - The Workspace SID to create the Task Queue under
* `friendly_name` - The name of the Task Queue
* `event_callback_url` - The callback URL for Task Queue Events
* `task_order` - How TaskRouter will assign workers tasks on the queue
* `assignment_activity_name` - The Assignment Activity Name for the Task Queue
* `assignment_activity_sid` - The Assignment Activity SID for the Task Queue
* `reservation_activity_name` - The Reservation Activity Name for the Task Queue
* `reservation_activity_sid` - The Reservation Activity SID for the Task Queue
* `target_workers` - Worker selection criteria for any tasks that enter the Task Queue
* `max_reserved_workers` - The max number of workers to create a reservation for
* `date_created` - The date that the Task Queue was created
* `date_updated` - The date that the Task Queue was updated
* `url` - The url of the Task Queue
