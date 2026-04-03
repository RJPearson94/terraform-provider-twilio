---
page_title: "twilio_conversations_user Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_user Resource

Manages a conversation user. See the [API docs](https://www.twilio.com/docs/conversations/api/user-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "test-user"
}
```

## Schema

### Required

- `identity` (String) The unique identity string for the user. Changing this forces a new resource
- `service_sid` (String) The SID of the conversations service. Changing this forces a new resource

### Optional

- `attributes` (String) A JSON string of attributes associated with the user. Defaults to `{}`
- `friendly_name` (String) A human-readable label for the user
- `role_sid` (String) The SID of the role to assign to the user
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this user
- `date_created` (String) The date and time the user was created, in RFC 3339 format
- `date_updated` (String) The date and time the user was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `is_notifiable` (Boolean) Whether the user has a potentially valid push channel registration for notifications
- `is_online` (Boolean) Whether the user is actively connected to the service
- `sid` (String) The unique SID assigned to this user by Twilio
- `url` (String) The absolute URL of the user resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A user can be imported using the `/Services/{serviceSid}/Users/{sid}` format, e.g.

```shell
terraform import twilio_conversations_user.user /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
