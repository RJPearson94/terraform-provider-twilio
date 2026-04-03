---
page_title: "twilio_chat_service Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_service Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat service. See the [API docs](https://www.twilio.com/docs/chat/rest/service-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_service" "service" {
  sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_chat_service.service
}
```

## Schema

### Required

- `sid` (String) The SID of the chat service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this chat service
- `date_created` (String) The date and time the chat service was created, in RFC 3339 format
- `date_updated` (String) The date and time the chat service was last updated, in RFC 3339 format
- `default_channel_creator_role_sid` (String) The SID of the default role assigned to the creator of a channel
- `default_channel_role_sid` (String) The SID of the default role assigned to members of a channel
- `default_service_role_sid` (String) The SID of the default role assigned to users of the chat service
- `friendly_name` (String) A human-readable label for the chat service
- `id` (String) The ID of this resource.
- `limits` (List of Object) The limits configuration for the chat service (see [below for nested schema](#nestedatt--limits))
- `media` (List of Object) The media configuration for the chat service (see [below for nested schema](#nestedatt--media))
- `notifications` (List of Object) The notification configuration for the chat service (see [below for nested schema](#nestedatt--notifications))
- `post_webhook_retry_count` (Number) The number of retry attempts for post-event webhook requests
- `post_webhook_url` (String) The URL for post-event webhook requests
- `pre_webhook_retry_count` (Number) The number of retry attempts for pre-event webhook requests
- `pre_webhook_url` (String) The URL for pre-event webhook requests
- `reachability_enabled` (Boolean) Whether the reachability indicator is enabled for the chat service
- `read_status_enabled` (Boolean) Whether the message read status feature is enabled
- `typing_indicator_timeout` (Number) The duration in seconds after which a typing indicator times out
- `url` (String) The absolute URL of the chat service resource
- `webhook_filters` (List of String) The list of webhook event triggers subscribed to
- `webhook_method` (String) The HTTP method used for webhook requests

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--limits"></a>
### Nested Schema for `limits`

Read-Only:

- `channel_members` (Number)
- `user_channels` (Number)


<a id="nestedatt--media"></a>
### Nested Schema for `media`

Read-Only:

- `compatibility_message` (String)
- `size_limit_mb` (Number)


<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Read-Only:

- `added_to_channel` (List of Object) (see [below for nested schema](#nestedobjatt--notifications--added_to_channel))
- `invited_to_channel` (List of Object) (see [below for nested schema](#nestedobjatt--notifications--invited_to_channel))
- `log_enabled` (Boolean)
- `new_message` (List of Object) (see [below for nested schema](#nestedobjatt--notifications--new_message))
- `removed_from_channel` (List of Object) (see [below for nested schema](#nestedobjatt--notifications--removed_from_channel))

<a id="nestedobjatt--notifications--added_to_channel"></a>
### Nested Schema for `notifications.added_to_channel`

Read-Only:

- `enabled` (Boolean)
- `sound` (String)
- `template` (String)


<a id="nestedobjatt--notifications--invited_to_channel"></a>
### Nested Schema for `notifications.invited_to_channel`

Read-Only:

- `enabled` (Boolean)
- `sound` (String)
- `template` (String)


<a id="nestedobjatt--notifications--new_message"></a>
### Nested Schema for `notifications.new_message`

Read-Only:

- `badge_count_enabled` (Boolean)
- `enabled` (Boolean)
- `sound` (String)
- `template` (String)


<a id="nestedobjatt--notifications--removed_from_channel"></a>
### Nested Schema for `notifications.removed_from_channel`

Read-Only:

- `enabled` (Boolean)
- `sound` (String)
- `template` (String)
