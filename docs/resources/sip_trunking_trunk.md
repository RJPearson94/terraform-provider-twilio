---
page_title: "twilio_sip_trunking_trunk Resource - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_trunk Resource

Manages a SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/trunk-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_trunk" "trunk" {
  friendly_name = "twilio-test"
}
```

## Schema

### Optional

- `cnam_lookup_enabled` (Boolean) Whether CNAM (Caller Name) lookup is enabled for the trunk. Defaults to `false`
- `disaster_recovery_method` (String) The HTTP method used to call the disaster recovery URL. Valid values are `GET` or `POST`
- `disaster_recovery_url` (String) The URL to call in the event of a disaster recovery failover
- `domain_name` (String) The unique domain name for the SIP trunk
- `friendly_name` (String) A human-readable label for the SIP trunk (up to 64 characters)
- `recording` (Block List, Max: 1) A block to configure recording settings for the SIP trunk (see [below for nested schema](#nestedblock--recording))
- `secure` (Boolean) Whether secure SIP (SIPS) is required for the trunk. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `transfer_mode` (String) The call transfer mode for the SIP trunk. Valid values are `enable-all`, `sip-only`, or `disable-all`. Defaults to `disable-all`

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk
- `auth_type` (String) The authentication type configured for the SIP trunk
- `auth_type_set` (List of String) The set of authentication types configured for the SIP trunk
- `date_created` (String) The date and time the SIP trunk was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP trunk by Twilio
- `url` (String) The absolute URL of the SIP trunk resource

<a id="nestedblock--recording"></a>
### Nested Schema for `recording`

Optional:

- `mode` (String) The recording mode for the SIP trunk. Valid values are `do-not-record`, `record-from-ringing`, `record-from-answer`, `record-from-ringing-dual`, or `record-from-answer-dual`. Defaults to `do-not-record`
- `trim` (String) Whether to trim silence from recordings. Valid values are `trim-silence` or `do-not-trim`. Defaults to `do-not-trim`


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A SIP trunk can be imported using the `/Trunks/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_trunk.trunk /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
