---
page_title: "twilio_conversations_address_configuration_webhook Resource - twilio"
subcategory: "Conversations"
description: |-
  
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

## Schema

### Required

- `address` (String) The address (e.g. phone number) for the configuration. Changing this forces a new resource
- `type` (String) The type of address. Valid values are `sms` or `whatsapp`. Changing this forces a new resource
- `webhook_filters` (List of String) The list of webhook event filters to subscribe to
- `webhook_url` (String) The URL to send webhook requests to

### Optional

- `enabled` (Boolean) Whether auto-creation is enabled for this address configuration. Defaults to `true`
- `friendly_name` (String) A human-readable label for the address configuration
- `service_sid` (String) The SID of the conversations service
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `webhook_method` (String) The HTTP method for the webhook. Valid values are `GET` or `POST`. Defaults to `POST`

### Read-Only

- `account_sid` (String) The SID of the account that owns this address configuration
- `date_created` (String) The date and time the address configuration was created, in RFC 3339 format
- `date_updated` (String) The date and time the address configuration was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `integration_type` (String) The type of auto-creation integration
- `sid` (String) The unique SID assigned to this address configuration by Twilio
- `url` (String) The absolute URL of the address configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

An address configuration can be imported using the `/Configuration/Addresses/{sid}` format, e.g.

```shell
terraform import twilio_conversations_address_configuration_webhook.address_configuration_webhook /Configuration/Addresses/IGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
