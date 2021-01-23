---
page_title: "Twilio Conversations Conversations"
subcategory: "Conversations"
---

# twilio_conversations_conversations Data Source

Use this data source to access information about the conversations associated with an existing conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_conversations" "conversations" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "conversations" {
  value = data.twilio_conversations_conversations.conversations
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the conversations are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `account_sid` - The SID of the account the conversations are associated with
- `service_sid` - The SID of the service the conversations are associated with
- `conversations` - A list of `conversation` blocks as documented below

---

A `conversation` block supports the following:

- `sid` - The SID of the conversation
- `unique_name` - The unique name of the conversation
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

- `read` - (Defaults to 10 minutes) Used when retrieving conversations
