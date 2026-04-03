---
page_title: "twilio_proxy_short_code Data Source - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_short_code Data Source

Use this data source to access information about an existing Proxy short code resource. See the [API docs](https://www.twilio.com/docs/proxy/api/short-code) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_short_code" "short_code" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_code" {
  value = data.twilio_proxy_short_code.short_code
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Proxy service
- `sid` (String) The SID of the Proxy short code to read

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Proxy short code
- `capabilities` (List of Object) The capabilities of the Proxy short code (see [below for nested schema](#nestedatt--capabilities))
- `date_created` (String) The date and time the Proxy short code was created, in RFC 3339 format
- `date_updated` (String) The date and time the Proxy short code was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `is_reserved` (Boolean) Whether the short code is reserved and not assigned to a Proxy session
- `iso_country` (String) The ISO country code of the short code
- `short_code` (String) The short code value
- `url` (String) The absolute URL of the Proxy short code resource

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
