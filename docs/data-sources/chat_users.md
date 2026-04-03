---
page_title: "twilio_chat_users Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_users Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the users associated with an existing Programmable Chat service. See the [API docs](https://www.twilio.com/docs/chat/rest/user-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_users" "roles" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "users" {
  value = data.twilio_chat_users.users
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the users
- `id` (String) The ID of this resource.
- `users` (List of Object) The list of users (see [below for nested schema](#nestedatt--users))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--users"></a>
### Nested Schema for `users`

Read-Only:

- `attributes` (String)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `identity` (String)
- `is_notifiable` (Boolean)
- `is_online` (Boolean)
- `joined_channels_count` (Number)
- `role_sid` (String)
- `sid` (String)
- `url` (String)
