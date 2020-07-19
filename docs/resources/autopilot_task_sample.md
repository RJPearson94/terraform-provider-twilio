---
page_title: "Twilio Autopilot Task Sample"
subcategory: "Autopilot"
---

# twilio_autopilot_task_sample Resource

Manages a Autopilot Task Sample

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
}

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_task.task.assistant_sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "en-US"
  tagged_text   = "test"
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the task sample to
- `task_sid` - (Mandatory) The SID of the task to attach the sample to
- `language` - (Mandatory) The language of the sample
- `tagged_text` - (Mandatory) The labelled/ tagged sample text
- `source_channel` - (Optional) The channel the sample was captured on

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the sample (Same as the SID)
- `sid` - The SID of the sample (Same as the ID)
- `account_sid` - The Account SID associated with the sample
- `assistant_sid` - The SID of the assistant to attach the task to
- `task_sid` - The SID of the task to attach the sample to
- `language` - The language of the sample
- `tagged_text` - The labelled/ tagged sample text
- `source_channel` - The channel the sample was captured on
- `date_created` - The date in RFC3339 format that the sample was created
- `date_updated` - The date in RFC3339 format that the sample was updated
- `url` - The url of the sample
