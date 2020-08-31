---
page_title: "Twilio Programmable Chat Service"
subcategory: "Programmable Chat"
---

# twilio_chat_service Data Source

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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the service

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the SID)
- `sid` - The SID of the service (Same as the ID)
- `account_sid` - The account SID associated with the service
- `default_channel_creator_role_sid` - The role SID that is associated with a user when they join a new channel
- `default_channel_role_sid` - The role SID that is associated with a user when they are added to a channel
- `default_service_role_sid` - The role SID that is associated with a user when they are added to the service
- `friendly_name` - The friendly name of the service
- `limits` - A `limits` block as documented below.
- `media` - A `media` block as documented below.
- `notifications` - A `notifications` block as documented below.
- `post_webhook_retry_count` - The number of attempts to retry a failed webhook call
- `post_webhook_url` - The webhook URL
- `pre_webhook_retry_count` - The number of attempts to retry a failed webhook call
- `pre_webhook_url` - The webhook URL
- `webhook_filters` - The events which trigger the webhook
- `webhook_method` - The HTTP method to trigger the webhook
- `reachability_enabled` - Whether the reachability indicator (for Programmable Chat) is enabled
- `read_status_enabled` - Whether the message consumption horizon (for Programmable Chat) is enabled
- `typing_indicator_timeout` - How many seconds should the service wait after receiving a `started typing` event before assuming a user is no longer typing
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

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

An `added_to_channel` block supports the following:

- `enabled` - Whether notifications for a user being added to a channel are enabled
- `template` - The template message that is sent when the notification is activated
- `sound` - The sound played when the notification is activated

---

An `invited_to_channel` block supports the following:

- `enabled` - Whether notifications for a user being invited to a channel are enabled
- `template` - The template message that is sent when the notification is activated
- `sound` - The sound played when the notification is activated

---

A `removed_from_channel` block supports the following:

- `enabled` - Whether notifications for a user being removed from a channel are enabled
- `template` - The template message that is sent when the notification is activated
- `sound` - The sound played when the notification is activated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the role
