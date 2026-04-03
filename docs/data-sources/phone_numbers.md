---
page_title: "twilio_phone_numbers Data Source - twilio"
subcategory: "Phone Numbers"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account to retrieve phone numbers for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.
- `phone_numbers` (List of Object) A list of phone numbers associated with the account (see [below for nested schema](#nestedatt--phone_numbers))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--phone_numbers"></a>
### Nested Schema for `phone_numbers`

Read-Only:

- `address_requirements` (String)
- `address_sid` (String)
- `beta` (Boolean)
- `bundle_sid` (String)
- `capabilities` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--capabilities))
- `date_created` (String)
- `date_updated` (String)
- `emergency_address_sid` (String)
- `emergency_status` (String)
- `fax` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--fax))
- `friendly_name` (String)
- `identity_sid` (String)
- `messaging` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--messaging))
- `origin` (String)
- `phone_number` (String)
- `sid` (String)
- `status` (String)
- `status_callback_method` (String)
- `status_callback_url` (String)
- `trunk_sid` (String)
- `voice` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--voice))

<a id="nestedobjatt--phone_numbers--capabilities"></a>
### Nested Schema for `phone_numbers.capabilities`

Read-Only:

- `fax` (Boolean)
- `mms` (Boolean)
- `sms` (Boolean)
- `voice` (Boolean)


<a id="nestedobjatt--phone_numbers--fax"></a>
### Nested Schema for `phone_numbers.fax`

Read-Only:

- `application_sid` (String)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `url` (String)


<a id="nestedobjatt--phone_numbers--messaging"></a>
### Nested Schema for `phone_numbers.messaging`

Read-Only:

- `application_sid` (String)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `url` (String)


<a id="nestedobjatt--phone_numbers--voice"></a>
### Nested Schema for `phone_numbers.voice`

Read-Only:

- `application_sid` (String)
- `caller_id_lookup` (Boolean)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `url` (String)
