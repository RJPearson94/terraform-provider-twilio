---
page_title: "twilio_chat_service Resource - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_service Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat service. See the [API docs](https://www.twilio.com/docs/chat/rest/service-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name = "twilio-test"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the chat service

### Optional

- `limits` (Block List, Max: 1) The limits configuration for the chat service (see [below for nested schema](#nestedblock--limits))
- `media` (Block List, Max: 1) The media configuration for the chat service (see [below for nested schema](#nestedblock--media))
- `notifications` (Block List, Max: 1) The notification configuration for the chat service (see [below for nested schema](#nestedblock--notifications))
- `post_webhook_retry_count` (Number) The number of retry attempts for post-event webhook requests. Defaults to `0`
- `post_webhook_url` (String) The URL for post-event webhook requests
- `pre_webhook_retry_count` (Number) The number of retry attempts for pre-event webhook requests. Defaults to `0`
- `pre_webhook_url` (String) The URL for pre-event webhook requests
- `reachability_enabled` (Boolean) Whether the reachability indicator is enabled for the chat service. Defaults to `false`
- `read_status_enabled` (Boolean) Whether the message read status feature is enabled. Defaults to `true`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `typing_indicator_timeout` (Number) The duration in seconds after which a typing indicator times out. Defaults to `5`
- `webhook_filters` (List of String) The list of webhook event triggers to subscribe to
- `webhook_method` (String) The HTTP method used for webhook requests. Valid values are `POST` or `GET`. Defaults to `POST`

### Read-Only

- `account_sid` (String) The SID of the account that owns this chat service
- `date_created` (String) The date and time the chat service was created, in RFC 3339 format
- `date_updated` (String) The date and time the chat service was last updated, in RFC 3339 format
- `default_channel_creator_role_sid` (String) The SID of the default role assigned to the creator of a channel
- `default_channel_role_sid` (String) The SID of the default role assigned to members of a channel
- `default_service_role_sid` (String) The SID of the default role assigned to users of the chat service
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this chat service by Twilio
- `url` (String) The absolute URL of the chat service resource

<a id="nestedblock--limits"></a>
### Nested Schema for `limits`

Optional:

- `channel_members` (Number) The maximum number of members allowed in a channel. Defaults to `100`
- `user_channels` (Number) The maximum number of channels a user can join. Defaults to `250`


<a id="nestedblock--media"></a>
### Nested Schema for `media`

Optional:

- `compatibility_message` (String) The message to display when a media message has no text fallback

Read-Only:

- `size_limit_mb` (Number) The maximum media file size in megabytes


<a id="nestedblock--notifications"></a>
### Nested Schema for `notifications`

Optional:

- `added_to_channel` (Block List, Max: 1) The notification settings for being added to a channel (see [below for nested schema](#nestedblock--notifications--added_to_channel))
- `invited_to_channel` (Block List, Max: 1) The notification settings for being invited to a channel (see [below for nested schema](#nestedblock--notifications--invited_to_channel))
- `log_enabled` (Boolean) Whether notification logging is enabled. Defaults to `false`
- `new_message` (Block List, Max: 1) The notification settings for new messages (see [below for nested schema](#nestedblock--notifications--new_message))
- `removed_from_channel` (Block List, Max: 1) The notification settings for being removed from a channel (see [below for nested schema](#nestedblock--notifications--removed_from_channel))

<a id="nestedblock--notifications--added_to_channel"></a>
### Nested Schema for `notifications.added_to_channel`

Optional:

- `enabled` (Boolean) Whether added-to-channel notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for added-to-channel notifications
- `template` (String) The notification template for being added to a channel


<a id="nestedblock--notifications--invited_to_channel"></a>
### Nested Schema for `notifications.invited_to_channel`

Optional:

- `enabled` (Boolean) Whether invited-to-channel notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for invited-to-channel notifications
- `template` (String) The notification template for being invited to a channel


<a id="nestedblock--notifications--new_message"></a>
### Nested Schema for `notifications.new_message`

Optional:

- `badge_count_enabled` (Boolean) Whether the badge count is enabled for new message notifications. Defaults to `false`
- `enabled` (Boolean) Whether new message notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for new message notifications
- `template` (String) The notification template for new messages


<a id="nestedblock--notifications--removed_from_channel"></a>
### Nested Schema for `notifications.removed_from_channel`

Optional:

- `enabled` (Boolean) Whether removed-from-channel notifications are enabled. Defaults to `false`
- `sound` (String) The sound to play for removed-from-channel notifications
- `template` (String) The notification template for being removed from a channel



<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A service can be imported using the `/Services/{serviceSid}` format, e.g.

```shell
terraform import twilio_chat_service.service /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
