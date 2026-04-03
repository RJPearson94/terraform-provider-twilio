---
page_title: "twilio_conversations_service_notification Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_service_notification Data Source

Use this data source to access notification configuration for a conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-notification-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_service_notification" "service_notification" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service_notification" {
  value = data.twilio_conversations_service_notification.service_notification
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this service notification
- `added_to_conversation` (List of Object) Notification settings for when a user is added to a conversation (see [below for nested schema](#nestedatt--added_to_conversation))
- `id` (String) The ID of this resource.
- `log_enabled` (Boolean) Whether notification logging is enabled for the service
- `new_message` (List of Object) Notification settings for new messages (see [below for nested schema](#nestedatt--new_message))
- `removed_from_conversation` (List of Object) Notification settings for when a user is removed from a conversation (see [below for nested schema](#nestedatt--removed_from_conversation))
- `url` (String) The absolute URL of the service notification resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--added_to_conversation"></a>
### Nested Schema for `added_to_conversation`

Read-Only:

- `enabled` (Boolean)
- `sound` (String)
- `template` (String)


<a id="nestedatt--new_message"></a>
### Nested Schema for `new_message`

Read-Only:

- `badge_count_enabled` (Boolean)
- `enabled` (Boolean)
- `sound` (String)
- `template` (String)


<a id="nestedatt--removed_from_conversation"></a>
### Nested Schema for `removed_from_conversation`

Read-Only:

- `enabled` (Boolean)
- `sound` (String)
- `template` (String)
