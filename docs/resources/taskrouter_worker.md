---
page_title: "twilio_taskrouter_worker Resource - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_worker Resource

Manages a TaskRouter worker. See the [API docs](https://www.twilio.com/docs/taskrouter/api/worker) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

!> Removing the `activity_sid` from your configuration will cause the value to be retained after a Terraform apply. If you want to change the `activity_sid` value you will need to either create a new `twilio_taskrouter_activity` resource and set the `activity_sid` to the generated `sid` alternatively you can set the `activity_sid` to be the workspace default by using the `default_activity_sid` attribute on the `twilio_taskrouter_workspace` resource

## Example Usage

### Basic

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

### Custom activity

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "test"
  available     = true
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Worker"
  activity_sid  = twilio_taskrouter_activity.activity.sid
}
```

### Explicitly set the workspace default activity SID

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Worker"
  activity_sid  = twilio_taskrouter_workspace.workspace.default_activity_sid
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the worker
- `workspace_sid` (String) The SID of the TaskRouter workspace. Changing this forces a new resource

### Optional

- `activity_sid` (String) The SID of the activity the worker should be set to
- `attributes` (String) A JSON string of attributes for the worker. Defaults to `{}`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this worker
- `activity_name` (String) The friendly name of the worker's current activity
- `available` (Boolean) Whether the worker is available to receive tasks
- `date_created` (String) The date and time the worker was created, in RFC 3339 format
- `date_status_changed` (String) The date and time the worker's activity status last changed, in RFC 3339 format
- `date_updated` (String) The date and time the worker was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this worker by Twilio
- `url` (String) The absolute URL of the worker resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A worker can be imported using the `/Workspaces/{workspaceSid}/Workers/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_worker.worker /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workers/WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
