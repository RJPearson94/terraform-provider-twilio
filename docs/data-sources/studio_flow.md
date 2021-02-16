---
page_title: "Twilio Studio Flow"
subcategory: "Studio"
---

# twilio_studio_flow Data Source

Use this data source to access information about an existing studio flow. See the [API docs](https://www.twilio.com/docs/studio/rest-api/v2/flow) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

```hcl
data "twilio_studio_flow" "flow" {
  sid = "FWxxxxxxxxxxxxxxxx"
}

output "definition" {
  value = data.twilio_studio_flow.flow.definition
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the Studio flow

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the Studio flow (Same as the SID)
- `sid` - The SID of the Studio flow (Same as the ID)
- `friendly_name` - The name of the Studio flow
- `definition` - The flow definition JSON
- `status` - The status of the Studio flow
- `revision` - The revision number of the Studio flow
- `valid` - Whether the Studio flow is valid
- `date_created` - The date in RFC3339 format that the Studio flow was created
- `date_updated` - The date in RFC3339 format that the Studio flow was updated
- `url` - The URL of the Studio flow
- `webhook_url` - The webhook URL of the Studio flow

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the Studio flow
