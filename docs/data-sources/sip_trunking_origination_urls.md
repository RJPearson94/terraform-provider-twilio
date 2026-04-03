---
page_title: "twilio_sip_trunking_origination_urls Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_origination_urls Data Source

Use this data source to access information about the origination URLs associated with an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/originationurl-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_origination_urls" "origination_urls" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "origination_urls" {
  value = data.twilio_sip_trunking_origination_urls.origination_urls
}
```

## Schema

### Required

- `trunk_sid` (String) The SID of the SIP trunk to retrieve origination URLs for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the SIP trunk origination URLs
- `id` (String) The ID of this resource.
- `origination_urls` (List of Object) A list of origination URLs associated with the SIP trunk (see [below for nested schema](#nestedatt--origination_urls))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--origination_urls"></a>
### Nested Schema for `origination_urls`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `enabled` (Boolean)
- `friendly_name` (String)
- `priority` (Number)
- `sid` (String)
- `sip_url` (String)
- `url` (String)
- `weight` (Number)
