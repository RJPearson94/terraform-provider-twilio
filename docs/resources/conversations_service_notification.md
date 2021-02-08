---
page_title: "Twilio Conversations Service Notification"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the notification configuration with. Changing this forces a new resource to be created
- `new_message` - (Optional) A `new_message` block as documented below.
- `added_to_conversation` - (Optional) An `added_to_conversation` block as documented below.
- `removed_from_conversation` - (Optional) A `removed_from_conversation` block as documented below.
- `log_enabled` - (Optional) Whether notification logging is enabled

---

A `new_message` block supports the following:

- `enabled` - (Optional) Whether new message notifications are enabled
- `template` - (Optional) The message template which is used when a new message is added to the conversation
- `sound` - (Optional) The sound played when a new message is added to the conversation
- `badge_count_enabled` - (Optional) Whether message badges are enabled for the conversation

---

An `added_to_conversation` block supports the following:

- `enabled` - (Optional) Whether notifications for users being added to the conversation are enabled
- `template` - (Optional) The message template which is used when a user is added to the conversation
- `sound` - (Optional) The sound played when a user is added to the conversation

---

A `removed_from_conversation` block supports the following:

- `enabled` - (Optional) Whether notifications for users being removed from the conversation are enabled
- `template` - (Optional) The message template which is used when a user is removed from the conversation
- `sound` - (Optional) The sound played when a user is removed from the conversation

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

- `update` - (Defaults to 10 minutes) Used when updating the notification configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the notification configuration
