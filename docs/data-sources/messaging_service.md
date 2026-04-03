---
page_title: "twilio_messaging_service Data Source - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_service Data Source

Use this data source to access information about an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_service" "service" {
  sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_messaging_service.service
}
```

## Schema

### Required

- `sid` (String) The SID of the messaging service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this messaging service
- `area_code_geomatch` (Boolean) Whether area code geomatch is enabled on the messaging service
- `date_created` (String) The date and time the messaging service was created, in RFC 3339 format
- `date_updated` (String) The date and time the messaging service was last updated, in RFC 3339 format
- `fallback_method` (String) The HTTP method used to call the fallback URL
- `fallback_to_long_code` (Boolean) Whether fallback to long code is enabled for the messaging service
- `fallback_url` (String) The URL to call when an inbound message error is received
- `friendly_name` (String) A human-readable label for the messaging service
- `id` (String) The ID of this resource.
- `inbound_method` (String) The HTTP method used to call the inbound request URL
- `inbound_request_url` (String) The URL to call when a message is received by any phone number or short code in the messaging service
- `mms_converter` (Boolean) Whether the MMS converter is enabled for the messaging service
- `smart_encoding` (Boolean) Whether smart encoding is enabled for the messaging service
- `status_callback_url` (String) The URL to call when a message status change event is received
- `sticky_sender` (Boolean) Whether sticky sender is enabled on the messaging service
- `url` (String) The absolute URL of the messaging service resource
- `use_inbound_webhook_on_number` (Boolean) Whether the inbound webhook on the phone number is used for incoming messages
- `validity_period` (Number) The number of seconds that the messaging service will keep a message in the sending queue before it is considered failed

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
