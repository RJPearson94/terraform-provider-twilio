---
page_title: "Twilio SIP Domain Registration Credential List Mapping"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the credential list mapping is associated with
- `domain_sid` - (Mandatory) The SID of the domain the credential list mapping is associated with
- `sid` - (Mandatory) The SID of the credential list mapping

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential list mapping (Same as the `sid`)
- `sid` - The SID of the credential list mapping (Same as the `id`)
- `account_sid` - The account SID associated with the credential list mapping
- `domain_sid` - The domain SID associated with the credential list mapping
- `friendly_name` - The friendly name of the credential list mapping
- `date_created` - The date in RFC3339 format that the credential list mapping was created
- `date_updated` - The date in RFC3339 format that the credential list mapping was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the credential list mapping details
