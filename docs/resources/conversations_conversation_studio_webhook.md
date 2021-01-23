---
page_title: "Twilio Conversations Conversation Webhook (Studio)"
subcategory: "Conversations"
---

# twilio_conversations_conversation_studio_webhook Resource

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

resource "twilio_conversations_conversation_studio_webhook" "studio_webhook" {
  service_sid      = twilio_conversations_service.service.sid
  conversation_sid = twilio_conversations_conversation.conversation.sid
  flow_sid    = "FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the conversation with. Changing this forces a new resource to be created
- `conversation_sid` - (Mandatory) The conversation SID to associate the webhook with. Changing this forces a new resource to be created
- `flow_sid` - (Mandatory) The SID for the Studio flow which will be called
- `replay_after` - (Optional) Message Index to replay messages from

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the conversation webhook (Same as the SID)
- `sid` - The SID of the conversation webhook (Same as the ID)
- `account_sid` - The account SID associated with the conversation webhook
- `service_sid` - The service SID associated with the conversation webhook
- `conversation_sid` - The conversation SID to associate the webhook with
- `target` - The target of webhook. The value will be `studio`
- `flow_sid` - The SID for the studio flow which will be called
- `replay_after` - Message Index to replay messages from
- `date_created` - The date in RFC3339 format that the conversation webhook was created
- `date_updated` - The date in RFC3339 format that the conversation webhook was updated
- `url` - The URL of the conversation webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the conversation webhook
- `update` - (Defaults to 10 minutes) Used when updating the conversation webhook
- `read` - (Defaults to 5 minutes) Used when retrieving the conversation webhook
- `delete` - (Defaults to 10 minutes) Used when deleting the conversation webhook

## Import

A conversation webhook can be imported using the `/Services/{serviceSid}/Conversations/{conversationSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_conversations_conversation_studio_webhook.studio_webhook /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
