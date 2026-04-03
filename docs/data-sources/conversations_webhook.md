---
page_title: "twilio_conversations_webhook Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_webhook Data Source

Use this data source to access information about conversations webhook. See the [API docs](https://www.twilio.com/docs/conversations/api/webhook-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_webhook" "webhook" {}

output "webhook" {
  value = data.twilio_conversations_webhook.webhook
}
```

## Schema

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this webhook
- `filters` (List of String) The list of webhook event triggers subscribed to
- `id` (String) The ID of this resource.
- `method` (String) The HTTP method for the webhook
- `post_webhook_url` (String) The URL called after an event is sent to the webhook
- `pre_webhook_url` (String) The URL called before an event is sent to the webhook
- `target` (String) The target of the webhook
- `url` (String) The absolute URL of the webhook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
