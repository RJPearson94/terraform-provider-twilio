---
page_title: "twilio_sip_trunking_ip_access_control_list Resource - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_ip_access_control_list Resource

Manages an IP access control list. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}

resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_ip_access_control_list" "ip_access_control_list" {
  trunk_sid                  = twilio_sip_trunking_trunk.trunk.sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
}
```

## Schema

### Required

- `ip_access_control_list_sid` (String) The SID of the IP access control list to associate with the trunk. Changing this forces a new resource
- `trunk_sid` (String) The SID of the SIP trunk to associate the IP access control list with. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk IP access control list
- `date_created` (String) The date and time the SIP trunk IP access control list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk IP access control list was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP trunk IP access control list
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP trunk IP access control list by Twilio
- `url` (String) The absolute URL of the SIP trunk IP access control list resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

An IP access control list can be imported using the `/Trunks/{trunkSid}/IpAccessControlLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_ip_access_control_list.ip_access_control_list /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
