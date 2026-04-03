---
page_title: "twilio_phone_number Data Source - twilio"
subcategory: "Phone Numbers"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this phone number
- `sid` (String) The SID of the phone number to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `address_requirements` (String) The type of address required for this phone number
- `address_sid` (String) The SID of the address associated with this phone number
- `beta` (Boolean) Whether the phone number is a beta number new to the Twilio platform
- `bundle_sid` (String) The SID of the regulatory compliance bundle associated with this phone number
- `capabilities` (List of Object) The set of boolean capabilities of the phone number (see [below for nested schema](#nestedatt--capabilities))
- `date_created` (String) The date and time the phone number was created, in RFC 3339 format
- `date_updated` (String) The date and time the phone number was last updated, in RFC 3339 format
- `emergency_address_sid` (String) The SID of the emergency address associated with this phone number
- `emergency_status` (String) The emergency calling status of the phone number
- `fax` (List of Object) The fax settings for the phone number (see [below for nested schema](#nestedatt--fax))
- `friendly_name` (String) A human-readable label for the phone number
- `id` (String) The ID of this resource.
- `identity_sid` (String) The SID of the identity resource associated with this phone number
- `messaging` (List of Object) The messaging settings for the phone number (see [below for nested schema](#nestedatt--messaging))
- `origin` (String) The origin of the phone number, such as `twilio` or `hosted`
- `phone_number` (String) The phone number in E.164 format
- `status` (String) The current status of the phone number
- `status_callback_method` (String) The HTTP method used to call the status callback URL
- `status_callback_url` (String) The URL called for status callback events on the phone number
- `trunk_sid` (String) The SID of the SIP trunk that handles voice calls for this phone number
- `voice` (List of Object) The voice settings for the phone number (see [below for nested schema](#nestedatt--voice))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--capabilities"></a>
### Nested Schema for `capabilities`

Read-Only:

- `fax` (Boolean)
- `mms` (Boolean)
- `sms` (Boolean)
- `voice` (Boolean)


<a id="nestedatt--fax"></a>
### Nested Schema for `fax`

Read-Only:

- `application_sid` (String)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `url` (String)


<a id="nestedatt--messaging"></a>
### Nested Schema for `messaging`

Read-Only:

- `application_sid` (String)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `url` (String)


<a id="nestedatt--voice"></a>
### Nested Schema for `voice`

Read-Only:

- `application_sid` (String)
- `caller_id_lookup` (Boolean)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `url` (String)
