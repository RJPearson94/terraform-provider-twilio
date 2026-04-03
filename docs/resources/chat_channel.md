---
page_title: "twilio_chat_channel Resource - twilio"
subcategory: "Programmable Chat"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service. Changing this forces a new resource

### Optional

- `attributes` (String) A JSON string of custom attributes for the channel. Defaults to `{}`
- `friendly_name` (String) A human-readable label for the channel
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `type` (String) The visibility of the channel. Valid values are `public` or `private`. Defaults to `public`. Changing this forces a new resource
- `unique_name` (String) A unique name for the channel

### Read-Only

- `account_sid` (String) The SID of the account that owns this channel
- `created_by` (String) The identity of the user that created the channel
- `date_created` (String) The date and time the channel was created, in RFC 3339 format
- `date_updated` (String) The date and time the channel was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `members_count` (Number) The number of members in the channel
- `messages_count` (Number) The number of messages in the channel
- `sid` (String) The unique SID assigned to this channel by Twilio
- `url` (String) The absolute URL of the channel resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A channel can be imported using the `/Services/{serviceSid}/Channels/{sid}` format, e.g.

```shell
terraform import twilio_chat_channel.channel /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
