---
page_title: "twilio_conversations_conversation_webhook Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_conversation_webhook Resource

Manages a conversation scoped webhook. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}

resource "twilio_conversations_conversation_webhook" "webhook" {
  service_sid      = twilio_conversations_service.service.sid
  conversation_sid = twilio_conversations_conversation.conversation.sid
  webhook_url      = "https://test.com/new"
  filters          = ["onMessageAdded"]
}
```

## Schema

### Required

- `conversation_sid` (String) The SID of the conversation. Changing this forces a new resource
- `filters` (List of String) The list of webhook event filters to subscribe to
- `service_sid` (String) The SID of the conversations service. Changing this forces a new resource
- `webhook_url` (String) The URL to send webhook requests to

### Optional

- `method` (String) The HTTP method for the webhook. Valid values are `GET` or `POST`. Defaults to `POST`
- `replay_after` (Number) The message index to replay messages from. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this conversation webhook
- `date_created` (String) The date and time the conversation webhook was created, in RFC 3339 format
- `date_updated` (String) The date and time the conversation webhook was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this conversation webhook by Twilio
- `target` (String) The target of the conversation webhook
- `url` (String) The absolute URL of the conversation webhook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A conversation webhook can be imported using the `/Services/{serviceSid}/Conversations/{conversationSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_conversations_conversation_webhook.webhook /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
