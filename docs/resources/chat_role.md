---
page_title: "twilio_chat_role Resource - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_role Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat role. See the [API docs](https://www.twilio.com/docs/chat/rest/role-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name = "twilio-test"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-role"
  type          = "channel"
  permissions = [
    "sendMessage",
    "leaveChannel"
  ]
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the role. Changing this forces a new resource
- `permissions` (List of String) The list of permissions granted to the role
- `service_sid` (String) The SID of the Programmable Chat service. Changing this forces a new resource
- `type` (String) The type of role. Valid values are `channel` or `deployment`. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this role
- `date_created` (String) The date and time the role was created, in RFC 3339 format
- `date_updated` (String) The date and time the role was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this role by Twilio
- `url` (String) The absolute URL of the role resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A role can be imported using the `/Services/{serviceSid}/Roles/{sid}` format, e.g.

```shell
terraform import twilio_chat_role.role /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
