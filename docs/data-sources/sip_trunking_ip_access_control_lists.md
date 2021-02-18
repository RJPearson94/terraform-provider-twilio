---
page_title: "Twilio SIP Trunking IP Access Control Lists"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_ip_access_control_lists Data Source

Use this data source to access information about the IP access control lists associated with an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_ip_access_control_lists" "ip_access_control_lists" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_access_control_lists" {
  value = data.twilio_sip_trunking_ip_access_control_lists.ip_access_control_lists
}
```

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the IP access control lists are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `trunk_sid`)
- `trunk_sid` - The SID of the SIP trunk the IP access control lists are associated with (Same as `id`)
- `account_sid` - The account SID associated with the IP access control lists
- `ip_access_control_lists` - A list of `ip_access_control_list` blocks as documented below

---

An `ip_access_control_list` block supports the following:

- `sid` - The SID of the IP access control list
- `friendly_name` - The friendly name of the IP access control list
- `date_created` - The date in RFC3339 format that the IP access control list was created
- `date_updated` - The date in RFC3339 format that the IP access control list was updated
- `url` - The URL of the IP access control list resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving IP access control lists
