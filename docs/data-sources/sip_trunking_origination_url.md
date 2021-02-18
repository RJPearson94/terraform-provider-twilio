---
page_title: "Twilio SIP Trunking Origination URL"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_origination_url Data Source

Use this data source to access information about an existing origination URL. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/originationurl-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid       = "OUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "origination_url" {
  value = data.twilio_sip_trunking_origination_url.origination_url
}
```

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the origination URL is associated with
- `sid` - (Mandatory) The SID of the origination URL

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the origination URL (Same as the `sid`)
- `sid` - The SID of the origination URL (Same as the `id`)
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the origination URL details
