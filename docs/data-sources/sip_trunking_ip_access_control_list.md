---
page_title: "Twilio SIP Trunking IP Access Control List"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_ip_access_control_list Data Source

Use this data source to access information about an existing IP access control list. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
data "twilio_sip_trunking_ip_access_control_list" "ip_access_control_list" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid       = "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_access_control_list" {
  value = data.twilio_sip_trunking_ip_access_control_list.ip_access_control_list
}
```

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the IP access control list is associated with
- `sid` - (Mandatory) The SID of the IP access control list

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the IP access control list (Same as the `sid`)
- `sid` - The SID of the IP access control list (Same as the `id`)
- `account_sid` - The account SID associated with the IP access control list
- `trunk_sid` - The trunk SID associated with the IP access control list
- `friendly_name` - The friendly name of the IP access control list
- `date_created` - The date in RFC3339 format that the IP access control list was created
- `date_updated` - The date in RFC3339 format that the IP access control list was updated
- `url` - The URL of the IP access control list resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the IP access control list details
