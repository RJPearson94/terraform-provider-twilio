---
page_title: "Twilio TaskRouter Workspace"
subcategory: "TaskRouter"
---

# twilio_taskrouter_workspace Resource

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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the workspace

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
- `url` - The URL of the workspace

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the workspace
