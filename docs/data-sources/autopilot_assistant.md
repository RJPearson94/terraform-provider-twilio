---
page_title: "Twilio Autopilot Assistant"
subcategory: "Autopilot"
---

# twilio_autopilot_assistant Data Source

Use this data source to access information about an existing Autopilot assistant. See the [API docs](https://www.twilio.com/docs/autopilot/api/assistant) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_assistant" "assistant" {
  sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "assistant" {
  value = data.twilio_autopilot_assistant.assistant
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the assistant

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the assistant (Same as the `sid`)
- `sid` - The SID of the assistant (Same as the `id`)
- `account_sid` - The account SID associated with the assistant
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
- `url` - The URL of the assistant

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the assistant
