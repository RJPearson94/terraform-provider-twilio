---
page_title: "twilio_taskrouter_activity Resource - twilio"
subcategory: "TaskRouter"
description: |-
  
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

## Schema

### Required

- `friendly_name` (String) A human-readable label for the activity
- `workspace_sid` (String) The SID of the TaskRouter workspace. Changing this forces a new resource

### Optional

- `available` (Boolean) Whether the worker is available when in this activity state. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this activity
- `date_created` (String) The date and time the activity was created, in RFC 3339 format
- `date_updated` (String) The date and time the activity was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this activity by Twilio
- `url` (String) The absolute URL of the activity resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A activity can be imported using the `/Workspaces/{workspaceSid}/Activities/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_activity.activity /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
