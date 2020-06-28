---
page_title: "Twilio Serverless Service"
subcategory: "Serverless"
---

# twilio_serverless_service Resource

Manages a Serverless service

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `unique_name` - (Mandatory) The unique name of the service
- `friendly_name` - (Mandatory) The name of the service
- `include_credentials` - (Optional) Whether or not credentials are included in the service runtime
- `ui_editable` - (Optional) Whether or not the service is editable in the console

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the SID)
- `sid` - The SID of the service (Same as the ID)
- `account_sid` - The Account SID of the service is deployed into
- `unique_name` - The unique name of the service
- `friendly_name` - The name of the service
- `include_credentials` - Whether or not credentials are included in the service runtime
- `ui_editable` - Whether or not the service is editable in the console
- `date_created` - The date that the service was created
- `date_updated` - The date that the service was updated
- `url` - The url of the service
