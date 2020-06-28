---
page_title: "Twilio Programmable Chat Service"
subcategory: "Programmable Chat"
---

# twilio_chat_service Resource

Manages a Chat Service

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name   = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The Friendly Name of the Chat Service
- `limits` - (Optional) A `limits` block as documented below.
- `media` - (Optional) A `media` block as documented below.
- `notifications` - (Optional) A `notifications` block as documented below.
- `post_webhook_retry_count` - (Optional) The number of attempt to retry a failed webhook call
- `post_webhook_url` - (Optional) The webhook url
- `pre_webhook_retry_count` - (Optional) The number of attempt to retry a failed webhook call
- `pre_webhook_url` - (Optional) The webhook url
- `webhook_filters` - (Optional) The events which trigger the webhook
- `webhook_method` - (Optional) The HTTP method to trigger the webhook. Valid values are `POST` or `GET`
- `reachability_enabled` - (Optional) Whether the Reachability Indicator (for Programmable Chat) is enabled
- `read_status_enabled` - (Optional) Whether the Message Consumption Horizon (for Programmable Chat) is enabled
- `typing_indicator_timeout` - (Optional) How many seconds should the service wait after receiving a `started typing` event before assuming a user is no longer typing

---

A `limits` block supports the following:

- `channel_members` - (Optional) The max number of members that can be added to a channel
- `user_channels` - (Optional) The max number of users that can be a member of

---

A `media` block supports the following:

- `compatibility_message` - (Optional) The placeholder message for media messages which has no text
  
---

A `notifications` block supports the following:

- `log_enabled` - (Optional) Whether notification logs are enabled
- `new_message` - (Optional) A `new_message` block as documented below.
- `added_to_channel` - (Optional) A `added_to_channel` block as documented below.
- `invited_to_channel` - (Optional) A `added_to_channel` block as documented below.
- `removed_from_channel` - (Optional) A `removed_from_channel` block as documented below.

---

A `new_message` block supports the following:

- `enabled` - (Optional) Whether notifications for a new message are enabled
- `template` - (Optional) The template message that is sent when a new message is received
- `sound` - (Optional) The sound played when the notification is activated
- `badge_count_enabled` - (Optional) Whether bade counts are enabled

---

A `added_to_channel` block supports the following:

- `enabled` - (Optional) Whether notifications for a user being added to a channel are enabled
- `template` - (Optional) The template message that is sent when the notification is activated
- `sound` - (Optional) The sound played when the notification is activated

---

A `invited_to_channel` block supports the following:

- `enabled` - (Optional) Whether notifications for a user being invited to a channel are enabled
- `template` - (Optional) The template message that is sent when the notification is activated
- `sound` - (Optional) The sound played when the notification is activated

---

A `removed_from_channel` block supports the following:

- `enabled` - (Optional) Whether notifications for a user being removed from a channel are enabled
- `template` - (Optional) The template message that is sent when the notification is activated
- `sound` - (Optional) The sound played when the notification is activated

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the SID)
- `sid` - The SID of the service (Same as the ID)
- `account_sid` - The Account SID associated with the service
- `default_channel_creator_role_sid` - The Role SID that is associated with a user when they join a new channel
- `default_channel_role_sid` - The Role SID that is associated with a user when they are added to a channel
- `default_service_role_sid` - The Role SID that is associated with a user when they are added to the service
- `friendly_name` - The Friendly Name of the Chat Service
- `limits` - A `limits` block as documented below.
- `media` - A `media` block as documented below.
- `notifications` - A `notifications` block as documented below.
- `post_webhook_retry_count` - The number of attempt to retry a failed webhook call
- `post_webhook_url` - The webhook url
- `pre_webhook_retry_count` - The number of attempt to retry a failed webhook call
- `pre_webhook_url` - The webhook url
- `webhook_filters` - The events which trigger the webhook
- `webhook_method` - The HTTP method to trigger the webhook
- `reachability_enabled` - Whether the Reachability Indicator (for Programmable Chat) is enabled
- `read_status_enabled` - Whether the Message Consumption Horizon (for Programmable Chat) is enabled
- `typing_indicator_timeout` - How many seconds should the service wait after receiving a `started typing` event before assuming a user is no longer typing
- `date_created` - The date that the service was created
- `date_updated` - The date that the service was updated
- `url` - The url of the service

---

A `limits` block supports the following:

- `channel_members` - he max number of members that can be added to a channel
- `user_channels` - The max number of users that can be a member of

---

A `media` block supports the following:

- `compatibility_message` - The placeholder message for media messages which has no text
- `size_limit_mb` - The media file size limit in Mb
  
---

A `notifications` block supports the following:

- `log_enabled` - Whether notification logs are enabled
- `new_message` - A `new_message` block as documented below.
- `added_to_channel` - A `added_to_channel` block as documented below.
- `invited_to_channel` - A `added_to_channel` block as documented below.
- `removed_from_channel` - A `removed_from_channel` block as documented below.

---

A `new_message` block supports the following:

- `enabled` - Whether notifications for a new message are enabled
- `template` - The template message that is sent when a new message is received
- `sound` - The sound played when the notification is activated
- `badge_count_enabled` - Whether bade counts are enabled

---

A `added_to_channel` block supports the following:

- `enabled` - Whether notifications for a user being added to a channel are enabled
- `template` - The template message that is sent when the notification is activated
- `sound` - The sound played when the notification is activated

---

A `invited_to_channel` block supports the following:

- `enabled` - Whether notifications for a user being invited to a channel are enabled
- `template` - The template message that is sent when the notification is activated
- `sound` - The sound played when the notification is activated

---

A `removed_from_channel` block supports the following:

- `enabled` - Whether notifications for a user being removed from a channel are enabled
- `template` - The template message that is sent when the notification is activated
- `sound` - The sound played when the notification is activated
