---
page_title: "twilio_sip_trunking_origination_url Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_origination_url Data Source

Use this data source to access information about an existing origination URL. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/originationurl-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid       = "OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "origination_url" {
  value = data.twilio_sip_trunking_origination_url.origination_url
}
```

## Schema

### Required

- `sid` (String) The SID of the SIP trunk origination URL to look up
- `trunk_sid` (String) The SID of the SIP trunk the origination URL belongs to

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk origination URL
- `date_created` (String) The date and time the SIP trunk origination URL was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk origination URL was last updated, in RFC 3339 format
- `enabled` (Boolean) Whether the origination URL is enabled and available for use
- `friendly_name` (String) A human-readable label for the SIP trunk origination URL
- `id` (String) The ID of this resource.
- `priority` (Number) The priority of the origination URL. Lower values have higher priority
- `sip_url` (String) The SIP address to route origination calls to
- `url` (String) The absolute URL of the SIP trunk origination URL resource
- `weight` (Number) The weight of the origination URL, used for load balancing among URLs with the same priority

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
