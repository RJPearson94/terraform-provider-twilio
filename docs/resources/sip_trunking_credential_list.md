---
page_title: "twilio_sip_trunking_credential_list Resource - twilio"
subcategory: "SIP Trunking"
description: |-
  
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

## Schema

### Required

- `credential_list_sid` (String) The SID of the SIP credential list to associate with the trunk. Changing this forces a new resource
- `trunk_sid` (String) The SID of the SIP trunk to associate the credential list with. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk credential list
- `date_created` (String) The date and time the SIP trunk credential list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk credential list was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP trunk credential list
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP trunk credential list by Twilio
- `url` (String) The absolute URL of the SIP trunk credential list resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A credential list can be imported using the `/Trunks/{trunkSid}/CredentialLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_credential_list.credential_list /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/CredentialLists/CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
