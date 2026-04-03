---
page_title: "twilio_sip_trunking_ip_access_control_lists Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_ip_access_control_lists Data Source

Use this data source to access information about the IP access control lists associated with an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/ipaccesscontrollist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_ip_access_control_lists" "ip_access_control_lists" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "ip_access_control_lists" {
  value = data.twilio_sip_trunking_ip_access_control_lists.ip_access_control_lists
}
```

## Schema

### Required

- `trunk_sid` (String) The SID of the SIP trunk to retrieve IP access control lists for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the SIP trunk IP access control lists
- `id` (String) The ID of this resource.
- `ip_access_control_lists` (List of Object) A list of IP access control lists associated with the SIP trunk (see [below for nested schema](#nestedatt--ip_access_control_lists))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--ip_access_control_lists"></a>
### Nested Schema for `ip_access_control_lists`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
- `url` (String)
