---
page_title: "twilio_conversations_conversation Data Source - twilio"
subcategory: "Conversations"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the conversations service
- `sid` (String) The SID of the conversation to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this conversation
- `attributes` (String) A JSON string of attributes associated with the conversation
- `date_created` (String) The date and time the conversation was created, in RFC 3339 format
- `date_updated` (String) The date and time the conversation was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the conversation
- `id` (String) The ID of this resource.
- `messaging_service_sid` (String) The SID of the messaging service associated with the conversation
- `state` (String) The state of the conversation
- `timers` (List of Object) Timer settings for the conversation (see [below for nested schema](#nestedatt--timers))
- `unique_name` (String) A unique, developer-assigned name for the conversation
- `url` (String) The absolute URL of the conversation resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--timers"></a>
### Nested Schema for `timers`

Read-Only:

- `date_closed` (String)
- `date_inactive` (String)
