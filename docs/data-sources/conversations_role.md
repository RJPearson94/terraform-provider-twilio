---
page_title: "twilio_conversations_role Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_role Data Source

Use this data source to access information about an existing conversations role. See the [API docs](https://www.twilio.com/docs/conversations/api/role-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_role" "role" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "role" {
  value = data.twilio_conversations_role.role
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service
- `sid` (String) The SID of the role to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this role
- `date_created` (String) The date and time the role was created, in RFC 3339 format
- `date_updated` (String) The date and time the role was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the role
- `id` (String) The ID of this resource.
- `permissions` (List of String) The list of permissions granted to the role
- `type` (String) The type of role
- `url` (String) The absolute URL of the role resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
