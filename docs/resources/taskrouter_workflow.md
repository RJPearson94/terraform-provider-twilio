---
page_title: "Twilio TaskRouter Workflow"
subcategory: "TaskRouter"
---

# twilio_taskrouter_workflow Resource

Manages a TaskRouter workflow. See the [API docs](https://www.twilio.com/docs/taskrouter/api/workflow) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "twilio-test"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Queue"
}

resource "twilio_taskrouter_workflow" "workflow" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "Test Workflow"
  configuration = jsonencode({
    "task_routing" : {
      "filters" : [],
      "default_filter" : {
        "queue" : twilio_taskrouter_task_queue.task_queue.sid
      }
    }
  })
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The TaskRouter workspace SID to associate the workflow with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the workflow. The value cannot be an empty string
- `configuration` - (Mandatory) JSON string of workflow configuration
- `assignment_callback_url` - (Optional) Assignment callback URL. The default value is an empty string/ no configuration specified
- `fallback_assignment_callback_url` - (Optional) Fallback assignment callback URL. The default value is an empty string/ no configuration specified
- `task_reservation_timeout` - (Optional) Maximum time the task can be unassigned for before it times out. The timeout must be between `1` and `86400` seconds (inclusive). The default value is `120`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the workflow (Same as the `sid`)
- `sid` - The SID of the workflow (Same as the `id`)
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

- `create` - (Defaults to 10 minutes) Used when creating the workflow
- `update` - (Defaults to 10 minutes) Used when updating the workflow
- `read` - (Defaults to 5 minutes) Used when retrieving the workflow
- `delete` - (Defaults to 10 minutes) Used when deleting the workflow

## Import

A workflow can be imported using the `/Workspaces/{workspaceSid}/Workflows/{sid}` format, e.g.

```shell
terraform import twilio_taskrouter_workflow.workflow /Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Workflows/WFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
