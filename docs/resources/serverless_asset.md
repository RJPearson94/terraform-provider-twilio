---
page_title: "Twilio Serverless Asset"
subcategory: "Serverless"
---

# twilio_serverless_asset Resource

Manages a Serverless asset

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the asset is managed under. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the asset

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the asset (Same as the SID)
- `sid` - The SID of the asset (Same as the ID)
- `account_sid` - The Account SID of the asset is deployed into
- `service_sid` - The Service SID of the asset is managed under
- `friendly_name` - The name of the asset
- `date_created` - The date that the asset was created
- `date_updated` - The date that the asset was updated
- `url` - The url of the asset
