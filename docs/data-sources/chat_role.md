---
page_title: "twilio_chat_role Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_role Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat role. See the [API docs](https://www.twilio.com/docs/chat/rest/role-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_role" "role" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "role" {
  value = data.twilio_chat_role.role
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service
- `sid` (String) The SID of the role

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this role
- `date_created` (String) The date and time the role was created, in RFC 3339 format
- `date_updated` (String) The date and time the role was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the role
- `id` (String) The ID of this resource.
- `permissions` (List of String) The list of permissions granted to the role
- `type` (String) The type of role. Values are `channel` or `deployment`
- `url` (String) The absolute URL of the role resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
