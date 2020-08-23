---
page_title: "Twilio Messaging Alpha Sender"
subcategory: "Messaging"
---

# twilio_messaging_alpha_sender Resource

Manages a messaging alphanumeric sender resource

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_messaging_alpha_sender" "alpha_sender" {
  service_sid  = twilio_messaging_service.service.sid
  alpha_sender = "test"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The messaging service SID associated with the alpha sender. Changing this forces a new resource to be created
- `alpha_sender` - (Mandatory) The alpha sender name to associated with the messaging service. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the alpha sender resource (Same as the SID)
- `sid` - The SID of the alpha sender resource (Same as the ID)
- `service_sid` - The messaging service SID associated with the alpha sender
- `account_sid` - The account SID associated with the alpha sender
- `capabilities` - The capabilities that are enabled for the alpha sender
- `alpha_sender` - The alpha sender name to associated with the messaging service
- `date_created` - The date in RFC3339 format that the messaging alpha sender resource was created
- `date_updated` - The date in RFC3339 format that the messaging alpha sender resource was updated
- `url` - The url of the messaging alpha sender resource
