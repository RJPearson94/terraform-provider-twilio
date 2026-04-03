---
page_title: "twilio_chat_channel_webhooks Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel_webhooks Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the webhooks associated with an existing Programmable Chat service and channel. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-webhook-resource) for more information

~> This is a generic data source which can be used to retrieve channel webhooks regardless of there type (webhook, trigger, studio)

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_webhooks" "webhooks" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhooks" {
  value = data.twilio_chat_channel_webhooks.webhooks
}
```

## Schema

### Required

- `channel_sid` (String) The SID of the chat channel
- `service_sid` (String) The SID of the Programmable Chat service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the channel webhooks
- `id` (String) The ID of this resource.
- `webhooks` (List of Object) The list of channel webhooks (see [below for nested schema](#nestedatt--webhooks))

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
- `type` (String)
- `url` (String)

<a id="nestedobjatt--webhooks--configuration"></a>
### Nested Schema for `webhooks.configuration`

Read-Only:

- `filters` (List of String)
- `flow_sid` (String)
- `method` (String)
- `retry_count` (Number)
- `triggers` (List of String)
- `webhook_url` (String)
