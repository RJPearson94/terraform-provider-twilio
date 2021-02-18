---
page_title: "Twilio Programmable Chat Channel Member"
subcategory: "Programmable Chat"
---

# twilio_chat_channel_member Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat channel member. See the [API docs](https://www.twilio.com/docs/chat/rest/member-resource) for more information

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

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "twilio-test-user"
}

resource "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  identity    = twilio_chat_user.user.identity
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the channel member with. Changing this forces a new resource to be created
- `channel_sid` - (Mandatory) The channel SID to associate the channel member with. Changing this forces a new resource to be created
- `identity` - (Mandatory) The identity of the chat user. Changing this forces a new resource to be created
- `attributes` - (Optional) JSON string of member attributes
- `role_sid` - (Optional) The role SID assignment to the member

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel member (Same as the `sid`)
- `sid` - The SID of the channel member (Same as the `id`)
- `account_sid` - The account SID associated with the channel member
- `service_sid` - The service SID associated with the channel member
- `channel_sid` - The channel SID associated with the channel member
- `identity` - The identity of the chat member
- `attributes` - JSON string of member attributes
- `role_sid` - The role SID assignment to the member
- `last_consumed_message_index` - The index of the last message read by the member
- `last_consumption_timestamp` - The timestamp of the last message read by the member
- `date_created` - The date in RFC3339 format that the channel member was created
- `date_updated` - The date in RFC3339 format that the channel member was updated
- `url` - The URL of the channel member

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the channel member
- `update` - (Defaults to 10 minutes) Used when updating the channel member
- `read` - (Defaults to 5 minutes) Used when retrieving the channel member
- `delete` - (Defaults to 10 minutes) Used when deleting the channel member

## Import

A channel member can be imported using the `/Services/{serviceSid}/Channels/{channelSid}/Members/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel_member.member /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
