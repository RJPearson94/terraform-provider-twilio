---
page_title: "twilio_sip_credential_list Resource - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_credential_list Resource

Manages a credential list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP credential list. Changing this forces a new resource
- `friendly_name` (String) A human-readable label for the SIP credential list

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP credential list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP credential list was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP credential list by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A credential list can be imported using the `/Accounts/{accountSid}/CredentialLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_credential_list.credential_list /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
