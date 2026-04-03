---
page_title: "twilio_sip_domain_credential_list_mapping Resource - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_domain_credential_list_mapping Resource

Manages a credential list mapping. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-credentiallistmapping-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_domain" "domain" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_name = "test.sip.twilio.com"
}

resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}

resource "twilio_sip_domain_credential_list_mapping" "credential_list_mapping" {
  account_sid         = twilio_sip_domain.domain.account_sid
  domain_sid          = twilio_sip_domain.domain.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP domain credential list mapping. Changing this forces a new resource
- `credential_list_sid` (String) The SID of the SIP credential list to map to the domain for call authentication. Changing this forces a new resource
- `domain_sid` (String) The SID of the SIP domain to map the credential list to. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP domain credential list mapping was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP domain credential list mapping was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP domain credential list mapping
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP domain credential list mapping by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

An domain can be imported using the `Accounts/{accountSid}/Domains/{domainSid}/Auth/Calls/CredentialListMappings/{sid}` format, e.g.

```shell
terraform import twilio_sip_domain_credential_list_mapping.credential_list_mapping /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Domains/DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
