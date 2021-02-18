---
page_title: "Twilio Proxy Phone Numbers"
subcategory: "Proxy"
---

# twilio_proxy_phone_numbers Data Source

Use this data source to access information about the phone numbers associated with an existing Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_phone_numbers" "phone_numbers" {
  service_sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_numbers" {
  value = data.twilio_proxy_phone_numbers.phone_numbers
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the phone numbers are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `service_sid`)
- `service_sid` - The SID of the proxy service the phone numbers are associated with (Same as the `id`)
- `account_sid` - The account SID associated with the phone number
- `phone_numbers` - A list of `phone_number` blocks as documented below

---

A `phone_number` block supports the following:

- `sid` - The SID of the Twilio phone number associated with the proxy service
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

- `read` - (Defaults to 10 minutes) Used when retrieving phone numbers
