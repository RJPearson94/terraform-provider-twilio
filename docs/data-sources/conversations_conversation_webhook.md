---
page_title: "twilio_conversations_conversation_webhook Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_conversation_webhook Data Source

Use this data source to access information about an existing conversation webhook. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource) for more information

~> This is a generic data source which can be used to retrieve conversation webhook info regardless of the type (webhook, trigger, studio)

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_conversation_webhook" "webhook" {
  service_sid      = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  conversation_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid              = "WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhook" {
  value = data.twilio_conversations_conversation_webhook.webhook
}
```

## Schema

### Required

- `conversation_sid` (String) The SID of the conversation
- `service_sid` (String) The SID of the conversations service
- `sid` (String) The SID of the conversation webhook to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this conversation webhook
- `configuration` (List of Object) The configuration of the conversation webhook (see [below for nested schema](#nestedatt--configuration))
- `date_created` (String) The date and time the conversation webhook was created, in RFC 3339 format
- `date_updated` (String) The date and time the conversation webhook was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `target` (String) The target of the conversation webhook
- `url` (String) The absolute URL of the conversation webhook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `filters` (List of String)
- `flow_sid` (Number)
- `method` (String)
- `replay_after` (Number)
- `triggers` (List of String)
- `webhook_url` (String)
