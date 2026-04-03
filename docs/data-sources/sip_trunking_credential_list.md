---
page_title: "twilio_sip_trunking_credential_list Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_credential_list Data Source

Use this data source to access information about an existing credential list. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
data "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid       = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential_list" {
  value = data.twilio_sip_trunking_credential_list.credential_list
}
```

## Schema

### Required

- `sid` (String) The SID of the SIP trunk credential list to look up
- `trunk_sid` (String) The SID of the SIP trunk the credential list belongs to

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk credential list
- `date_created` (String) The date and time the SIP trunk credential list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk credential list was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP trunk credential list
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the SIP trunk credential list resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
