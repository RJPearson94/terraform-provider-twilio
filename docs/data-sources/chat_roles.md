---
page_title: "twilio_chat_roles Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
---

# twilio_chat_roles Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the roles associated with an existing Programmable Chat service. See the [API docs](https://www.twilio.com/docs/chat/rest/role-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_roles" "roles" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "roles" {
  value = data.twilio_chat_roles.roles
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the roles
- `id` (String) The ID of this resource.
- `roles` (List of Object) The list of roles (see [below for nested schema](#nestedatt--roles))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--roles"></a>
### Nested Schema for `roles`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `permissions` (List of String)
- `sid` (String)
- `type` (String)
- `url` (String)
