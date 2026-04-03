---
page_title: "twilio_chat_channel_member Resource - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel_member Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat channel member. See the [API docs](https://www.twilio.com/docs/chat/rest/member-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

!> If the `role_sid` is managed via Terraform and the `role_sid` is removed from the configuration file. The old value will be retained on the next apply.

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

## Schema

### Required

- `channel_sid` (String) The SID of the chat channel. Changing this forces a new resource
- `identity` (String) The unique identity string of the member. Changing this forces a new resource
- `service_sid` (String) The SID of the Programmable Chat service. Changing this forces a new resource

### Optional

- `attributes` (String) A JSON string of custom attributes for the channel member. Defaults to `{}`
- `role_sid` (String) The SID of the role to assign to the channel member
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this channel member
- `date_created` (String) The date and time the channel member was created, in RFC 3339 format
- `date_updated` (String) The date and time the channel member was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `last_consumed_message_index` (Number) The index of the last message the member has read in the channel
- `last_consumption_timestamp` (String) The date and time the member last consumed a message, in RFC 3339 format
- `sid` (String) The unique SID assigned to this channel member by Twilio
- `url` (String) The absolute URL of the channel member resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A channel member can be imported using the `/Services/{serviceSid}/Channels/{channelSid}/Members/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel_member.member /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Members/MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
