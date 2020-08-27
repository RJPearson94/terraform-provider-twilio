---
page_title: "Twilio Autopilot Model Build"
subcategory: "Autopilot"
---

# twilio_autopilot_model_build Resource

Manages a Autopilot model build

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

resource "twilio_autopilot_model_build" "model_build" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"

  polling {
    enabled = true
  }

  depends_on = [
    twilio_autopilot_task_sample.task_sample
  ]
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the model build to. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the model build
- `status_callback` - (Optional) The callback url to post build statuses to. Changing this forces a new resource to be created
- `polling` - (Optional) A `polling` block as documented below.

---

A `polling` block supports the following:

- `enabled` - (Required) Enable or or disable polling of the build.
- `max_attempts` - (Optional) The maximum number of polling attempts. Default is 24
- `delay_in_ms` - (Optional) The time in milliseconds to wait between polling attempts. Default is 5000ms

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the model build (Same as the SID)
- `sid` - The SID of the model build (Same as the ID)
- `account_sid` - The Account SID associated with the model build
- `unique_name` - The unique name of the model build
- `status_callback` - The callback url to post build statuses to
- `status` - The current build status
- `error_code` - The error code of the build if the status is failed
- `build_duration` - The duration of the build (in seconds)
- `date_created` - The date in RFC3339 format that the model build was created
- `date_updated` - The date in RFC3339 format that the model build was updated
- `url` - The url of the model build resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the model build
- `update` - (Defaults to 10 minutes) Used when updating the model build
- `read` - (Defaults to 5 minutes) Used when retrieving the model build
- `delete` - (Defaults to 10 minutes) Used when deleting the model build

!> When polling is enabled, each request is constrained by the read timeout defined above

## Import

A model build can be imported using the `/Assistants/{assistantSid}/ModelBuilds/{sid}` format, e.g.

```shell
terraform import twilio_autopilot_model_build.model_build /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/ModelBuilds/UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
