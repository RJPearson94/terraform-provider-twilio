---
page_title: "twilio_chat_channel_member Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel_member Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat channel member. See the [API docs](https://www.twilio.com/docs/chat/rest/member-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_member" "member" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "member" {
  value = data.twilio_chat_channel_member.member
}
```

## Schema

### Required

- `channel_sid` (String) The SID of the chat channel
- `service_sid` (String) The SID of the Programmable Chat service
- `sid` (String) The SID of the channel member

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this channel member
- `attributes` (String) A JSON string of custom attributes for the channel member
- `date_created` (String) The date and time the channel member was created, in RFC 3339 format
- `date_updated` (String) The date and time the channel member was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `identity` (String) The unique identity string of the member
- `last_consumed_message_index` (Number) The index of the last message the member has read in the channel
- `last_consumption_timestamp` (String) The date and time the member last consumed a message, in RFC 3339 format
- `role_sid` (String) The SID of the role assigned to the channel member
- `url` (String) The absolute URL of the channel member resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
