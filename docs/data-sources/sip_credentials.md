---
page_title: "Twilio SIP Credentials"
subcategory: "SIP"
---

# twilio_sip_credentials Data Source

Use this data source to access information about the credentials associated with an existing account and credential list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-credential-resource) for more information

## Example Usage

```hcl
data "twilio_sip_credentials" "credentials" {
  account_sid         = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  credential_list_sid = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credentials" {
  value = data.twilio_sip_credentials.credentials
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the credentials are associated with
- `credential_list_sid` - (Mandatory) The SID of the credential list the credentials are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `account_sid/credential_list_sid`
- `account_sid` - The SID of the account the credentials are associated with
- `credential_list_sid` - The SID of the credential list the credentials are associated with
- `credentials` - A list of `credential` blocks as documented below

---

A `credential` block supports the following:

- `sid` - The SID of the credential
- `username` - The credential username
- `date_created` - The date in RFC3339 format that the credential was created
- `date_updated` - The date in RFC3339 format that the credential was updated

!> For security reasons, the API does not return the password so cannot be returned in the data lookup

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the credentials
