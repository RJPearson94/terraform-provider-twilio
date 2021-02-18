---
page_title: "Twilio Conversations Conversation"
subcategory: "Conversations"
---

# twilio_conversations_conversation Data Source

Use this data source to access information about an existing conversations conversation. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_conversation" "conversation" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "conversation" {
  value = data.twilio_conversations_conversation.conversation
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the conversation is associated with
- `sid` - (Mandatory) The SID of the conversation

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the conversation (Same as the `sid`)
- `sid` - The SID of the conversation (Same as the `id`)
- `account_sid` - The account SID associated with the conversation
- `unique_name` - The unique name of the conversation
- `service_sid` - The service SID associated with the conversation
- `friendly_name` - The friendly name of the conversation
- `attributes` - JSON string of attributes
- `messaging_service_sid` - The messaging service SID associated with the conversation
- `state` - The state of the conversation
- `timers` - A `timer` block as documented below
- `date_created` - The date in RFC3339 format that the conversation was created
- `date_updated` - The date in RFC3339 format that the conversation was updated
- `url` - The URL of the conversation

---

A `timer` block supports the following:

- `date_closed` - The date in RFC3339 format that the conversation will close
- `date_inactive` - The date in RFC3339 format that the conversation will be marked as inactive

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the conversation
