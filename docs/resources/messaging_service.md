---
page_title: "Twilio Programmable Messaging Service"
subcategory: "Programmable Messaging"
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

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The friendly name of the service
- `area_code_geomatch` - (Optional) Whether to use attempt to use a local phone number to send a message. The default value is `true`

~> This feature is only available in specific countries, see the [Twilio docs](https://www.twilio.com/docs/messaging/services#area-code-geomatch) more information

- `fallback_method` - (Optional) The HTTP method to call the fallback URL. Valid values are `POST` or `GET`. The default value is `POST`
- `fallback_to_long_code` - (Optional) Whether to attempt to use a long code to resend a message when delivery of a message fails using a short code. The default value is `true`
- `fallback_url` - (Optional) The URL which will be called when an error occurs fetching or executing the TwiML from the inbound request URL
- `inbound_method` - (Optional) The HTTP method to call the inbound request URL. Valid values are `POST` or `GET`. The default value is `POST`
- `inbound_request_url` - (Optional) The URL which will be called when an inbound message is received for any associated short code or phone number
- `mms_converter` - (Optional) Whether to convert MMS messages to SMS messages and include a URL to the content when the carrier cannot receive MMS messages. The default value is `true`
- `smart_encoding` - (Optional) Whether to enable detection and replacement of Unicode characters that are easy to miss. The default value is `true`
- `status_callback_url` - (Optional) The URL which will be called when a message delivery status is changed
- `sticky_sender` - (Optional) Whether to ensure the end-user receives messages from the same phone number. The default value is `true`
- `use_inbound_webhook_on_number` - (Optional) Whether to use the webhook that is configured on the phone number. The default value is `false`
- `validity_period` - (Optional) How long (in seconds) messages sent from the messaging service are valid for. The value must be between `1` and `14400` (inclusive). The default value is `14400`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `sid`)
- `sid` - The SID of the service (Same as the `id`)
- `account_sid` - The Account SID associated with the service
- `friendly_name` - The friendly name of the service
- `area_code_geomatch` - Whether to use attempt to use a local phone number to send a message
- `fallback_method` - The HTTP method to call the fallback URL
- `fallback_to_long_code` - Whether to attempt to use a long code to resent a message when delivery of a message fails using a short code
- `fallback_url` - The URL which will be called when an error occurs fetching or executing the TwiML from the inbound request URL.
- `inbound_method` - The HTTP method to call the inbound request URL
- `inbound_request_url` - The URL which will be called when an inbound message is received for any associated short code or phone number
- `mms_converter` - Whether to convert MMS messages to SMS messages and include a URL to the content when the carrier cannot receive MMS messages
- `smart_encoding` - Whether to enable detection and replacement of Unicode characters that are easy to miss
- `status_callback_url` - The URL which will be called when a message delivery status is changed
- `sticky_sender` - Whether to ensure the end-user receives messages from the same phone number
- `validity_period` - How long (in seconds) messages sent from the messaging service are valid for
- `use_inbound_webhook_on_number` - Whether to use the webhook that is configured on the phone number
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the service
- `update` - (Defaults to 10 minutes) Used when updating the service
- `read` - (Defaults to 5 minutes) Used when retrieving the service
- `delete` - (Defaults to 10 minutes) Used when deleting the service

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_messaging_service.service /Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
