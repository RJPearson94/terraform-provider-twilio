---
page_title: "Twilio Autopilot Model Build"
subcategory: "Autopilot"
---

# twilio_autopilot_model_build Data Source

Use this data source to access information about an existing Autopilot model build. See the [API docs](https://www.twilio.com/docs/autopilot/api/model-build) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_model_build" "model_build" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid = "UGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "model_build" {
  value = data.twilio_autopilot_model_build.model_build
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the model build is associated with
- `sid` - (Mandatory) The SID of the model build

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the model build (Same as the SID)
- `sid` - The SID of the model build (Same as the ID)
- `account_sid` - The account SID associated with the model build
- `unique_name` - The unique name of the model build
- `status` - The current model build status
- `error_code` - The error code of the model build if the status is failed
- `build_duration` - The duration of the model build (in seconds)
- `date_created` - The date in RFC3339 format that the model build was created
- `date_updated` - The date in RFC3339 format that the model build was updated
- `url` - The url of the model build resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the model build
