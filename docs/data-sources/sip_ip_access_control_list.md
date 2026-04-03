---
page_title: "twilio_sip_ip_access_control_list Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_ip_access_control_list Data Source

Use this data source to access information about an existing IP access control list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) for more information

## Example Usage

```hcl
data "twilio_sip_ip_access_control_list" "ip_access_control_list" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ALXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_access_control_list" {
  value = data.twilio_sip_ip_access_control_list.ip_access_control_list
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP IP access control list
- `sid` (String) The SID of the SIP IP access control list to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the SIP IP access control list was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP IP access control list was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the SIP IP access control list
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
