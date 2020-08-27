---
page_title: "Twilio Programmable Chat Channel Webhook"
subcategory: "Programmable Chat"
---

# twilio_chat_channel_webhook Resource

Manages a Chat Channel Webhook

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
  webhook_url = "http://localhost/new"
  filters     = ["onMessageSent"]
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID associated with the channel webhook. Changing this forces a new resource to be created
- `channel_sid` - (Mandatory) The Channel SID associated with the channel webhook. Changing this forces a new resource to be created
- `method` - (Optional) The HTTP method to trigger the webhook. Valid values are `POST` or `GET`
- `webhook_url` - (Mandatory) The webhook url
- `filters` - (Mandatory) The filter conditions that trigger the webhook
- `retry_count` - (Optional) The number of attempt to retry a failed webhook call

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel webhook (Same as the SID)
- `sid` - The SID of the channel webhook (Same as the ID)
- `account_sid` - The Account SID associated with the channel webhook
- `service_sid` - The Service SID associated with the channel webhook
- `channel_sid` - The Channel SID associated with the channel webhook
- `type` - The type of webhook. The value will be webhook
- `method` - The HTTP method to trigger the webhook
- `webhook_url` - The webhook url
- `filters` - The filter conditions that trigger the webhook
- `retry_count` - The number of attempt to retry a failed webhook call
- `date_created` - The date in RFC3339 format that the channel webhook was created
- `date_updated` - The date in RFC3339 format that the channel webhook was updated
- `url` - The url of the channel webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the channel webhook
- `update` - (Defaults to 10 minutes) Used when updating the channel webhook
- `read` - (Defaults to 5 minutes) Used when retrieving the channel webhook
- `delete` - (Defaults to 10 minutes) Used when deleting the channel webhook

## Import

A channel webhook can be imported using the `/Services/{serviceSid}/Channels/{channelSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel_webhook.webhook /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
