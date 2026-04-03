---
page_title: "twilio_conversations_configuration Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_configuration Resource

Manages configuration for the conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversation configuration for the account. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

!> Removing the `default_service_sid`, `default_closed_timer`, `default_inactive_timer` or `default_messaging_service_sid` from your configuration will cause the corresponding value to be retained after a Terraform apply. If you want to change any of the value you will need to update your configuration to set an appropriate value

## Example Usage

```hcl
resource "twilio_conversations_configuration" "configuration" {}
```

## Schema

### Optional

- `default_closed_timer` (String) The default ISO 8601 duration after which a conversation will be automatically closed
- `default_inactive_timer` (String) The default ISO 8601 duration after which a conversation will be marked as inactive
- `default_messaging_service_sid` (String) The SID of the default messaging service
- `default_service_sid` (String) The SID of the default conversations service
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this configuration
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
- `update` (String)
