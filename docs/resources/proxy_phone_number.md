---
page_title: "twilio_proxy_phone_number Resource - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_phone_number Resource

Manages a Proxy phone number. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

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

## Schema

### Required

- `service_sid` (String) The SID of the Proxy service. Changing this forces a new resource

### Optional

- `is_reserved` (Boolean) Whether the phone number is reserved and not assigned to a Proxy session
- `phone_number` (String) The phone number in E.164 format to add to the Proxy service. Conflicts with `sid`. Changing this forces a new resource
- `sid` (String) The SID of the Twilio phone number to add to the Proxy service. Conflicts with `phone_number`. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Proxy phone number
- `capabilities` (List of Object) The capabilities of the Proxy phone number (see [below for nested schema](#nestedatt--capabilities))
- `date_created` (String) The date and time the Proxy phone number was created, in RFC 3339 format
- `date_updated` (String) The date and time the Proxy phone number was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the Proxy phone number
- `id` (String) The ID of this resource.
- `in_use` (Number) The number of active Proxy sessions assigned to this phone number
- `iso_country` (String) The ISO country code of the phone number
- `url` (String) The absolute URL of the Proxy phone number resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--capabilities"></a>
### Nested Schema for `capabilities`

Read-Only:

- `fax_inbound` (Boolean)
- `fax_outbound` (Boolean)
- `mms_inbound` (Boolean)
- `mms_outbound` (Boolean)
- `restriction_fax_domestic` (Boolean)
- `restriction_mms_domestic` (Boolean)
- `restriction_sms_domestic` (Boolean)
- `restriction_voice_domestic` (Boolean)
- `sip_trunking` (Boolean)
- `sms_inbound` (Boolean)
- `sms_outbound` (Boolean)
- `voice_inbound` (Boolean)
- `voice_outbound` (Boolean)

## Import

A phone number can be imported using the `/Services/{serviceSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_proxy_phone_number.phone_number /Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
