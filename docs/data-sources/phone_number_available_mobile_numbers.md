---
page_title: "Twilio Available Mobile Phone Numbers"
subcategory: "Phone Numbers"
---

# twilio_phone_number_available_mobile_numbers Data Source

Use this data source to search for available mobile phone numbers. See the [API docs](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-mobile-resource) for more information

## Example Usage

```hcl
data "twilio_phone_number_available_mobile_numbers" "available_mobile_numbers" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  iso_country = "GB"
}

output "available_mobile_numbers" {
  value = data.twilio_phone_number_available_mobile_numbers.available_mobile_numbers.available_phone_numbers
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the phone number is associated with
- `iso_country` - (Mandatory) The ISO country to search for phone numbers
- `limit` - (Optional) The maximum number of available phone numbers to return
- `area_code` - (Optional) To search for phone numbers in an area code
- `allow_beta_numbers` - (Optional) Whether to include beta phone numbers
- `contains_number_pattern` - (Optional) The pattern to search for phone numbers
- `exclude_address_requirements` - (Optional) A `exclude_address_requirement` block as documented below
- `location` - (Optional) A `location` block as documented below
- `capabilities` - (Optional) A `capability` block as documented below

---

An `exclude_address_requirement` block supports the following:

- `all` - Whether to exclude phone numbers which have any address requirements
- `local` - Whether to exclude phone numbers which have local address requirements
- `foreign` - Whether to exclude phone numbers which have foreign address requirements

---

A `location` block supports the following:

- `in_postal_code` - To search for phone numbers in the postal area
- `in_region` - To search for phone numbers in a region
- `in_lata` - To search for phone numbers in a Local Address and Transport Area (LATA)
- `in_locality` - To search for phone numbers in a specific locality
- `in_rate_center` - To search for phone numbers in a specific rate center
- `near_number` - To search for phone numbers near an existing phone number
- `near_lat_long` - To search for phone numbers near a latitude and longitude
- `distance` - To search for phone numbers within n miles of a lat long or number

---

A `capability` block supports the following:

- `fax_enabled` - Whether to include fax enabled phone numbers
!> the `fax_enabled` attribute has been deprecated and will be removed in future version

- `sms_enabled` - Whether to include sms enabled phone numbers
- `mms_enabled` - Whether to include mms enabled phone numbers
- `voice_enabled` - Whether to include voice-enabled phone numbers

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the phone number (Same as the SID)
- `sid` - The SID of the phone number (Same as the ID)
- `account_sid` - The account SID the phone number is associated with
- `iso_country` - The ISO country of the phone number to each
- `available_phone_numbers` - A list of `available_phone_number` blocks as documented below

---

An `available_phone_number` block supports the following:

- `friendly_name` - The friendly name of the phone number
- `phone_number` - The phone number
- `address_requirements` - The address requirements of the phone number
- `capabilities` - A `capability` block as documented below
- `lata` - The Local Address and Transport Area (LATA) of the phone number
- `rate_center` - The rate centre of the phone number
- `latitude` - The latitude of the phone number's location
- `longitude` - The longitude of the phone number's location
- `locality` - The locality of the phone number's location
- `region` - The state or providence abbreviation of the phone number's location
- `postal_code` - The postal code of the phone number's location

---

A `capability` block supports the following:

- `fax` - Whether the phone number supports fax
!> the `fax` attribute has been deprecated and will be removed in future version

- `sms` - Whether the phone number supports SMS
- `mms` - Whether the phone number supports MMS
- `voice` - Whether the phone number supports voice

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the phone numbers
