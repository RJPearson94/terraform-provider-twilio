---
page_title: "Twilio TaskRouter Workspace Configuration"
subcategory: "TaskRouter"
---

# twilio_taskrouter_workspace_configuration Resource

Manages a TaskRouter workspace configuration. This resource is part of the TaskRouter Workspace API but is managed as a separate resource in Terraform to prevent cyclic dependencies. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workspace) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

!> This resource modifies the Twilio TaskRouter workspace configuration. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

!> Removing the `default_activity_sid` or `timeout_activity_sid` from your configuration will cause the corresponding value to be retained after a Terraform apply. If you want to change any of the values you will need to either create a new `twilio_taskrouter_activity` resource and set the argument to the generated `sid`. Alternatively, you can set the activity SID to one of the activities that were created when the service was created.

!> Twilio will throw an error if you try to delete an activity if it's attached as either the default or timeout activity SID of a workspace. If you use the `twilio_taskrouter_activity` resource, you will need to either remove the resource from the Terraform state or update the configuration to reference an activity that is not known by Terraform i.e. one of the activities created when the workspace was created, then the activity resource can be deleted.

## Example Usage

### Basic

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
}
```

### With Activity Resource

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Activity"
  available     = true
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  default_activity_sid = twilio_taskrouter_activity.activity.sid
}
```

### With Activities Data Source

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

data "twilio_taskrouter_activities" "activities" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Offline"
}

resource "twilio_taskrouter_workspace_configuration" "workspace_configuration" {
  workspace_sid        = twilio_taskrouter_workspace.workspace.sid
  timeout_activity_sid = data.twilio_taskrouter_activities.activities.activities[0].sid
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace
- `default_activity_sid` - (Optional) SID of the default activity
- `timeout_activity_sid` - (Optional) SID of the timeout activity

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the workspace (Same as the `workspace_sid`)
- `workspace_sid` - The SID of the workspace (Same as the `id`)
- `default_activity_name` - Name of the default activity
- `default_activity_sid` - SID of the default activity
- `timeout_activity_name` - Name of the timeout activity
- `timeout_activity_sid` - SID of the timeout activity

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the workspace configuration
- `update` - (Defaults to 10 minutes) Used when updating the workspace configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the workspace configuration
