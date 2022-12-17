---
page_title: "Twilio Conversations Address Configuration - Webhook"
subcategory: "Conversations"
---

# twilio_conversations_address_configuration_webhook Resource

Manages address configuration for a conversation service. See the [API docs](https://www.twilio.com/docs/conversations/api/address-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> you can only configure an address once. If you specify configuration multiple times for the same address, an error will be returned

## Example Usage

```hcl
resource "twilio_conversations_address_configuration_webhook" "address_configuration_webhook" {
  address         = "+4471234567890"
  type            = "type"
  webhook_filters = ["onMessageAdded"]
  webhook_url     = "https://localhost/webhook"
}
```

## Argument Reference

The following arguments are supported:

- `address` - (Required) The phone number or whatsapp number to be configured
- `type` - (Required) The address type. Valid values include: `sms` or `whatsapp`
- `webhook_url` - (Required) The URL to call when an event is triggered
- `webhook_filters` - (Required) A list of events which will trigger a call to the webhook URL. Valid values include: `onMessageAdded`, `onMessageUpdated`, `onMessageRemoved`, `onConversationUpdated`, `onConversationStateUpdated`, `onConversationRemoved`, `onParticipantAdded`, `onParticipantUpdated`, `onParticipantRemoved` or `onDeliveryUpdated`
- `webhook_method` - (Optional) The HTTP method should be used to call the webhook URL. Valid values are `GET` or `POST`. The default value is `POST`
- `service_sid` - (Optional) The conversation service to associate the address configuration with. If no value is supplied the webhook conversation service will be used
- `friendly_name` - (Optional) The friendly name of the address configuration
- `enabled` - (Optional) Whether conversation should auto-create when messages are received at the configured address

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the address configuration (Same as the `sid`)
- `sid` - The SID of the address configuration (Same as the `id`)
- `account_sid` - The account SID associated with the address configuration
- `address` - The phone number or whatsapp number that has been configured
- `type` - The address type
- `service_sid` - The conversation service associated with the address configuration
- `webhook_url` - The URL that is called when an event is triggered
- `webhook_filters` - A list of events that triggers a call to the webhook URL
- `webhook_method` - The HTTP method used to call the webhook URL
- `integration_type` - The integration type used. This should always be set to `webhook`
- `friendly_name` - The friendly name of the address configuration
- `enabled` - Whether conversation should auto-create when messages are received at the configured address
- `date_created` - The date in RFC3339 format that the address configuration was created
- `date_updated` - The date in RFC3339 format that the address configuration was updated
- `url` - The URL of the address configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the address configuration
- `update` - (Defaults to 10 minutes) Used when updating the address configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the address configuration
- `delete` - (Defaults to 10 minutes) Used when deleting the address configuration

## Import

An address configuration can be imported using the `/Configuration/Addresses/{sid}` format, e.g.

```shell
terraform import twilio_conversations_address_configuration_webhook.address_configuration_webhook /Configuration/Addresses/IGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
