---
page_title: "Twilio Serverless Function"
subcategory: "Serverless"
---

# twilio_serverless_function Resource

Manages a Serverless function

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_function" "function" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the function is managed under. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the function

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the function (Same as the SID)
- `sid` - The SID of the function (Same as the ID)
- `account_sid` - The Account SID of the function is deployed into
- `service_sid` - The Service SID of the function is managed under
- `friendly_name` - The name of the function
- `date_created` - The date that the function was created
- `date_updated` - The date that the function was updated
- `url` - The url of the function
