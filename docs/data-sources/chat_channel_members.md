---
page_title: "twilio_chat_channel_members Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_channel_member Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the members associated with an existing Programmable Chat service and channel. See the [API docs](https://www.twilio.com/docs/chat/rest/member-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_members" "members" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "members" {
  value = data.twilio_chat_channel_members.members
}
```

## Schema

### Required

- `channel_sid` (String) The SID of the chat channel
- `service_sid` (String) The SID of the Programmable Chat service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the channel members
- `id` (String) The ID of this resource.
- `members` (List of Object) The list of channel members (see [below for nested schema](#nestedatt--members))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `attributes` (String)
- `date_created` (String)
- `date_updated` (String)
- `identity` (String)
- `last_consumed_message_index` (Number)
- `last_consumption_timestamp` (String)
- `role_sid` (String)
- `sid` (String)
- `url` (String)
