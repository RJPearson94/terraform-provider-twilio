---
page_title: "Twilio TaskRouter Activity"
subcategory: "TaskRouter"
---

# twilio_taskrouter_activity Resource

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

- `friendly_name` - (Mandatory) The name of the activity
- `workspaceSid` - (Mandatory) The workspace SID to create the activity under. Changing this forces a new resource to be created
- `available` - (Optional) Whether the activity is available to accept tasks in TaskRouter. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the activity (Same as the SID)
- `sid` - The SID of the activity (Same as the ID)
- `account_sid` - The account SID of the activity is deployed into
- `workspaceSid` - The workspace SID to create the activity under.
- `friendly_name` - The name of the activity
- `available` -  Whether the activity is available to accept tasks in Task Router
- `date_created` - The date in RFC3339 format that the activity was created
- `date_updated` - The date in RFC3339 format that the activity was updated
- `url` - The url of the activity

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the activity
- `update` - (Defaults to 10 minutes) Used when updating the activity
- `read` - (Defaults to 5 minutes) Used when retrieving the activity
- `delete` - (Defaults to 10 minutes) Used when deleting the activity
