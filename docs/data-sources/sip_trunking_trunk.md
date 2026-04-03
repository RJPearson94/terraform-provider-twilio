---
page_title: "twilio_sip_trunking_trunk Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_trunk Data Source

Use this data source to access information about an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/trunk-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_trunk" "trunk" {
  sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "trunk" {
  value = data.twilio_sip_trunking_trunk.trunk
}
```

## Schema

### Required

- `sid` (String) The SID of the SIP trunk to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk
- `auth_type` (String) The authentication type configured for the SIP trunk
- `auth_type_set` (List of String) The set of authentication types configured for the SIP trunk
- `cnam_lookup_enabled` (Boolean) Whether CNAM (Caller Name) lookup is enabled for the trunk
- `date_created` (String) The date and time the SIP trunk was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk was last updated, in RFC 3339 format
- `disaster_recovery_method` (String) The HTTP method used to call the disaster recovery URL
- `disaster_recovery_url` (String) The URL called in the event of a disaster recovery failover
- `domain_name` (String) The unique domain name for the SIP trunk
- `friendly_name` (String) A human-readable label for the SIP trunk
- `id` (String) The ID of this resource.
- `recording` (List of Object) The recording settings for the SIP trunk (see [below for nested schema](#nestedatt--recording))
- `secure` (Boolean) Whether secure SIP (SIPS) is required for the trunk
- `transfer_mode` (String) The call transfer mode for the SIP trunk
- `url` (String) The absolute URL of the SIP trunk resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--recording"></a>
### Nested Schema for `recording`

Read-Only:

- `mode` (String)
- `trim` (String)
