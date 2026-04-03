---
page_title: "twilio_chat_channel_webhook Resource - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel_webhook Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat channel webhook. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-webhook-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-channel"
}

resource "twilio_chat_channel_webhook" "webhook" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  webhook_url = "https://test.com/new"
  filters     = ["onMessageSent"]
}
```

## Schema

### Required

- `channel_sid` (String) The SID of the chat channel. Changing this forces a new resource
- `filters` (List of String) The list of events that trigger the webhook
- `service_sid` (String) The SID of the Programmable Chat service. Changing this forces a new resource
- `webhook_url` (String) The URL to send webhook requests to

### Optional

- `method` (String) The HTTP method used for webhook requests. Valid values are `GET` or `POST`. Defaults to `POST`
- `retry_count` (Number) The number of retry attempts for failed webhook requests. Defaults to `0`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this channel webhook
- `date_created` (String) The date and time the channel webhook was created, in RFC 3339 format
- `date_updated` (String) The date and time the channel webhook was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this channel webhook by Twilio
- `type` (String) The type of the channel webhook
- `url` (String) The absolute URL of the channel webhook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A channel webhook can be imported using the `/Services/{serviceSid}/Channels/{channelSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel_webhook.webhook /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
