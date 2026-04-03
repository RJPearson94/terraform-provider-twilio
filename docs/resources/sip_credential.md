---
page_title: "twilio_sip_credential Resource - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_credential Resource

Manages a credential list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-credential-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}

resource "twilio_sip_credential" "credential" {
  account_sid         = twilio_sip_credential_list.credential_list.account_sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
  username            = "test"
  password            = "test"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP credential. Changing this forces a new resource
- `credential_list_sid` (String) The SID of the credential list that this SIP credential belongs to. Changing this forces a new resource
- `password` (String, Sensitive) The password for the SIP credential. Must be at least 12 characters and contain an uppercase letter, lowercase letter, and number. Sensitive -- will not be shown in logs or plans
- `username` (String) The username for the SIP credential (1 to 32 characters). Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP credential was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP credential was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP credential by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A credential list can be imported using the `/Accounts/{accountSid}/CredentialLists/{credentialListSid}/Credentials/{sid}` format, e.g.

```shell
terraform import twilio_sip_credential.credential /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
