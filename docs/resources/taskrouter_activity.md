---
page_title: "Twilio TaskRouter Activity"
subcategory: "TaskRouter"
---

# twilio_taskrouter_activity Resource

Manages a TaskRouter activity. See the [API docs](https://www.twilio.com/docs/taskrouter/api/activity) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

!> Twilio will throw an error if you try to delete an activity if it's attached as either the default or timeout activity SID of a workspace (this can be managed via the `twilio_taskrouter_workspace_configuration` resource). If you use this resource, you will need to either remove the resource from the Terraform state or update the configuration to reference an activity that is not known by Terraform i.e. one of the activities created when the workspace was created, then the activity resource can be deleted.

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
- `workspace_sid` - (Mandatory) The TaskRouter workspace SID to associate the activity with. Changing this forces a new resource to be created
- `available` - (Optional) Whether the activity is available to accept tasks in TaskRouter. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the activity (Same as the `sid`)
- `sid` - The SID of the activity (Same as the `id`)
- `account_sid` - The account SID of the activity is deployed into
- `workspace_sid` - The workspace SID to create the activity under.
- `friendly_name` - The name of the activity
- `available` - Whether the activity is available to accept tasks in TaskRouter
- `date_created` - The date in RFC3339 format that the activity was created
- `date_updated` - The date in RFC3339 format that the activity was updated
- `url` - The URL of the activity

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the activity
- `update` - (Defaults to 10 minutes) Used when updating the activity
- `read` - (Defaults to 5 minutes) Used when retrieving the activity
- `delete` - (Defaults to 10 minutes) Used when deleting the activity

## Import

A activity can be imported using the `/Workspaces/{workspaceSid}/Activities/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_activity.activity /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
