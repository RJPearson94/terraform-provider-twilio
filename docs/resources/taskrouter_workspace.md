---
page_title: "Twilio TaskRouter Workspace"
subcategory: "TaskRouter"
---

# twilio_taskrouter_workspace Resource

Manages a TaskRouter workspace

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "Test Workspace"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
  template               = "FIFO"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The name of the workspace
- `event_callback_url` - (Optional) The event callback URL
- `event_filters` - (Optional) list of event callback filters
- `multi_task_enabled` - (Optional) Whether or not multitasking is enabled
- `template` - (Optional) TaskRouter template to use
- `prioritize_queue_order` - (Optional) Determine how TaskRouter prioritizes incoming tasks. Options are `LIFO` or `FIFO`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the workspace (Same as the SID)
- `sid` - The SID of the workspace (Same as the ID)
- `account_sid` - The account SID of the workspace is deployed into
- `friendly_name` - The name of the workspace
- `event_callback_url` - The event callback URL
- `event_filters` - The event callback filter
- `multi_task_enabled` - Whether or not multitasking is enabled
- `template` - TaskRouter template to use
- `prioritize_queue_order` - Determine how TaskRouter prioritizes incoming
- `default_activity_name` - Name of default activity
- `default_activity_sid` - SID of default activity
- `timeout_activity_name` - Name of timeout activity
- `timeout_activity_sid` - SID of timeout activity
- `date_created` - The date in RFC3339 format that the workspace was created
- `date_updated` - The date in RFC3339 format that the workspace was updated
- `url` - The url of the workspace

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the workspace
- `update` - (Defaults to 10 minutes) Used when updating the workspace
- `read` - (Defaults to 5 minutes) Used when retrieving the workspace
- `delete` - (Defaults to 10 minutes) Used when deleting the workspace

## Import

A workspace can be imported using the `"/Workspaces/{sid}"` format, e.g.

```shell
terraform import twilio_taskrouter_workspace.workspace /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
