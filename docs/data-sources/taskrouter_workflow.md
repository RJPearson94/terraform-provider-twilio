---
page_title: "Twilio TaskRouter Workflow"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the workflow is associated with
- `sid` - (Mandatory) The SID of the workflow

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the workflow (Same as the SID)
- `sid` - The SID of the workflow (Same as the ID)
- `account_sid` - The account SID of the workflow is deployed into
- `workspace_sid` - The workspace SID to create the workflow under
- `friendly_name` - The name of the workflow
- `configuration` - JSON string of workflow configuration
- `assignment_callback_url` - Assignment callback URL
- `fallback_assignment_callback_url` - Fallback assignment callback URL
- `task_reservation_timeout` - Maximum time the task can be unassigned for before it times out
- `document_content_type` - The MIME type of the document
- `date_created` - The date in RFC3339 format that the workflow was created
- `date_updated` - The date in RFC3339 format that the workflow was updated
- `url` - The URL of the workflow

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the workflow
