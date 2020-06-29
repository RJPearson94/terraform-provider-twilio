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
- `event_callback_url` - (Optional) The Event Callback URL
- `events_filter` - (Optional) The Event Callback Filter
- `multi_task_enabled` - (Optional) Whether or not Multitasking is enabled
- `template` - (Optional) Task Router template to use
- `prioritize_queue_order` - (Optional) Determine how TaskRouter prioritizes incoming tasks. Options are `LIFO` or `FIFO`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the workspace (Same as the SID)
- `sid` - The SID of the workspace (Same as the ID)
- `account_sid` - The Account SID of the workspace is deployed into
- `friendly_name` - The name of the workspace
- `event_callback_url` - The Event Callback URL
- `events_filter` - The Event Callback Filter
- `multi_task_enabled` - Whether or not Multitasking is enabled
- `template` - Task Router template to use
- `prioritize_queue_order` - Determine how TaskRouter prioritizes incoming
- `default_activity_name` - Name of default activity
- `default_activity_sid` - Sid of default activity
- `timeout_activity_name` - Name of timeout activity
- `timeout_activity_sid` - Sid of timeout activity
- `date_created` - The date in RFC3339 format that the workspace was created
- `date_updated` - The date in RFC3339 format that the workspace was updated
- `url` - The url of the workspace
