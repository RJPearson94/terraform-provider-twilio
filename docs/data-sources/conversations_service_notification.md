---
page_title: "Twilio Conversations Service Notification"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the notification configuration is associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the Service SID)
- `service_sid` - The service SID associated with the notification configuration (Same as ID)
- `new_message` - A `new_message` block as documented below.
- `added_to_conversation` - An `added_to_conversation` block as documented below.
- `removed_from_conversation` - A `removed_from_conversation` block as documented below.
- `log_enabled` - (Optional) Whether notification logging is enabled
- `url` - The URL of the notification configuration

---

A `new_message` block supports the following:

- `enabled` - Whether new message notifications are enabled
- `template` - The message template which is used when a new message is added to the conversation
- `sound` - The sound played when a new message is added to the conversation
- `badge_count_enabled` - Whether message badges are enabled for the conversation

---

An `added_to_conversation` block supports the following:

- `enabled` - Whether notifications for users being added to the conversation are enabled
- `template` - The message template which is used when a user is added to the conversation
- `sound` - The sound played when a user is added to the conversation

---

A `removed_from_conversation` block supports the following:

- `enabled` - Whether notifications for users being removed from the conversation are enabled
- `template` - The message template which is used when a user is removed from the conversation
- `sound` - The sound played when a user is removed from the conversation

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the notification configuration
