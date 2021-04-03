---
page_title: "Twilio SIP Credential"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the credential with. Changing this forces a new resource to be created
- `credential_list_sid` - (Mandatory) The credential list SID to associate the credential with. Changing this forces a new resource to be created
- `username` - (Mandatory) The credential username. Changing this forces a new resource to be created. The length of the string must be between `1` and `64` characters (inclusive)
- `password` - (Mandatory) The credential password. The length of the string must be between at least `12` characters and contain at least 1 `uppercase character`, 1 `lowercase character` and 1 `number`.

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential (Same as the `sid` )
- `sid` - The SID of the credential (Same as the `id`)
- `account_sid` - The account SID associated with the credential
- `credential_list_sid` - The credential list SID associated with the credential
- `username` - The credential username
- `date_created` - The date in RFC3339 format that the credential was created
- `date_updated` - The date in RFC3339 format that the credential was updated

!> For security reasons, the API does not return the password. So if the password is changed outside of Terraform then the drift may not be detected

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the credential
- `update` - (Defaults to 10 minutes) Used when updating the credential
- `read` - (Defaults to 5 minutes) Used when retrieving the credential
- `delete` - (Defaults to 10 minutes) Used when deleting the credential

## Import

A credential list can be imported using the `/Accounts/{accountSid}/CredentialLists/{credentialListSid}/Credentials/{sid}` format, e.g.

```shell
terraform import twilio_sip_credential.credential /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
