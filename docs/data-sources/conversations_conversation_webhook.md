---
page_title: "Twilio Conversations Conversation Webhook"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the webhook is associated with
- `conversation_sid` - (Mandatory) The SID of the conversation the webhook is associated with
- `sid` - (Mandatory) The SID of the webhook

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the conversation webhook (Same as the SID)
- `sid` - The SID of the conversation webhook (Same as the ID)
- `account_sid` - The account SID associated with the conversation webhook
- `service_sid` - The service SID associated with the conversation webhook
- `conversation_sid` - The conversation SID to associate the webhook with
- `target` - The target of webhook
- `configuration` - The `configuration` block as documented below
- `date_created` - The date in RFC3339 format that the conversation webhook was created
- `date_updated` - The date in RFC3339 format that the conversation webhook was updated
- `url` - The URL of the conversation webhook

---

A `configuration` block supports the following:

- `method` - The HTTP method to trigger the conversation webhook
- `webhook_url` - The webhook URL
- `filters` - The filter conditions that trigger the conversation webhook
- `flow_sid` - The SID for the studio flow which will be called
- `replay_after` - Message Index to replay messages from
- `triggers` - The keywords which trigger the conversation webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the webhook webhook
