---
page_title: "twilio_sip_trunking_phone_numbers Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_phone_numbers Data Source

Use this data source to access information about the phone numbers associated with an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_phone_numbers" "phone_numbers" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_numbers" {
  value = data.twilio_sip_trunking_phone_numbers.phone_numbers
}
```

## Schema

### Required

- `trunk_sid` (String) The SID of the SIP trunk to retrieve phone numbers for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the SIP trunk phone numbers
- `id` (String) The ID of this resource.
- `phone_numbers` (List of Object) A list of phone numbers associated with the SIP trunk (see [below for nested schema](#nestedatt--phone_numbers))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--phone_numbers"></a>
### Nested Schema for `phone_numbers`

Read-Only:

- `address_requirements` (String)
- `beta` (Boolean)
- `capabilities` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--capabilities))
- `date_created` (String)
- `date_updated` (String)
- `fax` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--fax))
- `friendly_name` (String)
- `messaging` (List of Object) (see [below for nested schema](#nestedobjatt--phone_numbers--messaging))
- `phone_number` (String)
- `sid` (String)
- `status_callback_method` (String)
- `status_callback_url` (String)
- `url` (String)
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
