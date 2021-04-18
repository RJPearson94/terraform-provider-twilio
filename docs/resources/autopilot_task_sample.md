---
page_title: "Twilio Autopilot Task Sample"
subcategory: "Autopilot"
---

# twilio_autopilot_task_sample Resource

Manages an Autopilot task sample. See the [API docs](https://www.twilio.com/docs/autopilot/api/task-sample) for more information

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

resource "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = twilio_autopilot_task.task.assistant_sid
  task_sid      = twilio_autopilot_task.task.sid
  language      = "en-US"
  tagged_text   = "test"
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to associate the sample with. Changing this forces a new resource to be created
- `task_sid` - (Mandatory) The SID of the task to associate the sample with. Changing this forces a new resource to be created
- `language` - (Mandatory) The language of the sample
- `tagged_text` - (Mandatory) The labelled/ tagged sample text
- `source_channel` - (Optional) The channel the sample was captured on. Valid values are `voice`, `sms`, `chat`, `alexa`, `google-assistant` or `slack`. The default value is `voice`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the sample (Same as the `sid`)
- `sid` - The SID of the sample (Same as the `id`)
- `account_sid` - The account SID associated with the sample
- `assistant_sid` - The SID of the assistant to attach the sample to
- `task_sid` - The SID of the task to attach the sample to
- `language` - The language of the sample
- `tagged_text` - The labelled/ tagged sample text
- `source_channel` - The channel the sample was captured on
- `date_created` - The date in RFC3339 format that the sample was created
- `date_updated` - The date in RFC3339 format that the sample was updated
- `url` - The URL of the sample

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the sample
- `update` - (Defaults to 10 minutes) Used when updating the sample
- `read` - (Defaults to 5 minutes) Used when retrieving the sample
- `delete` - (Defaults to 10 minutes) Used when deleting the sample

## Import

A task sample can be imported using the `/Assistants/{assistantSid}/Tasks/{taskSid}/Samples/{sid}` format, e.g.

```shell
terraform import twilio_autopilot_task_sample.task_sample /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Tasks/UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Samples/UFXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
