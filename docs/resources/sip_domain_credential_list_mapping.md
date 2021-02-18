---
page_title: "Twilio SIP Domain Credential List Mapping"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the domain with. Changing this forces a new resource to be created
- `domain_sid` - (Mandatory) The domain SID to associate the credential list mapping with. Changing this forces a new resource to be created
- `credential_list_sid` - (Mandatory) The credential list SID to associate the credential list mapping with. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential list mapping (Same as the `sid` & `credential_list_sid`)
- `sid` - The SID of the credential list mapping (Same as the `id` & `credential_list_sid`)
- `account_sid` - The account SID associated with the credential list mapping
- `domain_sid` - The domain SID associated with the credential list mapping
- `credential_list_sid` - The credential list SID associated with the credential list mapping
- `friendly_name` - The friendly name of the credential list mapping
- `date_created` - The date in RFC3339 format that the credential list mapping was created
- `date_updated` - The date in RFC3339 format that the credential list mapping was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the credential list mapping
- `read` - (Defaults to 5 minutes) Used when retrieving the credential list mapping
- `delete` - (Defaults to 10 minutes) Used when deleting the credential list mapping

## Import

An domain can be imported using the `Accounts/{accountSid}/Domains/{domainSid}/Auth/Calls/CredentialListMappings/{sid}` format, e.g.

```shell
terraform import twilio_sip_domain_credential_list_mapping.credential_list_mapping /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Domains/DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/CredentialListMappings/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
