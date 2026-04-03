---
page_title: "twilio_sip_domain_ip_access_control_list_mappings Data Source - twilio"
subcategory: "SIP"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account that owns the SIP domain IP access control list mappings
- `domain_sid` (String) The SID of the SIP domain to retrieve IP access control list mappings for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.
- `ip_access_control_list_mappings` (List of Object) A list of IP access control list mappings for the SIP domain (see [below for nested schema](#nestedatt--ip_access_control_list_mappings))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--ip_access_control_list_mappings"></a>
### Nested Schema for `ip_access_control_list_mappings`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
