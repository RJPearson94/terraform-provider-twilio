---
page_title: "Twilio TaskRouter Workspaces"
subcategory: "TaskRouter"
---

# twilio_taskrouter_workspaces Data Source

Use this data source to access information about existing TaskRouter workspaces. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workspace) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

### Basic

```hcl
data "twilio_taskrouter_workspaces" "workspaces" {}

output "workspaces" {
  value = data.twilio_taskrouter_workspaces.workspaces
}
```

### Search for Flex Task Assignment Workspace (applicable to Flex projects)

```hcl
data "twilio_taskrouter_workspaces" "workspaces" {
  friendly_name = "Flex Task Assignment"
}

output "flex_workspace" {
  value = data.twilio_taskrouter_workspaces.workspaces[0].sid
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Optional) Search for all workspaces which have the friendly name specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `account_sid`)
- `account_sid` - The account SID associated with the workspaces (Same as the `id`)
- `workspaces` - A list of `workspace` blocks as documented below

---

A `workspace` block supports the following:

- `sid` - The SID of the workspace (Same as the `id`)
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

- `read` - (Defaults to 10 minutes) Used when retrieving workspaces
