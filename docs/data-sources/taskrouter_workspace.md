---
page_title: "twilio_taskrouter_workspace Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_workspace Data Source

Use this data source to access information about an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workspace) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_workspace" "workspace" {
  sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "workspace" {
  value = data.twilio_taskrouter_workspace.workspace
}
```

## Schema

### Required

- `sid` (String) The SID of the workspace to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this workspace
- `date_created` (String) The date and time the workspace was created, in RFC 3339 format
- `date_updated` (String) The date and time the workspace was last updated, in RFC 3339 format
- `default_activity_name` (String) The name of the default activity for the workspace
- `default_activity_sid` (String) The SID of the default activity for the workspace
- `event_callback_url` (String) The URL to call when an event is fired in the workspace
- `event_filters` (List of String) A list of event types the workspace is subscribed to
- `friendly_name` (String) A human-readable label for the workspace
- `id` (String) The ID of this resource.
- `multi_task_enabled` (Boolean) Whether multi-tasking is enabled for the workspace
- `prioritize_queue_order` (String) The order in which task queues are prioritized
- `timeout_activity_name` (String) The name of the timeout activity for the workspace
- `timeout_activity_sid` (String) The SID of the timeout activity for the workspace
- `url` (String) The absolute URL of the workspace resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
