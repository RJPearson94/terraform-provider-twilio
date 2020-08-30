---
page_title: "Twilio Proxy Short Codes"
subcategory: "Proxy"
---

# twilio_proxy_short_codes Data Source

Use this data source to access information about the short codes associated with an existing Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/short-code) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_short_codes" "short_codes" {
  service_sid  = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_codes" {
  value = data.twilio_proxy_short_codes.short_codes
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the short codes are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `service_sid` - The SID of the proxy service the short codes are associated with
- `account_sid` - The account SID associated with the short code
- `short_codes` - A list of `short_code` blocks as documented below

---

A `short_code` block supports the following:

- `sid` - The SID of the Twilio short code to associated with the proxy service
- `is_reserved` - Whether the short code is reserved
- `short_code` - The short code associated with the SID
- `iso_country` - The ISO country associated with the SID
- `capabilities` - A `capabilities` block as documented below.
- `date_created` - The date in RFC3339 format that the proxy short code resource was created
- `date_updated` - The date in RFC3339 format that the proxy short code resource was updated
- `url` - The url of the proxy short code resource

---

A `capabilities` block supports the following:

- `fax_inbound` - Whether the short code is able to accept inbound faxes
- `fax_outbound` - Whether the short code is able to send outbound faxes
- `mms_inbound` - Whether the short code is able to accept inbound MMS's
- `mms_outbound` - Whether the short code is able to send outbound MMS's
- `restriction_fax_domestic` - Whether the short code is restricted to domestic faxes
- `restriction_mms_domestic` - Whether the short code is restricted to domestic MMS'
- `restriction_sms_domestic` - Whether the short code is restricted to domestic SMS's
- `restriction_voice_domestic` - Whether the short code is restricted to domestic voice calls
- `sip_trunking` - Whether the short code supports SIP trunking
- `sms_inbound` - Whether the short code is able to accept inbound SMS's
- `sms_outbound` - Whether the short code is able to send outbound SMS's
- `voice_inbound` - Whether the short code is able to accept inbound voice calls
- `voice_outbound` - Whether the short code is able to make outbound voice calls

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving short codes
