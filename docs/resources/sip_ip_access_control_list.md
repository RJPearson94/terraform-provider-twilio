---
page_title: "Twilio SIP IP Access Control List"
subcategory: "SIP"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the IP access control list with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the IP access control list. The value cannot be an empty string

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the IP access control list (Same as the `sid`)
- `sid` - The SID of the IP access control list (Same as the `id`)
- `account_sid` - The account SID associated with the IP access control list
- `friendly_name` - The friendly name of the IP access control list
- `date_created` - The date in RFC3339 format that the IP access control list was created
- `date_updated` - The date in RFC3339 format that the IP access control list was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the IP access control list
- `update` - (Defaults to 10 minutes) Used when updating the IP access control list
- `read` - (Defaults to 5 minutes) Used when retrieving the IP access control list
- `delete` - (Defaults to 10 minutes) Used when deleting the IP access control list

## Import

An IP access control list can be imported using the `Accounts/{accountSid}/IpAccessControlLists/{sid}` format, e.g.

```shell
terraform import twilio_sip_ip_access_control_list.ip_access_control_list /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/IpAccessControlLists/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
