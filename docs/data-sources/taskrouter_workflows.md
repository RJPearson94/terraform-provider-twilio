---
page_title: "Twilio TaskRouter Workflows"
subcategory: "TaskRouter"
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

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the workflows are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `workspace_sid`)
- `workspace_sid` - The SID of the workspace the workflows are associated with (Same as the `id`)
- `account_sid` - The account SID associated with the workflows
- `workflows` - A list of `workflow` blocks as documented below

---

A `workflow` block supports the following:

- `sid` - The SID of the workflow
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

- `read` - (Defaults to 10 minutes) Used when retrieving workflows
