---
page_title: "Twilio Programmable Chat Channel"
subcategory: "Programmable Chat"
---

# twilio_chat_channel Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat channel. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name = "twilio-test"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-channel"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID to associate the channel with. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the channel
- `unique_name` - (Optional) The unique name of the channel
- `attributes` - (Optional) JSON string of channel attributes. The default value is `{}`
- `type` - (Optional) The type of channel. Valid values are `public` or `private`. The default value is `public`. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel (Same as the `sid`)
- `sid` - The SID of the channel (Same as the `id`)
- `account_sid` - The account SID associated with the channel
- `service_sid` - The service SID associated with the channel
- `friendly_name` - The friendly name of the channel
- `unique_name` - The unique name of the channel
- `attributes` - JSON string of channel attributes
- `type` - The type of channel
- `created_by` - Who created the chat channel
- `members_count` - The number of members that are associated with the channel
- `messages_count` - The number of messages that are associated with the channel
- `date_created` - The date in RFC3339 format that the channel was created
- `date_updated` - The date in RFC3339 format that the channel was updated
- `url` - The URL of the channel

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the channel
- `update` - (Defaults to 10 minutes) Used when updating the channel
- `read` - (Defaults to 5 minutes) Used when retrieving the channel
- `delete` - (Defaults to 10 minutes) Used when deleting the channel

## Import

A channel can be imported using the `/Services/{serviceSid}/Channels/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel.channel /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
