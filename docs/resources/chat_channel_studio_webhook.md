---
page_title: "Twilio Programmable Chat Channel Webhook (Studio)"
subcategory: "Programmable Chat"
---

# twilio_chat_channel_studio_webhook Resource

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

resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID to associate the channel webhook with. Changing this forces a new resource to be created
- `channel_sid` - (Mandatory) The Channel SID to associate the channel webhook with. Changing this forces a new resource to be created
- `flow_sid` - (Mandatory) The SID for the Studio flow which will be called
- `retry_count` - (Optional) The number of attempts to retry a failed webhook call

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel webhook (Same as the SID)
- `sid` - The SID of the channel webhook (Same as the ID)
- `account_sid` - The account SID associated with the channel webhook
- `service_sid` - The service SID associated with the channel webhook
- `channel_sid` - The channel SID associated with the channel webhook
- `type` - The type of webhook. The value will be `studio`
- `flow_sid` - The SID for the studio flow which will be called
- `retry_count` - The number of attempts to retry a failed webhook call
- `date_created` - The date in RFC3339 format that the channel webhook was created
- `date_updated` - The date in RFC3339 format that the channel webhook was updated
- `url` - The URL of the channel webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the channel webhook
- `update` - (Defaults to 10 minutes) Used when updating the channel webhook
- `read` - (Defaults to 5 minutes) Used when retrieving the channel webhook
- `delete` - (Defaults to 10 minutes) Used when deleting the channel webhook

## Import

A channel webhook can be imported using the `/Services/{serviceSid}/Channels/{channelSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel_studio_webhook.webhook /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
