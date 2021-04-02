---
page_title: "Twilio TaskRouter Worker"
subcategory: "TaskRouter"
---

# twilio_taskrouter_worker Resource

Manages a TaskRouter worker. See the [API docs](https://www.twilio.com/docs/taskrouter/api/worker) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

!> Removing the `activity_sid` from you configuration will cause the SID to be set to workpace default activity SID

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Worker"
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The TaskRouter workspace SID to associate the worker with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the worker
- `attributes` - (Optional) JSON string of worker attributes. The default value is `{}`
- `activity_sid` - (Optional) Activity SID to be assigned to the worker

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the worker (Same as the `sid`)
- `sid` - The SID of the worker (Same as the `id`)
- `account_sid` - The account SID of the worker is deployed into
- `workspace_sid` - The workspace SID to create the worker under
- `friendly_name` - The name of the worker
- `attributes` - JSON string of worker attributes
- `activity_sid` - Activity SID to be assigned to the worker
- `activity_name` - Friendly name of the activity
- `available` - Is the worker available to receive tasks
- `date_created` - The date in RFC3339 format that the worker was created
- `date_updated` - The date in RFC3339 format that the worker was updated
- `date_status_changed` - The date in RFC3339 format that the worker status was changed
- `url` - The URL of the worker

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the worker
- `update` - (Defaults to 10 minutes) Used when updating the worker
- `read` - (Defaults to 5 minutes) Used when retrieving the worker
- `delete` - (Defaults to 10 minutes) Used when deleting the worker

## Import

A worker can be imported using the `/Workspaces/{workspaceSid}/Workers/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_worker.worker /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
