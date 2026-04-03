---
page_title: "twilio_sip_ip_address Resource - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_ip_address Resource

Manages an IP address. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}

resource "twilio_sip_ip_address" "ip_address" {
  account_sid                = twilio_sip_ip_access_control_list.ip_access_control_list.account_sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
  friendly_name              = "Test"
  ip_address                 = "127.0.0.1"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP IP address. Changing this forces a new resource
- `friendly_name` (String) A human-readable label for the SIP IP address
- `ip_access_control_list_sid` (String) The SID of the IP access control list that this IP address belongs to. Changing this forces a new resource
- `ip_address` (String) The IP address to allow or deny in the access control list

### Optional

- `cidr_length_prefix` (Number) The CIDR prefix length for the IP address range. Defaults to `32`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP IP address was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP IP address was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP IP address by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

An IP address can be imported using the `Accounts/{accountSid}/IpAccessControlLists/{ipAccessControlListSid}/IpAddresses/{sid}` format, e.g.

```shell
terraform import twilio_sip_ip_address.ip_address /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
