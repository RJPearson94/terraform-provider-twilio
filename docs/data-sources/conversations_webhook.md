---
page_title: "Twilio Conversations Webhook"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

N/A - This data source has no arguments

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `account_sid`)
- `account_sid` - The account SID associated with the webhook (Same as the `id`)
- `target` - The target of webhook. The value will be webhook
- `method` - The HTTP method to trigger the webhook
- `pre_webhook_url` - The pre webhook URL
- `post_webhook_url` - The post webhook URL
- `filters` - The filter conditions that trigger the webhook
- `url` - The URL of the webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the webhook
