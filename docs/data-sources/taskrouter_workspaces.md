---
page_title: "twilio_taskrouter_workspaces Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
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
data "twilio_taskrouter_workspaces" "flex" {
  friendly_name = "Flex Task Assignment"
}

output "flex_workspace" {
  value = data.twilio_taskrouter_workspaces.flex.workspaces[0].sid
}
```

## Schema

### Optional

- `friendly_name` (String) Filter workspaces by friendly name
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the workspaces
- `id` (String) The ID of this resource.
- `workspaces` (List of Object) A list of workspaces (see [below for nested schema](#nestedatt--workspaces))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--workspaces"></a>
### Nested Schema for `workspaces`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `default_activity_name` (String)
- `default_activity_sid` (String)
- `event_callback_url` (String)
- `event_filters` (List of String)
- `friendly_name` (String)
- `multi_task_enabled` (Boolean)
- `prioritize_queue_order` (String)
- `sid` (String)
- `timeout_activity_name` (String)
- `timeout_activity_sid` (String)
- `url` (String)
