---
page_title: "Twilio SIP Trunking Credential List"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_credential_list Resource

Manages a credential list. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}

resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_credential_list" "credential_list" {
  trunk_sid           = twilio_sip_trunking_trunk.trunk.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}
```

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The trunk SID to associate the credential list with. Changing this forces a new resource to be created
- `credential_list_sid` - (Mandatory) The SIP credential list SID to associate the credential list with. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential list (Same as the `sid` & `credential_list_sid`)
- `sid` - The SID of the credential list (Same as the `id` & `credential_list_sid`)
- `account_sid` - The account SID associated with the credential list
- `trunk_sid` - The trunk SID associated with the credential list
- `credential_list_sid` - The SIP credential list SID associated with the credential list
- `friendly_name` - The friendly name of the credential list
- `date_created` - The date in RFC3339 format that the credential list was created
- `date_updated` - The date in RFC3339 format that the credential list was updated
- `url` - The URL of the credential list resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the credential list
- `read` - (Defaults to 5 minutes) Used when retrieving the credential list
- `delete` - (Defaults to 10 minutes) Used when deleting the credential list

## Import

A credential list can be imported using the `/Trunks/{trunkSid}/CredentialLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_credential_list.credential_list /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
