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
- `date_created` - The date in RFC3339 format that the asset was created
- `date_updated` - The date in RFC3339 format that the asset was updated
- `url` - The url of the asset

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the asset
- `update` - (Defaults to 10 minutes) Used when updating the asset
- `read` - (Defaults to 5 minutes) Used when retrieving the asset
- `delete` - (Defaults to 10 minutes) Used when deleting the asset

## Import

A asset can be imported using the `/Services/{serviceSid}/Assets/{sid}` format, e.g.

```shell
terraform import twilio_serverless_asset.asset /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
