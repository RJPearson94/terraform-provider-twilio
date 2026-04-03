---
page_title: "twilio_messaging_phone_number Resource - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_phone_number Resource

Manages a Programmable Messaging phone number resource. See the [API docs](https://www.twilio.com/docs/sms/services/api/phonenumber-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

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

## Schema

### Required

- `service_sid` (String) The SID of the messaging service to associate the phone number with. Changing this forces a new resource
- `sid` (String) The SID of the phone number to associate with the messaging service. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this phone number
- `capabilities` (List of String) The list of capabilities for the phone number
- `country_code` (String) The two-character ISO country code of the phone number
- `date_created` (String) The date and time the phone number was created, in RFC 3339 format
- `date_updated` (String) The date and time the phone number was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `phone_number` (String) The phone number in E.164 format
- `url` (String) The absolute URL of the phone number resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A phone number can be imported using the `/Services/{serviceSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_messaging_phone_number.phone_number /Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
