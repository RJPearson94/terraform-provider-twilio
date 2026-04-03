---
page_title: "twilio_chat_user Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_user Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat user. See the [API docs](https://www.twilio.com/docs/chat/rest/user-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_user" "role" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "user" {
  value = data.twilio_chat_user.user
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service
- `sid` (String) The SID of the user

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this user
- `attributes` (String) A JSON string of custom attributes for the user
- `date_created` (String) The date and time the user was created, in RFC 3339 format
- `date_updated` (String) The date and time the user was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the user
- `id` (String) The ID of this resource.
- `identity` (String) The unique identity string of the user
- `is_notifiable` (Boolean) Whether the user has a potentially valid push registration for the chat service
- `is_online` (Boolean) Whether the user is actively connected to the chat service
- `joined_channels_count` (Number) The number of channels the user has joined
- `role_sid` (String) The SID of the role assigned to the user
- `url` (String) The absolute URL of the user resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
