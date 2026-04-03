---
page_title: "twilio_proxy_phone_number Data Source - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_phone_number Data Source

Use this data source to access information about an existing Proxy phone number. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_phone_number" "phone_number" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_number" {
  value = data.twilio_proxy_phone_number.phone_number
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Proxy service
- `sid` (String) The SID of the Proxy phone number to read

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Proxy phone number
- `capabilities` (List of Object) The capabilities of the Proxy phone number (see [below for nested schema](#nestedatt--capabilities))
- `date_created` (String) The date and time the Proxy phone number was created, in RFC 3339 format
- `date_updated` (String) The date and time the Proxy phone number was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the Proxy phone number
- `id` (String) The ID of this resource.
- `in_use` (Number) The number of active Proxy sessions assigned to this phone number
- `is_reserved` (Boolean) Whether the phone number is reserved and not assigned to a Proxy session
- `iso_country` (String) The ISO country code of the phone number
- `phone_number` (String) The phone number in E.164 format
- `url` (String) The absolute URL of the Proxy phone number resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


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
