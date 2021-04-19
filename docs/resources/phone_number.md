---
page_title: "Twilio Phone Number"
subcategory: "Phone Numbers"
---

# twilio_phone_number Resource

Manages a phone number. See the [API docs](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) for more information

!> Removing the `friendly_name` or `emergency_status` from your configuration will cause the corresponding value to be retained after a Terraform apply. If you want to change any of the value you will need to update your configuration to set an appropriate value

## Example Usage

### With supplied phone number

```hcl
resource "twilio_phone_number" "phone_number" {
  account_sid  = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  phone_number = "+15005550006"
}

output "phone_number" {
  value = twilio_phone_number.phone_number.phone_number
}
```

### With search criteria

```hcl
data "twilio_account_details" "account_details" {}

resource "twilio_phone_number" "phone_number" {
  account_sid = data.twilio_account_details.account_details.sid

  search_criteria {
    type        = "mobile"
    iso_country = "GB"

    exclude_address_requirements {
      all = true
    }
  }
}

output "phone_number" {
  value = twilio_phone_number.phone_number.phone_number
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account to associate the phone number with. Changing this forces a new resource to be created.
- `friendly_name` - (Optional) The friendly name of the phone number
- `phone_number` - (Optional) The phone number to purchase. Changing this forces a new resource to be created. Conflicts with `area_code` and `search_criteria`.
- `area_code` - (Optional) The area code to purchase a phone number in. Changing this forces a new resource to be created. Conflicts with `phone_number` and `search_criteria`.
- `search_criteria` - (Optional) A `search_criteria` block as documented below. Conflicts with `area_code` and `phone_number`.
- `address_sid` - (Optional) The address SID the phone number is associated with
- `emergency_address_sid` - (Optional) The emergency address SID the phone number is associated with
- `emergency_status` - (Optional) The emergency status of the phone number. Valid values are `Active` or `Inactive`
- `messaging` - (Optional) A `messaging` block as documented below
- `trunk_sid` - (Optional) The trunk SID the phone number is associated with
- `voice` - (Optional) A `voice` block as documented below. Conflicts with `fax`.
- `fax` - (Optional) A `fax` block as documented below. Conflicts with `voice`.
- `identity_sid` - (Optional) The identity SID the phone number is associated with
- `bundle_sid` - (Optional) The bundle SID the phone number is associated with
- `status_callback_url` - (Optional) The URL to call on each status change
- `status_callback_method` - (Optional) The HTTP method that should be used to call the status callback URL. The default value is `POST`

~> Either the `phone_number`, `area_code` or `search_criteria` must be set

!> if the Twilio API doesn't return the voice receive mode field (this field hasn't been returned since Programmable Fax was disabled on some projects), then the provider will assume the configuration is for voice

---

A `messaging` block supports the following:

- `application_sid` - (Optional) The application SID which should be called on each incoming message
- `url` - (Optional) The URL which should be called on each incoming message
- `method` - (Optional) The HTTP method that should be used to call the URL. Valid values are `GET` or `POST`. The default value is `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method that should be used to call the fallback URL. Valid values are `GET` or `POST`. The default value is `POST`

---

A `voice` block supports the following:

- `application_sid` - (Optional) The application SID which should be called on each incoming call
- `url` - (Optional) The URL which should be called on each incoming call
- `method` - (Optional) The HTTP method that should be used to call the URL. Valid values are `GET` or `POST`. The default value is `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method that should be used to call the fallback URL. Valid values are `GET` or `POST`. The default value is `POST`
- `caller_id_lookup` - (Optional) Whether caller ID lookup is enabled for the phone number. The default value is `false`

---

A `fax` block supports the following:

- `application_sid` - (Optional) The application SID which should be called on each incoming fax
- `url` - (Optional) The URL which should be called on each incoming fax
- `method` - (Optional) The HTTP method that should be used to call the URL. Valid values are `GET` or `POST`. The default value is `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method that should be used to call the fallback URL. Valid values are `GET` or `POST`. The default value is `POST`

---

- `type` - (Mandatory) The type of phone number to purchase. Valid values are `local`, `mobile` or `toll_free`
- `iso_country` - (Mandatory) The ISO country to find a phone number
- `area_code` - (Optional) To find a phone number in an area code
- `allow_beta_numbers` - (Optional) Whether to include beta phone number in the search
- `contains_number_pattern` - (Optional) The pattern to find a phone numbers
- `exclude_address_requirements` - (Optional) A `exclude_address_requirement` block as documented below
- `location` - (Optional) A `location` block as documented below
- `capabilities` - (Optional) A `capability` block as documented below

---

An `exclude_address_requirement` block supports the following:

- `all` - Whether to find a phone number that does not have any address requirements
- `local` - Whether to find a phone number that does not have local address requirements
- `foreign` - Whether to find a phone number that does not have foreign address requirements

---

A `location` block supports the following:

- `in_postal_code` - To find a phone number in the postal area
- `in_region` - To find a phone number in a region
- `in_lata` - To find a phone number in a Local Address and Transport Area (LATA)
- `in_locality` - To find a phone number in a specific locality
- `in_rate_center` - To find a phone number in a specific rate center
- `near_number` - To find a phone number near an existing phone number
- `near_lat_long` - To find a phone number near a latitude and longitude
- `distance` - To find a phone number within n miles of a lat long or number

---

A `capability` block supports the following:

- `fax_enabled` - Whether to find a fax-enabled phone number
- `sms_enabled` - Whether to find an sms-enabled phone number
- `mms_enabled` - Whether to find an mms-enabled phone number
- `voice_enabled` - Whether to find a voice-enabled phone number

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the phone number (Same as the `sid`)
- `sid` - The SID of the phone number (Same as the `id`)
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

- `create` - (Defaults to 10 minutes) Used when creating the phone number
- `update` - (Defaults to 10 minutes) Used when updating the phone number
- `read` - (Defaults to 5 minutes) Used when retrieving the phone number
- `delete` - (Defaults to 10 minutes) Used when deleting the phone number

## Import

A flow can be imported using the `/Accounts/{accountSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_phone_number.phone_number /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> `search_criteria` cannot be imported
