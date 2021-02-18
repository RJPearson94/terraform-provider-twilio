---
page_title: "Twilio SIP Trunking Phone Number"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_phone_number Data Source

Use this data source to access information about an existing phone number. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid       = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_number" {
  value = data.twilio_sip_trunking_phone_number.phone_number
}
```

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the phone number is associated with
- `sid` - (Mandatory) The SID of the phone number

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the phone number (Same as the `sid`)
- `sid` - The SID of the phone number (Same as the `id`)
- `account_sid` - The account SID the phone number is associated with
- `friendly_name` - The friendly name of the phone number
- `phone_number` - The phone number
- `address_requirements` - The address requirements of the phone number
- `beta` - Whether the phone number is in beta on the platform
- `capabilities` - A `capability` block as documented below
- `messaging` - A `messaging` block as documented below
- `trunk_sid` - The SID of the SIP trunk the phone number is associated with
- `voice` - A `voice` block as documented below
- `fax` - A `fax` block as documented below
- `status_callback_url` - The URL to call on each status change
- `status_callback_method` - The HTTP method which should be used to call the status callback URL
- `date_created` - The date in RFC3339 format that the phone number was created
- `date_updated` - The date in RFC3339 format that the phone number was updated
- `url` - The URL of the phone number resource

!> the `voice` block will be defaulted to if the Twilio API doesn't return the voice receive mode field, this data isn't being returned since Programmable Fax was disabled on some accounts

---

A `capability` block supports the following:

- `fax` - Whether the phone number supports fax
- `sms` - Whether the phone number supports SMS
- `mms` - Whether the phone number supports MMS
- `voice` - Whether the phone number supports voice

---

A `messaging` block supports the following:

- `application_sid` - The application SID which should be called on each incoming message
- `url` - The URL which should be called on each incoming message
- `method` - The HTTP method which should be used to call the URL
- `fallback_url` - The fallback URL which should be called when the URL request fails
- `fallback_method` - The HTTP method which should be used to call the fallback URL

---

A `voice` block supports the following:

- `application_sid` - The application SID which should be called on each incoming call
- `url` - The URL which should be called on each incoming call
- `method` - The HTTP method which should be used to call the URL
- `fallback_url` - The fallback URL which should be called when the URL request fails
- `fallback_method` - The HTTP method which should be used to call the fallback URL
- `caller_id_lookup` - Whether caller ID lookup is enabled for the phone number

---

A `fax` block supports the following:

- `application_sid` - The application SID which should be called on each incoming fax
- `url` - The URL which should be called on each incoming fax
- `method` - The HTTP method which should be used to call the URL
- `fallback_url` - The fallback URL which should be called when the URL request fails
- `fallback_method` - The HTTP method which should be used to call the fallback URL

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the phone number details
