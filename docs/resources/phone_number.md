---
page_title: "twilio_phone_number Resource - twilio"
subcategory: "Phone Numbers"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this phone number. Changing this forces a new resource

### Optional

- `address_sid` (String) The SID of the address associated with this phone number
- `area_code` (String) The area code of the phone number to purchase. Exactly one of `phone_number`, `area_code`, or `search_criteria` must be specified. Changing this forces a new resource
- `bundle_sid` (String) The SID of the regulatory compliance bundle associated with this phone number
- `emergency_address_sid` (String) The SID of the emergency address associated with this phone number
- `emergency_status` (String) The emergency calling status of the phone number. Valid values are `Active` or `Inactive`
- `fax` (Block List, Max: 1) A block to configure fax settings for the phone number. Conflicts with `voice` (see [below for nested schema](#nestedblock--fax))
- `friendly_name` (String) A human-readable label for the phone number
- `identity_sid` (String) The SID of the identity resource associated with this phone number
- `messaging` (Block List, Max: 1) A block to configure messaging settings for the phone number (see [below for nested schema](#nestedblock--messaging))
- `phone_number` (String) The phone number in E.164 format to purchase. Exactly one of `phone_number`, `area_code`, or `search_criteria` must be specified. Changing this forces a new resource
- `search_criteria` (Block List, Max: 1) A block to define search criteria for finding an available phone number to purchase. Exactly one of `phone_number`, `area_code`, or `search_criteria` must be specified. Changing this forces a new resource (see [below for nested schema](#nestedblock--search_criteria))
- `status_callback_method` (String) The HTTP method used to call the status callback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_url` (String) The URL to call for status callback events on the phone number
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `trunk_sid` (String) The SID of the SIP trunk to route voice calls to
- `voice` (Block List, Max: 1) A block to configure voice settings for the phone number. Conflicts with `fax` (see [below for nested schema](#nestedblock--voice))

### Read-Only

- `address_requirements` (String) The type of address required for this phone number
- `beta` (Boolean) Whether the phone number is a beta number new to the Twilio platform
- `capabilities` (List of Object) The set of boolean capabilities of the phone number (see [below for nested schema](#nestedatt--capabilities))
- `date_created` (String) The date and time the phone number was created, in RFC 3339 format
- `date_updated` (String) The date and time the phone number was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `origin` (String) The origin of the phone number, such as `twilio` or `hosted`
- `sid` (String) The unique SID assigned to this phone number by Twilio
- `status` (String) The current status of the phone number

<a id="nestedblock--fax"></a>
### Nested Schema for `fax`

Optional:

- `application_sid` (String) The SID of the TwiML application to handle incoming faxes
- `fallback_method` (String) The HTTP method used to call the fax fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `fallback_url` (String) The URL to call when an error occurs while retrieving or executing the TwiML for incoming faxes
- `method` (String) The HTTP method used to call the fax URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `url` (String) The URL to call when the phone number receives an incoming fax


<a id="nestedblock--messaging"></a>
### Nested Schema for `messaging`

Optional:

- `application_sid` (String) The SID of the TwiML application to handle incoming messages
- `fallback_method` (String) The HTTP method used to call the messaging fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `fallback_url` (String) The URL to call when an error occurs while retrieving or executing the TwiML for incoming messages
- `method` (String) The HTTP method used to call the messaging URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `url` (String) The URL to call when the phone number receives an incoming message


<a id="nestedblock--search_criteria"></a>
### Nested Schema for `search_criteria`

Required:

- `iso_country` (String) The ISO 3166-1 alpha-2 country code to search for available phone numbers. Changing this forces a new resource
- `type` (String) The type of phone number to search for. Valid values are `local`, `mobile`, or `toll_free`. Changing this forces a new resource

Optional:

- `allow_beta_numbers` (Boolean) Whether to include beta phone numbers in the search results. Changing this forces a new resource
- `area_code` (Number) The area code to filter available phone numbers by. Changing this forces a new resource
- `capabilities` (Block List, Max: 1) A block to filter available phone numbers by their capabilities. Changing this forces a new resource (see [below for nested schema](#nestedblock--search_criteria--capabilities))
- `contains_number_pattern` (String) A pattern to match phone numbers against, using '*' as a wildcard. Changing this forces a new resource
- `exclude_address_requirements` (Block List, Max: 1) A block to exclude phone numbers that require an address. Changing this forces a new resource (see [below for nested schema](#nestedblock--search_criteria--exclude_address_requirements))
- `location` (Block List, Max: 1) A block to filter available phone numbers by location. Changing this forces a new resource (see [below for nested schema](#nestedblock--search_criteria--location))

<a id="nestedblock--search_criteria--capabilities"></a>
### Nested Schema for `search_criteria.capabilities`

Optional:

- `fax_enabled` (Boolean) Whether to filter for fax-capable phone numbers. Changing this forces a new resource
- `mms_enabled` (Boolean) Whether to filter for MMS-capable phone numbers. Changing this forces a new resource
- `sms_enabled` (Boolean) Whether to filter for SMS-capable phone numbers. Changing this forces a new resource
- `voice_enabled` (Boolean) Whether to filter for voice-capable phone numbers. Changing this forces a new resource


<a id="nestedblock--search_criteria--exclude_address_requirements"></a>
### Nested Schema for `search_criteria.exclude_address_requirements`

Optional:

- `all` (Boolean) Whether to exclude phone numbers that require any address. Changing this forces a new resource
- `foreign` (Boolean) Whether to exclude phone numbers that require a foreign address. Changing this forces a new resource
- `local` (Boolean) Whether to exclude phone numbers that require a local address. Changing this forces a new resource


<a id="nestedblock--search_criteria--location"></a>
### Nested Schema for `search_criteria.location`

Optional:

- `distance` (Number) The distance in miles from the `near_number` or `near_lat_long` to search within. Changing this forces a new resource
- `in_lata` (String) The LATA to filter available phone numbers by. Changing this forces a new resource
- `in_locality` (String) The locality (city) to filter available phone numbers by. Changing this forces a new resource
- `in_postal_code` (String) The postal code to filter available phone numbers by. Changing this forces a new resource
- `in_rate_center` (String) The rate center to filter available phone numbers by. Changing this forces a new resource
- `in_region` (String) The region (state or province) to filter available phone numbers by. Changing this forces a new resource
- `near_lat_long` (String) A latitude/longitude coordinate pair to search near, specified as `latitude,longitude`. Changing this forces a new resource
- `near_number` (String) A phone number to search near. Changing this forces a new resource



<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedblock--voice"></a>
### Nested Schema for `voice`

Optional:

- `application_sid` (String) The SID of the TwiML application to handle incoming voice calls
- `caller_id_lookup` (Boolean) Whether to perform a caller ID lookup on incoming voice calls. Defaults to `false`
- `fallback_method` (String) The HTTP method used to call the voice fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `fallback_url` (String) The URL to call when an error occurs while retrieving or executing the TwiML for incoming voice calls
- `method` (String) The HTTP method used to call the voice URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `url` (String) The URL to call when the phone number receives an incoming voice call


<a id="nestedatt--capabilities"></a>
### Nested Schema for `capabilities`

Read-Only:

- `fax` (Boolean)
- `mms` (Boolean)
- `sms` (Boolean)
- `voice` (Boolean)

## Import

A flow can be imported using the `/Accounts/{accountSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_phone_number.phone_number /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> `search_criteria` cannot be imported
