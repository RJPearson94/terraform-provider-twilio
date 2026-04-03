---
page_title: "twilio_sip_ip_address Data Source - twilio"
subcategory: "SIP"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP IP address
- `ip_access_control_list_sid` (String) The SID of the IP access control list that this IP address belongs to
- `sid` (String) The SID of the SIP IP address to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `cidr_length_prefix` (Number) The CIDR prefix length for the IP address range
- `date_created` (String) The date and time the SIP IP address was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP IP address was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP IP address
- `id` (String) The ID of this resource.
- `ip_address` (String) The IP address in the access control list

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
