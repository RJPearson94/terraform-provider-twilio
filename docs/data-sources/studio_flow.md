---
page_title: "Twilio Studio Flow"
subcategory: "Studio"
---

# twilio_studio_flow Data Source

Use this data source to access information about an existing studio flow. See the [API docs](https://www.twilio.com/docs/studio/rest-api/v2/flow) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

!> This API used to manage this resource is currently in beta and is subject to change

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

- `sid` - (Mandatory) The SID of the studio flow

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the studio flow (Same as the SID)
- `sid` - The SID of the studio flow (Same as the ID)
- `friendly_name` - The name of the studio flow
- `definition` - The Flow Definition JSON
- `status` - The status of the studio flow
- `revision` - The revision number of teh studio flow
- `valid` - Whether the studio flow is valid
- `date_created` - The date in RFC3339 format that the studio flow was created
- `date_updated` - The date in RFC3339 format that the studio flow was updated
- `url` - The url of the studio flow
- `webhook_url` - The webhook url of the studio flow

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the studio flow
