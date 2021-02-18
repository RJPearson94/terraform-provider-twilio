---
page_title: "Twilio Autopilot Task Field"
subcategory: "Autopilot"
---

# twilio_autopilot_task_field Resource

Manages an Autopilot task field. See the [API docs](https://www.twilio.com/docs/autopilot/api/task-field) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
}

resource "twilio_autopilot_task_field" "task_field" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  task_sid      = twilio_autopilot_task.task.sid
  unique_name   = "test"
  field_type    = "Twilio.YES_NO"
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to associate the task field with. Changing this forces a new resource to be created
- `task_sid` - (Mandatory) The SID of the task to associate the task field with. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the field. Changing this forces a new resource to be created
- `field_type` - (Mandatory) The type of field

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field (Same as the `sid`)
- `sid` - The SID of the field (Same as the `id`)
- `account_sid` - The account SID associated with the field
- `assistant_sid` - The SID of the assistant to attach the task to
- `task_sid` - The SID of the task to attach the field to
- `unique_name` - The unique name of the field
- `field_type` - The type of field
- `date_created` - The date in RFC3339 format that the field was created
- `date_updated` - The date in RFC3339 format that the field was updated
- `url` - The URL of the field

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the field
- `read` - (Defaults to 5 minutes) Used when retrieving the field
- `delete` - (Defaults to 10 minutes) Used when deleting the field

## Import

A task field can be imported using the `/Assistants/{assistantSid}/Tasks/{taskSid}/Fields/{sid}` format, e.g.

```shell
terraform import twilio_autopilot_task_field.task_field /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Fields/UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
