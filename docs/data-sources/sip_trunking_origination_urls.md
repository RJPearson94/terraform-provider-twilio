---
page_title: "Twilio SIP Trunking Origination URLs"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_origination_urls Data Source

Use this data source to access information about the origination URLs associated with an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/originationurl-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_origination_urls" "origination_urls" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "origination_urls" {
  value = data.twilio_sip_trunking_origination_urls.origination_urls
}
```

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the origination URLs are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `trunk_sid`)
- `trunk_sid` - The SID of the SIP trunk the origination URLs are associated with (Same as the `id`)
- `account_sid` - The account SID associated with the origination URLs
- `origination_urls` - A list of `origination_url` blocks as documented below

---

An `origination_url` block supports the following:

- `sid` - The SID of the origination URL
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

- `read` - (Defaults to 10 minutes) Used when retrieving origination URLs
