---
page_title: "Twilio SIP Domain IP Access Control List Mapping"
subcategory: "SIP"
---

# twilio_sip_domain_ip_access_control_list_mapping Data Source

Use this data source to access information about an existing IP access control list mapping. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource) for more information

## Example Usage

```hcl
data "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_sid  = "DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_access_control_list_mapping" {
  value = data.twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the IP access control list mapping is associated with
- `domain_sid` - (Mandatory) The SID of the domain the IP access control list mapping is associated with
- `sid` - (Mandatory) The SID of the IP access control list mapping

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the IP access control list mapping (Same as the `sid`)
- `sid` - The SID of the IP access control list mapping (Same as the `id`)
- `account_sid` - The account SID associated with the IP access control list mapping
- `domain_sid` - The domain SID associated with the IP access control list mapping
- `friendly_name` - The friendly name of the IP access control list mapping
- `date_created` - The date in RFC3339 format that the IP access control list mapping was created
- `date_updated` - The date in RFC3339 format that the IP access control list mapping was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the IP access control list mapping details
