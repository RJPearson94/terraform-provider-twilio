# taskrouter_activity

Manages a TaskRouter activity

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Workspace Activity"
  available     = true
}
```

## Argument Reference

The following arguments are supported:

* `friendly_name` - (Mandatory) The name of the Activity
* `workspaceSid` - (Mandatory) The Workspace SID to create the Activity under. Changing this forces a new resource to be created
* `available` - (Optional) Whether the Activity is available to accept tasks in Task Router. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Workspace (Same as the SID)
* `sid` - The SID of the Workspace (Same as the ID)
* `account_sid` - The Account SID of the Workspace is deployed into
* `workspaceSid` - The Workspace SID to create the Activity under.
* `friendly_name` - The name of the Activity
* `available` -  Whether the Activity is available to accept tasks in Task Router
* `date_created` - The date that the Activity was created
* `date_updated` - The date that the Activity was updated
* `url` - The url of the Workspace
