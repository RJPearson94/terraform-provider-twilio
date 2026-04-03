---
page_title: "twilio_taskrouter_workflow Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
---

# twilio_taskrouter_workflow Data Source

Use this data source to access information about an existing TaskRouter workflow. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workflow) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "workflow" {
  value = data.twilio_taskrouter_workflow.workflow
}
```

## Schema

### Required

- `sid` (String) The SID of the workflow to retrieve
- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this workflow
- `assignment_callback_url` (String) The URL to call when a task is assigned to a worker
- `configuration` (String) A JSON string of the workflow configuration
- `date_created` (String) The date and time the workflow was created, in RFC 3339 format
- `date_updated` (String) The date and time the workflow was last updated, in RFC 3339 format
- `document_content_type` (String) The MIME type of the workflow document
- `fallback_assignment_callback_url` (String) The URL to call when a task assignment event is not handled by the primary callback
- `friendly_name` (String) A human-readable label for the workflow
- `id` (String) The ID of this resource.
- `task_reservation_timeout` (Number) The timeout in seconds for a task reservation
- `url` (String) The absolute URL of the workflow resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
