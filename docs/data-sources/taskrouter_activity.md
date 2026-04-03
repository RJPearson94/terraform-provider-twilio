---
page_title: "twilio_taskrouter_activity Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_activity Data Source

Use this data source to access information about an existing TaskRouter activity. See the [API docs](https://www.twilio.com/docs/taskrouter/api/activity) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_activity" "activity" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "activity" {
  value = data.twilio_taskrouter_activity.activity
}
```

## Schema

### Required

- `sid` (String) The SID of the activity to retrieve
- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this activity
- `available` (Boolean) Whether the worker is available when in this activity state
- `date_created` (String) The date and time the activity was created, in RFC 3339 format
- `date_updated` (String) The date and time the activity was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the activity
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the activity resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
