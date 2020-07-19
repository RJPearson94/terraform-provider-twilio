---
page_title: "Twilio Autopilot Task Field"
subcategory: "Autopilot"
---

# twilio_autopilot_task_field Resource

Manages a Autopilot Task Field

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

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the task field to. Changing this forces a new resource to be created
- `task_sid` - (Mandatory) The SID of the task to attach the field to. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the field. Changing this forces a new resource to be created
- `field_type` - (Mandatory) The type of field

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field (Same as the SID)
- `sid` - The SID of the field (Same as the ID)
- `account_sid` - The Account SID associated with the field
- `assistant_sid` - The SID of the assistant to attach the task to
- `task_sid` - The SID of the task to attach the field to
- `unique_name` - The unique name of the field
- `field_type` - The type of field
- `date_created` - The date in RFC3339 format that the field was created
- `date_updated` - The date in RFC3339 format that the field was updated
- `url` - The url of the field
