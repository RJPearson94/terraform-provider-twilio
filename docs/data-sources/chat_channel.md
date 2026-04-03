---
page_title: "twilio_chat_channel Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat channel. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel" "channel" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "channel" {
  value = data.twilio_chat_channel.channel
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service
- `sid` (String) The SID of the channel

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this channel
- `attributes` (String) A JSON string of custom attributes for the channel
- `created_by` (String) The identity of the user that created the channel
- `date_created` (String) The date and time the channel was created, in RFC 3339 format
- `date_updated` (String) The date and time the channel was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the channel
- `id` (String) The ID of this resource.
- `members_count` (Number) The number of members in the channel
- `messages_count` (Number) The number of messages in the channel
- `type` (String) The visibility of the channel. Values are `public` or `private`
- `unique_name` (String) A unique name for the channel
- `url` (String) The absolute URL of the channel resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
