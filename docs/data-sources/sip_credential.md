---
page_title: "twilio_sip_credential Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_credential Data Source

Use this data source to access information about an existing credential. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-credential-resource) for more information

## Example Usage

```hcl
data "twilio_sip_credential" "credential" {
  account_sid         = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  credential_list_sid = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid                 = "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential" {
  value = data.twilio_sip_credential.credential
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP credential
- `credential_list_sid` (String) The SID of the credential list that this SIP credential belongs to
- `sid` (String) The SID of the SIP credential to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP credential was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP credential was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `username` (String) The username for the SIP credential

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
