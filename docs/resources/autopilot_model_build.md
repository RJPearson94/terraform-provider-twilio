---
page_title: "Twilio Autopilot Model Build"
subcategory: "Autopilot"
---

# twilio_autopilot_model_build Resource

Manages an Autopilot model build. See the [API docs](https://www.twilio.com/docs/autopilot/api/model-build) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

~> To allow terraform to correctly manage the lifecycle of the model build you should use `create_before_destroy` argument in the lifecycle block. The docs can be found [here](https://www.terraform.io/docs/configuration/resources.html#create_before_destroy)

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

  triggers = {
    redeployment = sha1(join(",", list(
      twilio_autopilot_task_sample.task_sample.sid,
      twilio_autopilot_task_sample.task_sample.language,
      twilio_autopilot_task_sample.task_sample.tagged_text,
    )))
  }

  lifecycle {
    create_before_destroy = true
  }

  polling {
    enabled = true
  }
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to associate the model build with. Changing this forces a new resource to be created
- `unique_name_prefix` - (Optional) The prefix to be concatenated with an unique ID to form the unique name of the model build
- `status_callback` - (Optional) The callback URL to post build statuses to. Changing this forces a new resource to be created
- `triggers` - (Optional) A map of key-value pairs which can be used to determine if changes have occurred and a redeployment is necessary. Changing this forces a new resource to be created
~> An alternative strategy is to use the [taint](https://www.terraform.io/docs/commands/taint.html) functionality of terraform.
- `polling` - (Optional) A `polling` block as documented below.

---

A `polling` block supports the following:

- `enabled` - (Required) Enable or disable polling of the build.
- `max_attempts` - (Optional) The maximum number of polling attempts. Default is 24
- `delay_in_ms` - (Optional) The time in milliseconds to wait between polling attempts. Default is 5000ms

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the model build (Same as the SID)
- `sid` - The SID of the model build (Same as the ID)
- `account_sid` - The account SID associated with the model build
- `unique_name_prefix` - The prefix to be concatenated with an unique ID to form the unique name of the model build
- `unique_name` - The unique name of the model build
- `status_callback` - The callback URL to post build statuses to
- `status` - The current model build status
- `triggers` - A map of key-value pairs which can be used to determine if changes have occurred and a redeployment is necessary.
- `error_code` - The error code of the model build if the status is failed
- `build_duration` - The duration of the model build (in seconds)
- `date_created` - The date in RFC3339 format that the model build was created
- `date_updated` - The date in RFC3339 format that the model build was updated
- `url` - The URL of the model build resource

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

!> The following arguments "unique_name_prefix", "triggers" and "polling" cannot be imported
