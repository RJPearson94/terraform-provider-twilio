---
page_title: "twilio_sip_credential_list Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_credential_list Data Source

Use this data source to access information about an existing credential list. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource) for more information

## Example Usage

```hcl
data "twilio_sip_credential_list" "credential_list" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential_list" {
  value = data.twilio_sip_credential_list.credential_list
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP credential list
- `sid` (String) The SID of the SIP credential list to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP credential list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP credential list was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP credential list
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
