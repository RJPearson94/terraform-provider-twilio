---
page_title: "Twilio SIP Trunking Origination URL"
subcategory: "SIP Trunking"
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

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the phone number is associated with. Changing this forces a new resource to be created
- `enabled` - (Mandatory)  Whether the origination URL is enabled
- `friendly_name` - (Mandatory) The friendly name of the origination URL
- `priority` - (Mandatory) The priority/ importance of the origination URL
- `sip_url` - (Mandatory) The SIP address to route origination calls to
- `weight` - (Mandatory) The weight/ share which is used to determine where the traffic is routed with origination URL of the same priority

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the origination URL (Same as the SID)
- `sid` - The SID of the origination URL (Same as the ID)
- `account_sid` - The account SID the origination URL is associated with
- `trunk_sid` - The SID of the SIP trunk the origination URL is associated with
- `enabled` - Whether the origination URL is enabled
- `friendly_name` - The friendly name of the origination URL
- `priority` - The priority/ importance of the origination URL
- `sip_url` - The SIP address to route origination calls to
- `weight` - The weight/ share which is used to determine where the traffic is routed with origination URL of the same priority
- `date_created` - The date in RFC3339 format that the origination URL was created
- `date_updated` - The date in RFC3339 format that the origination URL was updated
- `url` - The URL of the origination URL resource

## Import

A SIP trunk origination URL can be imported using the `/Trunks/{trunkSid}/OriginationUrls/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_origination_url.origination_url /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/OriginationUrls/OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
