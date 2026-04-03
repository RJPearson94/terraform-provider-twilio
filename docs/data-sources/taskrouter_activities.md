---
page_title: "twilio_taskrouter_activities Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_activities Data Source

Use this data source to access information about the activities associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/activity) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_activities" "activities" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "activities" {
  value = data.twilio_taskrouter_activities.activities
}
```

## Schema

### Required

- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `available` (Boolean) Filter activities by availability
- `friendly_name` (String) Filter activities by friendly name
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the activities
- `activities` (List of Object) A list of activities in the workspace (see [below for nested schema](#nestedatt--activities))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--activities"></a>
### Nested Schema for `activities`

Read-Only:

- `available` (Boolean)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
- `url` (String)
