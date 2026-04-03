---
page_title: "twilio_conversations_users Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_users Data Source

Use this data source to access information about the users associated with an existing conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/user-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_users" "users" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "users" {
  value = data.twilio_conversations_users.users
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these users
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
- `role_sid` (String)
- `sid` (String)
- `url` (String)
