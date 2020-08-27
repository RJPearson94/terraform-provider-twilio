---
page_title: "Twilio Proxy Phone Number"
subcategory: "Proxy"
---

# twilio_proxy_phone_number Resource

Manages a proxy phone number

## Example Usage

```hcl
resource "twilio_proxy_service" "service" {
  unique_name = "Test Proxy Service"
}

resource "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_service.service.sid
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  is_reserved = true
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of a Twilio proxy service. Changing this forces a new resource to be created
- `sid` - (Optional) The SID of a Twilio phone number to associate with the proxy. Changing this forces a new resource to be created. Conflicts with `phone_number`
- `phone_number` - (Optional) The phone number to associate with the proxy. Changing this forces a new resource to be created. Conflicts with `sid`
- `is_reserved` - (Optional) Whether the phone number is reserved

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the proxy phone number resource (Same as the SID)
- `sid` - The SID of a Twilio phone number to associate with the proxy (Same as the ID)
- `account_sid` - The Account SID of the phone number resource is deployed into
- `service_sid` - The SID of a Twilio proxy service
- `is_reserved` - Whether the phone number is reserved
- `phone_number` - The phone number to associate with the proxy
- `friendly_name` - The friendly name of the phone number
- `iso_country` - The ISO country of the phone number
- `in_use` - The number of active calls associated with the phone number
- `capabilities` - A `capabilities` block as documented below.
- `date_created` - The date in RFC3339 format that the proxy phone number resource was created
- `date_updated` - The date in RFC3339 format that the proxy phone number resource was updated
- `url` - The url of the proxy phone number resource

---

A `capabilities` block supports the following:

- `fax_inbound` - Whether the phone number is able to accept inbound faxes
- `fax_outbound` - Whether the phone number is able to send outbound faxes
- `mms_inbound` - Whether the phone number is able to accept inbound MMS's
- `mms_outbound` - Whether the phone number is able to send outbound MMS's
- `restriction_fax_domestic` - Whether the phone number is restricted to domestic faxes
- `restriction_mms_domestic` - Whether the phone number is restricted to domestic MMS's
- `restriction_sms_domestic` - Whether the phone number is restricted to domestic SMS's
- `restriction_voice_domestic` - Whether the phone number is restricted to domestic voice calls
- `sip_trunking` - Whether the phone number supports SIP trunking
- `sms_inbound` - Whether the phone number is able to accept inbound SMS's
- `sms_outbound` - Whether the phone number is able to send outbound SMS's
- `voice_inbound` - Whether the phone number is able to accept inbound voice calls
- `voice_outbound` - Whether the phone number is able to make outbound voice calls

---

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the phone number
- `update` - (Defaults to 10 minutes) Used when updating the phone number
- `read` - (Defaults to 5 minutes) Used when retrieving the phone number
- `delete` - (Defaults to 10 minutes) Used when deleting the phone number

## Import

A phone number can be imported using the `/Services/{serviceSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_proxy_phone_number.phone_number /Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
