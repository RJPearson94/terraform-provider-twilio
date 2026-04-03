---
page_title: "twilio_conversations_conversation_webhooks Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_conversation_webhooks Data Source

Use this data source to access information about the webhooks associated with an existing conversation service and conversation. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource) for more information

~> This is a generic data source which can be used to retrieve conversation webhooks regardless of there type (webhook, trigger, studio)

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_conversation_webhook" "webhooks" {
  service_sid      = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  conversation_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhooks" {
  value = data.twilio_conversations_conversation_webhook.webhooks
}
```

## Schema

### Required

- `conversation_sid` (String) The SID of the conversation
- `service_sid` (String) The SID of the conversations service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these conversation webhooks
- `id` (String) The ID of this resource.
- `webhooks` (List of Object) The list of conversation webhooks (see [below for nested schema](#nestedatt--webhooks))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--webhooks"></a>
### Nested Schema for `webhooks`

Read-Only:

- `configuration` (List of Object) (see [below for nested schema](#nestedobjatt--webhooks--configuration))
- `date_created` (String)
- `date_updated` (String)
- `sid` (String)
- `target` (String)
- `url` (String)

<a id="nestedobjatt--webhooks--configuration"></a>
### Nested Schema for `webhooks.configuration`

Read-Only:

- `filters` (List of String)
- `flow_sid` (Number)
- `method` (String)
- `replay_after` (Number)
- `triggers` (List of String)
- `webhook_url` (String)
