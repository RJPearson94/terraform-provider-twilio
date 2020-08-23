---
page_title: "Twilio Messaging Phone Number"
subcategory: "Messaging"
---

# twilio_messaging_phone_number Resource

Manages a messaging phone number resource

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_messaging_phone_number" "phone_number" {
  service_sid = twilio_messaging_service.service.sid
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The messaging service SID associated with the phone number. Changing this forces a new resource to be created
- `sid` - (Mandatory) The SID of the Twilio phone number to associated with the messaging service. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the phone number (Same as the SID)
- `sid` - The SID of the Twilio phone number to associated with the messaging service (Same as the ID)
- `service_sid` - The messaging service SID associated with the phone number
- `account_sid` - The account SID associated with the phone number
- `capabilities` - The capabilities that are enabled for the phone number
- `country_code` - The country code of the phone number
- `phone_number` - The phone number
- `date_created` - The date in RFC3339 format that the messaging phone number resource was created
- `date_updated` - The date in RFC3339 format that the messaging phone number resource was updated
- `url` - The url of the messaging phone number resource
