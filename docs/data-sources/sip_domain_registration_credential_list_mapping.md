---
page_title: "twilio_sip_domain_registration_credential_list_mapping Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_domain_registration_credential_list_mapping Data Source

Use this data source to access information about an existing SIP domain registration credential list mapping. See the [API docs](ttps://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource) for more information

## Example Usage

```hcl
data "twilio_sip_domain_registration_credential_list_mapping" "credential_list_mapping" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_sid  = "DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential_list_mapping" {
  value = data.twilio_sip_domain_registration_credential_list_mapping.credential_list_mapping
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP domain registration credential list mapping
- `domain_sid` (String) The SID of the SIP domain the credential list is mapped to for registration
- `sid` (String) The SID of the SIP domain registration credential list mapping to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP domain registration credential list mapping was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP domain registration credential list mapping was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP domain registration credential list mapping
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
