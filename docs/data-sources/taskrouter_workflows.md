---
page_title: "twilio_taskrouter_workflows Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_workflows Data Source

Use this data source to access information about the workflows associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workflow) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_workflows" "workflows" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "workflows" {
  value = data.twilio_taskrouter_workflows.workflows
}
```

## Schema

### Required

- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `friendly_name` (String) Filter workflows by friendly name
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the workflows
- `id` (String) The ID of this resource.
- `workflows` (List of Object) A list of workflows in the workspace (see [below for nested schema](#nestedatt--workflows))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--workflows"></a>
### Nested Schema for `workflows`

Read-Only:

- `assignment_callback_url` (String)
- `configuration` (String)
- `date_created` (String)
- `date_updated` (String)
- `document_content_type` (String)
- `fallback_assignment_callback_url` (String)
- `friendly_name` (String)
- `sid` (String)
- `task_reservation_timeout` (Number)
- `url` (String)
