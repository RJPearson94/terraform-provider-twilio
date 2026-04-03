---
page_title: "twilio_conversations_role Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_role Resource

Manages a conversation role. See the [API docs](https://www.twilio.com/docs/conversations/api/role-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_role" "role" {
  service_sid   = twilio_conversations_service.service.sid
  friendly_name = "twilio-test-role"
  type          = "conversation"
  permissions = [
    "sendMediaMessage",
    "sendMessage"
  ]
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the role. Changing this forces a new resource
- `permissions` (List of String) The list of permissions granted to the role
- `service_sid` (String) The SID of the conversations service. Changing this forces a new resource
- `type` (String) The type of role. Valid values are `conversation` or `service`. Changing this forces a new resource

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
terraform import twilio_conversations_role.role /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
