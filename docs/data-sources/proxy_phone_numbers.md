---
page_title: "twilio_proxy_phone_numbers Data Source - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_phone_numbers Data Source

Use this data source to access information about the phone numbers associated with an existing Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_phone_numbers" "phone_numbers" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_numbers" {
  value = data.twilio_proxy_phone_numbers.phone_numbers
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Proxy service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Proxy service
- `id` (String) The ID of this resource.
- `phone_numbers` (List of Object) A list of phone numbers associated with the Proxy service (see [below for nested schema](#nestedatt--phone_numbers))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--phone_numbers"></a>
### Nested Schema for `phone_numbers`

Read-Only:

- `capabilities` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--capabilities))
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `in_use` (Number)
- `is_reserved` (Boolean)
- `iso_country` (String)
- `phone_number` (String)
- `sid` (String)
- `url` (String)

<a id="nestedobjatt--phone_numbers--capabilities"></a>
### Nested Schema for `phone_numbers.capabilities`

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
