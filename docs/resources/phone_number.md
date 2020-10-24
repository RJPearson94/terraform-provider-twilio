---
page_title: "Twilio Phone Number"
subcategory: "Phone Numbers"
---

# twilio_phone_number Resource

Manages a phone number. See the [API docs](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) for more information

## Example Usage

```hcl
resource "twilio_phone_number" "phone_number" {
  account_sid  = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  phone_number = "+15005550006"
}

output "phone_number" {
  value = twilio_phone_number.phone_number.phone_number
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account to associate the phone number with. Changing this forces a new resource to be created.
- `friendly_name` - (Optional) The friendly name of the phone number
- `phone_number` - (Optional) The phone number to purchase. Changing this forces a new resource to be created. Conflicts with `area_code`.
- `area_code` - (Optional) The area code to purchase a phone number in. Changing this forces a new resource to be created. Conflicts with `phone_number`.
- `address_sid` - (Optional) The address SID the phone number is associated with
- `emergency_address_sid` - (Optional) The emergency address SID the phone number is associated with
- `emergency_status` - (Optional) The emergency status of the phone number
- `messaging` - (Optional) A `messaging` block as documented below
- `trunk_sid` - (Optional) The trunk SID the phone number is associated with
- `voice` - (Optional) A `voice` block as documented below. Conflicts with `fax`.
- `fax` - (Optional) A `fax` block as documented below. Conflicts with `voice`.
- `identity_sid` - (Optional) The identity SID the phone number is associated with
- `bundle_sid` - (Optional) The bundle SID the phone number is associated with
- `status_callback_url` - (Optional) The URL to call on each status change
- `status_callback_method` - (Optional) The HTTP method which should be used to call the status callback URL

~> Either the phone number or area code must be set

---

A `messaging` block supports the following:

- `application_sid` - (Optional) The application SID which should be called on each incoming message
- `url` - (Optional) The URL which should be called on each incoming message
- `method` - (Optional) The HTTP method which should be used to call the URL. Options are `GET` or `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method which should be used to call the fallback URL. Options are `GET` or `POST`

---

A `voice` block supports the following:

- `application_sid` - (Optional) The application SID which should be called on each incoming call
- `url` - (Optional) The URL which should be called on each incoming call
- `method` - (Optional) The HTTP method which should be used to call the URL. Options are `GET` or `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method which should be used to call the fallback URL. Options are `GET` or `POST`
- `caller_id_lookup` - (Optional) Whether caller ID lookup is enabled for the phone number

---

A `fax` block supports the following:

- `application_sid` - (Optional) The application SID which should be called on each incoming fax
- `url` - (Optional) The URL which should be called on each incoming fax
- `method` - (Optional) The HTTP method which should be used to call the URL. Options are `GET` or `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method which should be used to call the fallback URL. Options are `GET` or `POST`

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

- `create` - (Defaults to 10 minutes) Used when creating the phone number
- `update` - (Defaults to 10 minutes) Used when updating the phone number
- `read` - (Defaults to 5 minutes) Used when retrieving the phone number
- `delete` - (Defaults to 10 minutes) Used when deleting the phone number

## Import

A flow can be imported using the `/Accounts/{accountSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_phone_number.phone_number /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
