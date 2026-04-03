---
page_title: "twilio_proxy_short_codes Data Source - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_short_codes Data Source

Use this data source to access information about the short codes associated with an existing Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/short-code) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_short_codes" "short_codes" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_codes" {
  value = data.twilio_proxy_short_codes.short_codes
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
- `short_codes` (List of Object) A list of short codes associated with the Proxy service (see [below for nested schema](#nestedatt--short_codes))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--short_codes"></a>
### Nested Schema for `short_codes`

Read-Only:

- `capabilities` (List of Object) (see [below for nested schema](#nestedobjatt--short_codes--capabilities))
- `date_created` (String)
- `date_updated` (String)
- `is_reserved` (Boolean)
- `iso_country` (String)
- `short_code` (String)
- `sid` (String)
- `url` (String)

<a id="nestedobjatt--short_codes--capabilities"></a>
### Nested Schema for `short_codes.capabilities`

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
