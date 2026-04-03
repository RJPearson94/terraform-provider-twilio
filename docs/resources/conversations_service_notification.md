---
page_title: "twilio_conversations_service_notification Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_service_notification Resource

Manages notification configuration for a conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-notification-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversations service notifications. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service. Changing this forces a new resource

### Optional

- `added_to_conversation` (Block List, Max: 1) Notification settings for when a user is added to a conversation (see [below for nested schema](#nestedblock--added_to_conversation))
- `log_enabled` (Boolean) Whether notification logging is enabled for the service. Defaults to `false`
- `new_message` (Block List, Max: 1) Notification settings for new messages (see [below for nested schema](#nestedblock--new_message))
- `removed_from_conversation` (Block List, Max: 1) Notification settings for when a user is removed from a conversation (see [below for nested schema](#nestedblock--removed_from_conversation))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this service notification
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the service notification resource

<a id="nestedblock--added_to_conversation"></a>
### Nested Schema for `added_to_conversation`

Optional:

- `enabled` (Boolean) Whether added-to-conversation notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for added-to-conversation notifications
- `template` (String) The template for added-to-conversation notifications


<a id="nestedblock--new_message"></a>
### Nested Schema for `new_message`

Optional:

- `badge_count_enabled` (Boolean) Whether badge count is enabled for new message notifications. Defaults to `false`
- `enabled` (Boolean) Whether new message notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for new message notifications
- `template` (String) The template for new message notifications


<a id="nestedblock--removed_from_conversation"></a>
### Nested Schema for `removed_from_conversation`

Optional:

- `enabled` (Boolean) Whether removed-from-conversation notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for removed-from-conversation notifications
- `template` (String) The template for removed-from-conversation notifications


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
