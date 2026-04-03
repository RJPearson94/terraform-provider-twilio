---
page_title: "twilio_sip_ip_access_control_list Resource - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_ip_access_control_list Resource

Manages an IP access control list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP IP access control list. Changing this forces a new resource
- `friendly_name` (String) A human-readable label for the SIP IP access control list

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP IP access control list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP IP access control list was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP IP access control list by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

An IP access control list can be imported using the `Accounts/{accountSid}/IpAccessControlLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_ip_access_control_list.ip_access_control_list /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
