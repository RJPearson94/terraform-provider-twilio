---
page_title: "twilio_phone_number_available_local_numbers Data Source - twilio"
subcategory: "Phone Numbers"
description: |-
  
---

# twilio_phone_number_available_local_numbers Data Source

!> This resource is deprecated and will be removed in a future release of the provider. As the data source is refreshed on every plan, this data source cannot be used to purchase a phone number. Please use the `search_criteria` block on the `twilio_phone_number` resource instead

Use this data source to search for available local phone numbers. See the [API docs](https://www.twilio.com/docs/phone-numbers/api/availablephonenumber-local-resource) for more information

## Example Usage

```hcl
data "twilio_phone_number_available_local_numbers" "available_local_numbers" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  iso_country = "GB"
}

output "available_local_numbers" {
  value = data.twilio_phone_number_available_local_numbers.available_local_numbers.available_phone_numbers
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account to search for available local phone numbers
- `iso_country` (String) The ISO 3166-1 alpha-2 country code to search for available local phone numbers

### Optional

- `allow_beta_numbers` (Boolean) Whether to include beta phone numbers in the search results
- `area_code` (Number) The area code to filter available phone numbers by
- `capabilities` (Block List, Max: 1) A block to filter available phone numbers by their capabilities (see [below for nested schema](#nestedblock--capabilities))
- `contains_number_pattern` (String) A pattern to match phone numbers against, using '*' as a wildcard
- `exclude_address_requirements` (Block List, Max: 1) A block to exclude phone numbers that require an address (see [below for nested schema](#nestedblock--exclude_address_requirements))
- `limit` (Number) The maximum number of results to return
- `location` (Block List, Max: 1) A block to filter available phone numbers by location (see [below for nested schema](#nestedblock--location))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `available_phone_numbers` (List of Object) A list of available local phone numbers matching the search criteria (see [below for nested schema](#nestedatt--available_phone_numbers))
- `id` (String) The ID of this resource.

<a id="nestedblock--capabilities"></a>
### Nested Schema for `capabilities`

Optional:

- `fax_enabled` (Boolean) Whether to filter for fax-capable phone numbers
- `mms_enabled` (Boolean) Whether to filter for MMS-capable phone numbers
- `sms_enabled` (Boolean) Whether to filter for SMS-capable phone numbers
- `voice_enabled` (Boolean) Whether to filter for voice-capable phone numbers


<a id="nestedblock--exclude_address_requirements"></a>
### Nested Schema for `exclude_address_requirements`

Optional:

- `all` (Boolean) Whether to exclude phone numbers that require any address
- `foreign` (Boolean) Whether to exclude phone numbers that require a foreign address
- `local` (Boolean) Whether to exclude phone numbers that require a local address


<a id="nestedblock--location"></a>
### Nested Schema for `location`

Optional:

- `distance` (Number) The distance in miles from the `near_number` or `near_lat_long` to search within
- `in_lata` (String) The LATA to filter available phone numbers by
- `in_locality` (String) The locality (city) to filter available phone numbers by
- `in_postal_code` (String) The postal code to filter available phone numbers by
- `in_rate_center` (String) The rate center to filter available phone numbers by
- `in_region` (String) The region (state or province) to filter available phone numbers by
- `near_lat_long` (String) A latitude/longitude coordinate pair to search near, specified as `latitude,longitude`
- `near_number` (String) A phone number to search near


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--available_phone_numbers"></a>
### Nested Schema for `available_phone_numbers`

Read-Only:

- `address_requirements` (String)
- `beta` (Boolean)
- `capabilities` (List of Object) (see [below for nested schema](#nestedobjatt--available_phone_numbers--capabilities))
- `friendly_name` (String)
- `lata` (String)
- `latitude` (String)
- `locality` (String)
- `longitude` (String)
- `phone_number` (String)
- `postal_code` (String)
- `rate_center` (String)
- `region` (String)

<a id="nestedobjatt--available_phone_numbers--capabilities"></a>
### Nested Schema for `available_phone_numbers.capabilities`

Read-Only:

- `fax` (Boolean)
- `mms` (Boolean)
- `sms` (Boolean)
- `voice` (Boolean)
