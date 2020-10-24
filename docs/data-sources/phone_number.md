---
page_title: "Twilio Phone Number"
subcategory: "Phone Numbers"
---

# twilio_phone_number Data Source

Use this data source to access information about an existing phone number. See the [API docs](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) for more information

## Example Usage

```hcl
data "twilio_phone_number" "phone_number" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_number" {
  value = data.twilio_phone_number.phone_number.phone_number
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the phone number is associated with
- `sid` - (Mandatory) The SID of the phone number

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the phone number (Same as the SID)
- `sid` - The SID of the phone number (Same as the ID)
- `account_sid` - The account SID the phone number is associated with
- `friendly_name` - The friendly name of the phone number
- `phone_number` - The phone number
- `address_sid` - The address SID the phone number is associated with
- `address_requirements` - The address requirements of the phone number
- `beta` - Whether the phone number is in beta on the platform
- `capabilities` - A `capability` block as documented below
- `emergency_address_sid` - The emergency address SID the phone number is associated with
- `emergency_status` - The emergency status of the phone number
- `messaging` - A `messaging` block as documented below
- `trunk_sid` - The trunk SID the phone number is associated with
- `voice` - A `voice` block as documented below
- `fax` - A `fax` block as documented below
- `identity_sid` - The identity SID the phone number is associated with
- `bundle_sid` - The bundle SID the phone number is associated with
- `status` - The status of the phone number
- `status_callback_url` - The URL to call on each status change
- `status_callback_method` - The HTTP method which should be used to call the status callback URL
- `origin` - The origin of the phone number
- `date_created` - The date in RFC3339 format that the phone number was created
- `date_updated` - The date in RFC3339 format that the phone number was updated

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
