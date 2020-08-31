---
page_title: "Twilio Proxy Short Code"
subcategory: "Proxy"
---

# twilio_proxy_short_code Data Source

Use this data source to access information about an existing Proxy short code resource. See the [API docs](https://www.twilio.com/docs/proxy/api/short-code) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_short_code" "short_code" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_code" {
  value = data.twilio_proxy_short_code.short_code
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the short code is associated with
- `sid` - (Mandatory) The SID of the short code

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the proxy short code resource (Same as the SID)
- `sid` - The SID of a Twilio short code associated with the proxy (Same as the ID)
- `account_sid` - The account SID of the short code resource is deployed into
- `service_sid` - The SID of a Twilio proxy service
- `is_reserved` - Whether the short code is reserved
- `short_code` - The short code associated with the SID
- `iso_country` - The ISO country associated with the SID
- `capabilities` - A `capabilities` block as documented below.
- `date_created` - The date in RFC3339 format that the proxy short code resource was created
- `date_updated` - The date in RFC3339 format that the proxy short code resource was updated
- `url` - The URL of the proxy short code resource

---

A `capabilities` block supports the following:

- `fax_inbound` - Whether the short code can accept inbound faxes
- `fax_outbound` - Whether the short code can send outbound faxes
- `mms_inbound` - Whether the short code can accept inbound MMS's
- `mms_outbound` - Whether the short code can send outbound MMS's
- `restriction_fax_domestic` - Whether the short code is restricted to domestic faxes
- `restriction_mms_domestic` - Whether the short code is restricted to domestic MMS'
- `restriction_sms_domestic` - Whether the short code is restricted to domestic SMS's
- `restriction_voice_domestic` - Whether the short code is restricted to domestic voice calls
- `sip_trunking` - Whether the short code supports SIP trunking
- `sms_inbound` - Whether the short code can accept inbound SMS's
- `sms_outbound` - Whether the short code can send outbound SMS's
- `voice_inbound` - Whether the short code can accept inbound voice calls
- `voice_outbound` - Whether the short code can make outbound voice calls

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the short code
