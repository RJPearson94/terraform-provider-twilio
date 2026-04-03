---
page_title: "twilio_conversations_user Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_user Data Source

Use this data source to access information about an existing conversations user. See the [API docs](https://www.twilio.com/docs/conversations/api/user-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_user" "user" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "user" {
  value = data.twilio_conversations_user.user
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service
- `sid` (String) The SID of the user to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this user
- `attributes` (String) A JSON string of attributes associated with the user
- `date_created` (String) The date and time the user was created, in RFC 3339 format
- `date_updated` (String) The date and time the user was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the user
- `id` (String) The ID of this resource.
- `identity` (String) The unique identity string for the user
- `is_notifiable` (Boolean) Whether the user has a potentially valid push channel registration for notifications
- `is_online` (Boolean) Whether the user is actively connected to the service
- `role_sid` (String) The SID of the role assigned to the user
- `url` (String) The absolute URL of the user resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
