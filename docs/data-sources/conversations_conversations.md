---
page_title: "twilio_conversations_conversations Data Source - twilio"
subcategory: "Conversations"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the conversations service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these conversations
- `conversations` (List of Object) The list of conversations (see [below for nested schema](#nestedatt--conversations))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--conversations"></a>
### Nested Schema for `conversations`

Read-Only:

- `attributes` (String)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `messaging_service_sid` (String)
- `sid` (String)
- `state` (String)
- `timers` (List of Object) (see [below for nested schema](#nestedobjatt--conversations--timers))
- `unique_name` (String)
- `url` (String)

<a id="nestedobjatt--conversations--timers"></a>
### Nested Schema for `conversations.timers`

Read-Only:

- `date_closed` (String)
- `date_inactive` (String)
