---
page_title: "Twilio SIP IP Address"
subcategory: "SIP"
---

# twilio_sip_ip_address Data Source

Use this data source to access information about an existing IP Address. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource) for more information

## Example Usage

```hcl
data "twilio_sip_ip_address" "ip_address" {
  account_sid                = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  ip_access_control_list_sid = "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid                        = "IPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_address" {
  value = data.twilio_sip_ip_address.ip_address
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the IP address is associated with
- `ip_access_control_list_sid` - (Mandatory) The SID of the IP access control list the IP address is associated with
- `sid` - (Mandatory) The SID of the IP address

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

- `read` - (Defaults to 5 minutes) Used when retrieving the IP address details
