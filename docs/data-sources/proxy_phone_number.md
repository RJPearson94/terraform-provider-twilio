---
page_title: "Twilio Proxy Phone Number"
subcategory: "Proxy"
---

# twilio_proxy_phone_number Data Source

Use this data source to access information about an existing Proxy phone number. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_phone_number" "phone_number" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_number" {
  value = data.twilio_proxy_phone_number.phone_number
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the phone number is associated with
- `sid` - (Mandatory) The SID of the phone number

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the proxy phone number resource (Same as the SID)
- `sid` - The SID of a Twilio phone number associated with the proxy (Same as the ID)
- `account_sid` - The account SID of the phone number resource is deployed into
- `service_sid` - The SID of a Twilio proxy service
- `is_reserved` - Whether the phone number is reserved
- `phone_number` - The phone number associated with the proxy
- `friendly_name` - The friendly name of the phone number
- `iso_country` - The ISO country of the phone number
- `in_use` - The number of active calls associated with the phone number
- `capabilities` - A `capabilities` block as documented below.
- `date_created` - The date in RFC3339 format that the proxy phone number resource was created
- `date_updated` - The date in RFC3339 format that the proxy phone number resource was updated
- `url` - The URL of the proxy phone number resource

---

A `capabilities` block supports the following:

- `fax_inbound` - Whether the phone number can accept inbound faxes
- `fax_outbound` - Whether the phone number can send outbound faxes
- `mms_inbound` - Whether the phone number can accept inbound MMS's
- `mms_outbound` - Whether the phone number can send outbound MMS's
- `restriction_fax_domestic` - Whether the phone number is restricted to domestic faxes
- `restriction_mms_domestic` - Whether the phone number is restricted to domestic MMS's
- `restriction_sms_domestic` - Whether the phone number is restricted to domestic SMS's
- `restriction_voice_domestic` - Whether the phone number is restricted to domestic voice calls
- `sip_trunking` - Whether the phone number supports SIP trunking
- `sms_inbound` - Whether the phone number can accept inbound SMS's
- `sms_outbound` - Whether the phone number can send outbound SMS's
- `voice_inbound` - Whether the phone number can accept inbound voice calls
- `voice_outbound` - Whether the phone number can make outbound voice calls

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the phone number
