---
page_title: "twilio_conversations_webhook Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_webhook Resource

Manages the webhook configuration for the conversation service. See the [API docs](https://www.twilio.com/docs/conversations/api/webhook-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversation webhook configuration for the account. No new resources will be provisioned. Instead, the webhook configuration will be updated upon creation and the webhook configuration will remain after the destruction of the resource.

!> Removing any argument from your configuration will cause the corresponding value to be retained after a Terraform apply. If you want to change any of the values you will need to set the argument to the appropiate value

## Example Usage

```hcl
resource "twilio_conversations_webhook" "webhook" {}
```

## Schema

### Optional

- `filters` (List of String) The list of webhook event triggers to subscribe to
- `method` (String) The HTTP method for the webhook. Valid values are `GET` or `POST`
- `post_webhook_url` (String) The URL called after an event is sent to the webhook
- `pre_webhook_url` (String) The URL called before an event is sent to the webhook
- `target` (String) The target of the webhook. Valid values are `webhook` or `flex`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this webhook
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the webhook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
- `update` (String)
