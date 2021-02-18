---
page_title: "Twilio SIP Trunking IP Access Control List"
subcategory: "SIP Trunking"
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

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The trunk SID to associate the IP access control list with. Changing this forces a new resource to be created
- `ip_access_control_list_sid` - (Mandatory) The SIP IP access control list SID to associate the IP access control list with. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the IP access control list (Same as the `sid` & `ip_access_control_list_sid`)
- `sid` - The SID of the IP access control list (Same as the `id` & `ip_access_control_list_sid`)
- `account_sid` - The account SID associated with the IP access control list
- `trunk_sid` - The trunk SID associated with the IP access control list
- `ip_access_control_list_sid` - The SIP IP access control list SID associated with the IP access control
- `friendly_name` - The friendly name of the IP access control list
- `date_created` - The date in RFC3339 format that the IP access control list was created
- `date_updated` - The date in RFC3339 format that the IP access control list was updated
- `url` - The URL of the IP access control list resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the IP access control list
- `read` - (Defaults to 5 minutes) Used when retrieving the IP access control list
- `delete` - (Defaults to 10 minutes) Used when deleting the IP access control list

## Import

An IP access control list can be imported using the `/Trunks/{trunkSid}/IpAccessControlLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_ip_access_control_list.ip_access_control_list /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
