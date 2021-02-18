---
page_title: "Twilio SIP Domain IP Access Control List Mappings"
subcategory: "SIP"
---

# twilio_sip_domain_ip_access_control_list_mappings Data Source

Use this data source to access information about an existing IP access control list mappings associated with an existing account and domain. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource) for more information

## Example Usage

```hcl
data "twilio_sip_domain_ip_access_control_list_mappings" "ip_access_control_list_mappings" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_sid  = "DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_access_control_list_mappings" {
  value = data.twilio_sip_domain_ip_access_control_list_mappings.ip_access_control_list_mappings
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the IP access control list mappings are associated with
- `domain_sid` - (Mandatory) The SID of the domain the IP access control list mappings are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `account_sid/domain_sid`
- `account_sid` - The SID of the account the IP access control list mappings are associated with
- `domain_sid` - The SID of the domain the IP access control list mappings are associated with
- `ip_access_control_list_mappings` - A list of `ip_access_control_list_mapping` blocks as documented below

---

An `ip_access_control_list_mapping` block supports the following:

- `sid` - The SID of the IP access control list mapping
- `friendly_name` - The friendly name of the IP access control list mapping
- `date_created` - The date in RFC3339 format that the IP access control list mapping was created
- `date_updated` - The date in RFC3339 format that the IP access control list mapping was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the IP access control list mappings
