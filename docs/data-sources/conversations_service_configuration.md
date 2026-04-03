---
page_title: "twilio_conversations_service_configuration Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_service_configuration Data Source

Use this data source to access configuration for a conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_service_configuration" "configuration" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "configuration" {
  value = data.twilio_conversations_service_configuration.configuration
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `default_chat_service_role_sid` (String) The SID of the default role assigned to users when they join the service
- `default_conversation_creator_role_sid` (String) The SID of the default role assigned to the creator of a conversation
- `default_conversation_role_sid` (String) The SID of the default role assigned to users when they are added to a conversation
- `id` (String) The ID of this resource.
- `reachability_enabled` (Boolean) Whether the reachability indicator is enabled for the service
- `url` (String) The absolute URL of the service configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
