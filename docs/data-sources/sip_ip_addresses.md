---
page_title: "twilio_sip_ip_addresses Data Source - twilio"
subcategory: "SIP"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account that owns the SIP IP addresses
- `ip_access_control_list_sid` (String) The SID of the IP access control list to retrieve IP addresses from

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.
- `ip_addresses` (List of Object) A list of IP addresses in the IP access control list (see [below for nested schema](#nestedatt--ip_addresses))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--ip_addresses"></a>
### Nested Schema for `ip_addresses`

Read-Only:

- `cidr_length_prefix` (Number)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `ip_address` (String)
- `sid` (String)
