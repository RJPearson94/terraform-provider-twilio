---
page_title: "Twilio Phone Numbers"
subcategory: "Phone Numbers"
---

# twilio_phone_numbers Data Source

Use this data source to access information about the phone numbers associated with an existing account. See the [API docs](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) for more information

## Example Usage

```hcl
data "twilio_phone_numbers" "phone_numbers" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_numbers" {
  value = data.twilio_phone_numbers.phone_numbers
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the phone number is associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the account SID)
- `account_sid` - The SID of the account the queues are associated with
- `phone_numbers` - A list of `phone_number` blocks as documented below

---

An `phone_number` block supports the following:

- `sid` - The SID of the phone number
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

- `read` - (Defaults to 10 minutes) Used when retrieving the phone number details
