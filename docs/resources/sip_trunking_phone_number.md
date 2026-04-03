---
page_title: "twilio_sip_trunking_phone_number Resource - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_phone_number Resource

Manages a SIP trunk phone number. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/phonenumber-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid        = twilio_sip_trunking_trunk.trunk.sid
  phone_number_sid = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

## Schema

### Required

- `phone_number_sid` (String) The SID of the incoming phone number to associate with the trunk. Changing this forces a new resource
- `trunk_sid` (String) The SID of the SIP trunk to associate the phone number with. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this SIP trunk phone number
- `address_requirements` (String) The type of address required for this phone number
- `beta` (Boolean) Whether the phone number is a beta number new to the Twilio platform
- `capabilities` (List of Object) The set of boolean capabilities of the phone number (see [below for nested schema](#nestedatt--capabilities))
- `date_created` (String) The date and time the SIP trunk phone number was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP trunk phone number was last updated, in RFC 3339 format
- `fax` (List of Object) The fax settings for the phone number (see [below for nested schema](#nestedatt--fax))
- `friendly_name` (String) A human-readable label for the SIP trunk phone number
- `id` (String) The ID of this resource.
- `messaging` (List of Object) The messaging settings for the phone number (see [below for nested schema](#nestedatt--messaging))
- `phone_number` (String) The phone number in E.164 format
- `sid` (String) The unique SID assigned to this SIP trunk phone number by Twilio
- `status_callback_method` (String) The HTTP method used to call the status callback URL
- `status_callback_url` (String) The URL called for status callback events on the phone number
- `url` (String) The absolute URL of the SIP trunk phone number resource
- `voice` (List of Object) The voice settings for the phone number (see [below for nested schema](#nestedatt--voice))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
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

## Import

A SIP trunk phone number can be imported using the `/Trunks/{trunkSid}/PhoneNumbers/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_phone_number.phone_number /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/PhoneNumbers/PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
