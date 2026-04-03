---
page_title: "twilio_messaging_service Resource - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_service Resource

Manages a Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the messaging service

### Optional

- `area_code_geomatch` (Boolean) Whether to enable area code geomatch on the messaging service. Defaults to `true`
- `fallback_method` (String) The HTTP method used to call the fallback URL. Valid values are `POST` or `GET`. Defaults to `POST`
- `fallback_to_long_code` (Boolean) Whether to enable fallback to long code for the messaging service. Defaults to `true`
- `fallback_url` (String) The URL to call when an inbound message error is received. Must be a valid HTTP or HTTPS URL
- `inbound_method` (String) The HTTP method used to call the inbound request URL. Valid values are `POST` or `GET`. Defaults to `POST`
- `inbound_request_url` (String) The URL to call when a message is received by any phone number or short code in the messaging service. Must be a valid HTTP or HTTPS URL
- `mms_converter` (Boolean) Whether to enable the MMS converter for the messaging service. Defaults to `true`
- `smart_encoding` (Boolean) Whether to enable smart encoding for the messaging service. Defaults to `true`
- `status_callback_url` (String) The URL to call when a message status change event is received. Must be a valid HTTP or HTTPS URL
- `sticky_sender` (Boolean) Whether to enable sticky sender on the messaging service. Defaults to `true`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `use_inbound_webhook_on_number` (Boolean) Whether to use the inbound webhook on the phone number for incoming messages. Defaults to `false`
- `validity_period` (Number) The number of seconds that the messaging service will keep a message in the sending queue before it is considered failed. Value must be between `1` and `36000`. Defaults to `36000`

### Read-Only

- `account_sid` (String) The SID of the account that owns this messaging service
- `date_created` (String) The date and time the messaging service was created, in RFC 3339 format
- `date_updated` (String) The date and time the messaging service was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this messaging service by Twilio
- `url` (String) The absolute URL of the messaging service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_messaging_service.service /Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
