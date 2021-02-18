---
page_title: "Twilio SIP Credential"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the credential is associated with
- `credential_list_sid` - (Mandatory) The SID of the credential list the credential is associated with
- `sid` - (Mandatory) The SID of the credential

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential (Same as the `sid` )
- `sid` - The SID of the credential (Same as the `id`)
- `account_sid` - The account SID associated with the credential
- `credential_list_sid` - The credential list SID associated with the credential
- `username` - The credential username
- `date_created` - The date in RFC3339 format that the credential was created
- `date_updated` - The date in RFC3339 format that the credential was updated

!> For security reasons, the API does not return the password so cannot be returned in the data lookup

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the credential details
