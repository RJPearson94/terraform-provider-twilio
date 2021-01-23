---
page_title: "Twilio Conversations Conversation Webhooks"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the webhooks are associated with
- `conversation_sid` - (Mandatory) The SID of the conversation the webhooks are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `service_sid/conversation_sid`
- `account_sid` - The SID of the account the conversation webhooks are associated with
- `service_sid` - The SID of the service the conversation webhooks are associated with
- `conversation_sid` - The SID of the conversation the webhooks are associated with
- `webhooks` - A list of `webhook` blocks as documented below

---

A `webhook` block supports the following:

- `sid` - The SID of the conversation webhook
- `type` - The type of webhook
- `configuration` - The `configuration` block as documented below
- `date_created` - The date in RFC3339 format that the conversation webhook was created
- `date_updated` - The date in RFC3339 format that the conversation webhook was updated
- `url` - The URL of the conversation webhook

---

A `configuration` block supports the following:

- `method` - The HTTP method to trigger the conversation webhook
- `webhook_url` - The webhook URL
- `filters` - The filter conditions that trigger the conversation webhook
- `replay_after` - Message Index to replay messages from
- `flow_sid` - The SID for the studio flow which will be called
- `triggers` - The keywords which trigger the conversation webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving conversation webhooks
