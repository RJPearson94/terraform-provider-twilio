---
page_title: "twilio_sip_domain_registration_credential_list_mappings Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_domain_registration_credential_list_mappings Data Source

Use this data source to access information about existing SIP domain registration credential list mappings associated with an existing account and domain. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-domain-registration-credentiallistmapping-resource) for more information

## Example Usage

```hcl
data "twilio_sip_domain_registration_credential_list_mappings" "credential_list_mappings" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_sid  = "DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential_list_mappings" {
  value = data.twilio_sip_domain_registration_credential_list_mappings.credential_list_mappings
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns the SIP domain registration credential list mappings
- `domain_sid` (String) The SID of the SIP domain to retrieve registration credential list mappings for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `credential_list_mappings` (List of Object) A list of registration credential list mappings for the SIP domain (see [below for nested schema](#nestedatt--credential_list_mappings))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--credential_list_mappings"></a>
### Nested Schema for `credential_list_mappings`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
