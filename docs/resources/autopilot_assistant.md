---
page_title: "Twilio Autopilot Assistant"
subcategory: "Autopilot"
---

# twilio_autopilot_assistant Resource

Manages a Autopilot assistant

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "twilio-test"
  defaults      = <<EOF
{
 "defaults": {
  "assistant_initiation": "",
  "fallback": "http://localhost/fallback"
 }
}
EOF

  stylesheet    = <<EOF
{
 "style_sheet": {
  "voice": {
   "say_voice": "Polly.Matthew"
  }
 }
}
EOF
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Optional) The friendly name of the assistant
- `unique_name` - (Optional) The unique name of the assistant
- `log_queries` - (Optional) Whether or not queries are recorded/ logged
- `development_stage` - (Optional) The stage description for the assistant. Valid values are `in-development` or `in-production`.
- `callback_url` - (Optional) The URL the assistant will call back to when an event is fired
- `callback_events` - (Optional) A list of callback events strings which trigger the callback webhook
- `defaults` - (Optional) JSON string of an Autopilot defaults
- `stylesheet` - (Optional) JSON string of an Autopilot stylesheet

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the assistant (Same as the SID)
- `sid` - The SID of the assistant (Same as the ID)
- `account_sid` - The Account SID associated with the assistant
- `friendly_name` - The friendly name of the assistant
- `unique_name` - The unique name of the assistant
- `latest_model_build_sid` - The latest model build SID of the assistant
- `needs_model_build` - Whether or not a model build is required for the assistant
- `log_queries` - Whether or not queries are recorded/ logged
- `development_stage` - The stage description for the assistant
- `callback_url` - The URL the assistant will call back to when an event is fired
- `callback_events` - A list of callback events strings which trigger the callback webhook
- `defaults` - JSON string of an Autopilot defaults
- `stylesheet` - JSON string of an Autopilot stylesheet
- `date_created` - The date in RFC3339 format that the assistant was created
- `date_updated` - The date in RFC3339 format that the assistant was updated
- `url` - The url of the assistant

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the assistant
- `update` - (Defaults to 10 minutes) Used when updating the assistant
- `read` - (Defaults to 5 minutes) Used when retrieving the assistant
- `delete` - (Defaults to 10 minutes) Used when deleting the assistant

## Import

A assistant can be imported using the `/Assistants/{sid}` format, e.g.

```shell
terraform import twilio_autopilot_assistant.assistant /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
