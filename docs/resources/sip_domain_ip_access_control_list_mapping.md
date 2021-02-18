---
page_title: "Twilio SIP Domain IP Access Control List Mapping"
subcategory: "SIP"
---

# twilio_sip_domain_ip_access_control_list_mapping Resource

Manages an IP access control list mapping. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollistmapping-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_domain" "domain" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_name = "test.sip.twilio.com"
}

resource "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test"
}

resource "twilio_sip_domain_ip_access_control_list_mapping" "ip_access_control_list_mapping" {
  account_sid                = twilio_sip_domain.domain.account_sid
  domain_sid                 = twilio_sip_domain.domain.sid
  ip_access_control_list_sid = twilio_sip_ip_access_control_list.ip_access_control_list.sid
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the IP access control list mapping with. Changing this forces a new resource to be created
- `domain_sid` - (Mandatory) The domain SID to associate the IP access control list mapping with. Changing this forces a new resource to be created
- `ip_access_control_list_sid` - (Mandatory) The SIP IP access control list SID to associate the IP access control list mapping with. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the IP access control list mapping (Same as the `sid` & `ip_access_control_list_sid`)
- `sid` - The SID of the IP access control list mapping (Same as the `id` & `ip_access_control_list_sid`)
- `account_sid` - The account SID associated with the IP access control list mapping
- `domain_sid` - The domain SID associated with the IP access control list mapping
- `ip_access_control_list_sid` - The SIP IP access control list SID associated with the IP access control list mapping
- `friendly_name` - The friendly name of the IP access control list mapping
- `date_created` - The date in RFC3339 format that the IP access control list mapping was created
- `date_updated` - The date in RFC3339 format that the IP access control list mapping was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the IP access control list
- `read` - (Defaults to 5 minutes) Used when retrieving the IP access control list
- `delete` - (Defaults to 10 minutes) Used when deleting the IP access control list

## Import

An domain can be imported using the `Accounts/{accountSid}/Domains/{domainSid}/Auth/Calls/IpAccessControlListMappings/{sid}` format, e.g.

```shell
terraform import twilio_sip_domain_ip_access_control_list_mapping.ip_access_control_list_mapping /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Domains/DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Auth/Calls/IpAccessControlListMappings/ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
