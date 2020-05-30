# Data Source: twilio_studio_flow

Use this data source to access information about an existing Studio Flow

**NOTE:** This resource is in beta

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

* `sid` - The SID of the Studio Flow

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Studio Flow (Same as the SID)
* `sid` - The SID of the Studio Flow (Same as the ID)
* `friendly_name` - The name of the Studio Flow
* `definition` - The Flow Definition JSON
* `status` -  The status of the Studio Flow
* `revision` - The revision number of teh Studio Flow
* `valid` -  Whether the Studio Flow is valid
* `date_created` - The date that the Studio Flow was created
* `date_updated` - The date that the Studio Flow was updated
* `url` - The url of the Studio Flow
* `webhook_url` - The webhook url of the Studio Flow
