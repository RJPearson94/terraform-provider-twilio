---
page_title: "Twilio Conversations Webhook"
subcategory: "Conversations"
---

# twilio_conversations_webhook Resource

Manages the webhook configuration for the conversation service. See the [API docs](https://www.twilio.com/docs/conversations/api/webhook-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversation webhook configuration for the account. No new resources will be provisioned. Instead, the webhook configuration will be updated upon creation and the webhook configuration will remain after the destruction of the resource.

## Example Usage

```hcl
resource "twilio_conversations_webhook" "webhook" {}
```

## Argument Reference

The following arguments are supported:

- `target` - (Optional) The target of webhook. Valid values are `webhook` or `flex`
- `method` - (Optional) The HTTP method to trigger the webhook. Valid values are `POST` or `GET`
- `pre_webhook_url` - (Optional) The pre webhook URL
- `post_webhook_url` - (Optional) The post webhook URL
- `filters` - (Optional) The filter conditions that trigger the webhook

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the account SID)
- `account_sid` - The account SID associated with the webhook (Same as the ID)
- `target` - The target of webhook
- `method` - The HTTP method to trigger the webhook
- `pre_webhook_url` - The pre webhook URL
- `post_webhook_url` - The post webhook URL
- `filters` - The filter conditions that trigger the webhook
- `url` - The URL of the webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `update` - (Defaults to 10 minutes) Used when updating the webhook
- `read` - (Defaults to 5 minutes) Used when retrieving the webhook
