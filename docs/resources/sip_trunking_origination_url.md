---
page_title: "twilio_sip_trunking_origination_url Resource - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_origination_url Resource

Manages a SIP trunk origination URL. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/originationurl-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid     = twilio_sip_trunking_trunk.trunk.sid
  friendly_name = "twilio-test"
  enabled       = true
  priority      = 1
  sip_url       = "sip:test@test.com"
  weight        = 1
}
```

## Schema

### Required

- `enabled` (Boolean) Whether the origination URL is enabled and available for use
- `friendly_name` (String) A human-readable label for the SIP trunk origination URL
- `priority` (Number) The priority of the origination URL, from 0 to 65535. Lower values have higher priority
- `sip_url` (String) The SIP address to route origination calls to, must start with `sip:`
- `trunk_sid` (String) The SID of the SIP trunk to add the origination URL to. Changing this forces a new resource
- `weight` (Number) The weight of the origination URL, from 0 to 65535. Used for load balancing among URLs with the same priority

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk origination URL
- `date_created` (String) The date and time the SIP trunk origination URL was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk origination URL was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP trunk origination URL by Twilio
- `url` (String) The absolute URL of the SIP trunk origination URL resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A SIP trunk origination URL can be imported using the `/Trunks/{trunkSid}/OriginationUrls/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_origination_url.origination_url /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
