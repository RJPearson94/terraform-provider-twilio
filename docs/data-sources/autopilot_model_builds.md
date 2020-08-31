---
page_title: "Twilio Autopilot Model Builds"
subcategory: "Autopilot"
---

# twilio_autopilot_model_builds Data Source

Use this data source to access information about the model builds associated with an existing Autopilot assistant. See the [API docs](https://www.twilio.com/docs/autopilot/api/model-build) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_model_builds" "model_builds" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "model_builds" {
  value = data.twilio_autopilot_model_builds.model_builds
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the model builds are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the assistant SID)
- `account_sid` - The SID of the account the model builds are associated with
- `assistant_sid` - The SID of the assistant the model builds are associated with
- `model_builds` - A list of `model_build` blocks as documented below

---

A `model_build` block supports the following:

- `sid` - The SID of the model build
- `unique_name` - The unique name of the model build
- `status` - The current model build status
- `error_code` - The error code of the model build if the status is failed
- `build_duration` - The duration of the model build (in seconds)
- `date_created` - The date in RFC3339 format that the model build was created
- `date_updated` - The date in RFC3339 format that the model build was updated
- `url` - The URL of the model build resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving model builds
