# twilio_proxy_phone_number

Manages a Proxy Phone Number

## Example Usage

```hcl
resource "twilio_proxy_service" "service" {
  unique_name = "Test Proxy Service"
}

resource "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_service.service.sid
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  is_reserved = true
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of a Twilio Proxy Service. Changing this forces a new resource to be created
- `sid` - (Optional) The SID of a Twilio Phone Number to associate with the proxy
- `is_reserved` - (Optional) Whether the phone number is reserved

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the proxy phone number resource (Same as the SID)
- `sid` - The SID of a Twilio Phone Number to associate with the proxy (Same as the ID)
- `account_sid` - The Account SID of the phone number resource is deployed into
- `service_sid` - The SID of a Twilio Proxy Service
- `is_reserved` - Whether the phone number is reserved
- `phone_number` - The phone number associated with the SID
- `friendly_name` - The friendly name associated with the SID
- `iso_country` - The ISO country associated with the SID
- `in_use` - The number of active calls associated with the SID
- `date_created` - The date that the proxy phone number resource was created
- `date_updated` - The date that the proxy phone number resource was updated
- `url` - The url of the proxy phone number
