---
page_title: "twilio_chat_user Resource - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_user Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat user. See the [API docs](https://www.twilio.com/docs/chat/rest/user-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

!> If the `role_sid` is managed via Terraform and the `role_sid` is removed from the configuration file. The old value will be retained on the next apply.

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name = "twilio-test"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "twilio-test"
}
```

## Schema

### Required

- `identity` (String) The unique identity string of the user. Changing this forces a new resource
- `service_sid` (String) The SID of the Programmable Chat service. Changing this forces a new resource

### Optional

- `attributes` (String) A JSON string of custom attributes for the user. Defaults to `{}`
- `friendly_name` (String) A human-readable label for the user
- `role_sid` (String) The SID of the role to assign to the user
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this user
- `date_created` (String) The date and time the user was created, in RFC 3339 format
- `date_updated` (String) The date and time the user was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `is_notifiable` (Boolean) Whether the user has a potentially valid push registration for the chat service
- `is_online` (Boolean) Whether the user is actively connected to the chat service
- `joined_channels_count` (Number) The number of channels the user has joined
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
terraform import twilio_chat_role.role /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
