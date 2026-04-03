---
page_title: "twilio_chat_channel_webhook Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel_webhook Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat channel webhook. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-webhook-resource) for more information

~> This is a generic data source which can be used to retrieve channel webhook info regardless of the type (webhook, trigger, studio)

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_webhook" "webhook" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhook" {
  value = data.twilio_chat_channel_webhook.webhook
}
```

## Schema

### Required

- `channel_sid` (String) The SID of the chat channel
- `service_sid` (String) The SID of the Programmable Chat service
- `sid` (String) The SID of the channel webhook

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this channel webhook
- `configuration` (List of Object) The configuration of the channel webhook (see [below for nested schema](#nestedatt--configuration))
- `date_created` (String) The date and time the channel webhook was created, in RFC 3339 format
- `date_updated` (String) The date and time the channel webhook was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `type` (String) The type of the channel webhook
- `url` (String) The absolute URL of the channel webhook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `filters` (List of String)
- `flow_sid` (String)
- `method` (String)
- `retry_count` (Number)
- `triggers` (List of String)
- `webhook_url` (String)
