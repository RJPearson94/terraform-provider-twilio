---
page_title: "Twilio SIP IP Address"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the IP address with. Changing this forces a new resource to be created
- `ip_access_control_list_sid` - (Mandatory) The IP access control list SID to associate the IP address with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the IP address
- `ip_address` - (Mandatory) The IP address of the resource
- `cidr_length_prefix` - (Optional) The CIDR length prefix to use with the IP address. The default value ia `32`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the IP address (Same as the `sid`)
- `sid` - The SID of the IP address (Same as the `id`)
- `account_sid` - The account SID associated with the IP address
- `ip_access_control_list_sid` - The IP access control list SID associated with the IP address
- `friendly_name` - The friendly name of the IP address
- `ip_address` - The IP address of the resource
- `cidr_length_prefix` - The CIDR length prefix for the IP address
- `date_created` - The date in RFC3339 format that the IP address was created
- `date_updated` - The date in RFC3339 format that the IP address was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the IP address
- `update` - (Defaults to 10 minutes) Used when updating the IP address
- `read` - (Defaults to 5 minutes) Used when retrieving the IP address
- `delete` - (Defaults to 10 minutes) Used when deleting the IP address

## Import

An IP address can be imported using the `Accounts/{accountSid}/IpAccessControlLists/{ipAccessControlListSid}/IpAddresses/{sid}` format, e.g.

```shell
terraform import twilio_sip_ip_address.ip_address /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAddresses/IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
