---
page_title: "Twilio SIP IP Addresses"
subcategory: "SIP"
---

# twilio_sip_ip_addresses Data Source

Use this data source to access information about the IP Addresses associated with an existing account and IP access control list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaddress-resource) for more information

## Example Usage

```hcl
data "twilio_sip_ip_addresses" "ip_addresses" {
  account_sid                = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  ip_access_control_list_sid = "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_addresses" {
  value = data.twilio_sip_ip_addresses.ip_addresses
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the IP addresses are associated with
- `ip_access_control_list_sid` - (Mandatory) The SID of the IP access control list the IP addresses are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `account_sid/ip_access_control_list_sid`
- `account_sid` - The SID of the account the IP addresses are associated with
- `ip_access_control_list_sid` - The SID of the credential list the IP addresses are associated with
- `ip_addresses` - A list of `ip_address` blocks as documented below

---

An `ip_address` block supports the following:

- `sid` - The SID of the IP address
- `friendly_name` - The friendly name of the IP address
- `ip_address` - The IP address of the resource
- `cidr_length_prefix` - The CIDR length prefix for the IP address
- `date_created` - The date in RFC3339 format that the IP address was created
- `date_updated` - The date in RFC3339 format that the IP address was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the IP addresses
