---
page_title: "Twilio Conversations Address Configuration - Studio"
subcategory: "Conversations"
---

# twilio_conversations_address_configuration_studio Resource

Manages address configuration for a conversation service. See the [API docs](https://www.twilio.com/docs/conversations/api/address-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> you can only configure an address once. If you specify configuration multiple times for the same address, an error will be returned

## Example Usage

### Basic

```hcl
resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address     = "+4471234567890"
  type        = "sms"
  flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  retry_count = 1
}
```

### With Studio Flow

```hcl
resource "twilio_studio_flow" "flow" {
  friendly_name = "Test Flow"
  status        = "published"
  definition = jsonencode({
    "description" : "A New Flow",
    "flags" : {
      "allow_concurrent_calls" : true
    },
    "initial_state" : "Trigger",
    "states" : [
      {
        "name" : "Trigger",
        "properties" : {
          "offset" : {
            "x" : 0,
            "y" : 0
          }
        },
        "transitions" : [],
        "type" : "trigger"
      }
    ]
  })
}

resource "twilio_conversations_address_configuration_studio" "address_configuration_studio" {
  address     = "+4471234567890"
  type        = "sms"
  flow_sid    = twilio_studio_flow.flow.sid
  retry_count = 3
}
```

## Argument Reference

The following arguments are supported:

- `address` - (Required) The phone number or whatsapp number to be configured
- `type` - (Required) The address type. Valid values include: `sms` or `whatsapp`
- `flow_sid` - (Mandatory) The SID for the Studio flow which will be called
- `retry_count` - (Mandatory) The number of attempts to retry a failed webhook call. The value must be between `0` and `3` (inclusive)
- `service_sid` - (Optional) The conversation service to associate the address configuration with. If no value is supplied the studio conversation service will be used
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
- `flow_sid` - The SID for the studio flow which will be called
- `retry_count` - The number of attempts to retry a failed webhook call
- `integration_type` - The integration type used. This should always be set to `studio`
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
terraform import twilio_conversations_address_configuration_studio.address_configuration_studio /Configuration/Addresses/IGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
