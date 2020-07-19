---
page_title: "Twilio Autopilot Task"
subcategory: "Autopilot"
---

# twilio_autopilot_task Resource

Manages a Autopilot Task

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
  actions       = <<EOF
{
 "actions": [
  {
   "say": {
    "speech": "Hello World"
   }
  }
 ]
}
EOF
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the task to
- `unique_name` - (Mandatory) The unique name of the task
- `friendly_name` - (Optional) The friendly name of the task
- `actions_url` - (Optional) The url to retrieve the actions json. Conflicts with actions.
- `actions` - (Optional) JSON string of an Autopilot task. Conflicts with actions_url.

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the task (Same as the SID)
- `sid` - The SID of the task (Same as the ID)
- `account_sid` - The Account SID associated with the task
- `assistant_sid` - The SID of the assistant to attach the task to
- `unique_name` - The unique name of the task
- `friendly_name` - The friendly name of the task
- `actions_url` - The url to retrieve the actions json
- `actions` - JSON string of an Autopilot task
- `date_created` - The date in RFC3339 format that the task was created
- `date_updated` - The date in RFC3339 format that the task was updated
- `url` - The url of the task resource
