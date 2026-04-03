---
page_title: "twilio_conversations_roles Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_roles Data Source

Use this data source to access information about the roles associated with an existing conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/role-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_roles" "roles" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "roles" {
  value = data.twilio_conversations_roles.roles
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these roles
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
