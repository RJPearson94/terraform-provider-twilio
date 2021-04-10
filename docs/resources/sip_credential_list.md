---
page_title: "Twilio SIP Credential List"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the credential list with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the credential list

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential list (Same as the `sid` )
- `sid` - The SID of the credential list (Same as the `id`)
- `account_sid` - The account SID associated with the credential list
- `friendly_name` - The friendly name of the credential list
- `date_created` - The date in RFC3339 format that the credential list was created
- `date_updated` - The date in RFC3339 format that the credential list was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the credential list
- `update` - (Defaults to 10 minutes) Used when updating the credential list
- `read` - (Defaults to 5 minutes) Used when retrieving the credential list
- `delete` - (Defaults to 10 minutes) Used when deleting the credential list

## Import

A credential list can be imported using the `/Accounts/{accountSid}/CredentialLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_credential_list.credential_list /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
