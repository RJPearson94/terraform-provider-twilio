---
page_title: "twilio_conversations_service_configuration Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_service_configuration Resource

Manages configuration for a conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversations service configuration. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

!> Removing the `default_chat_service_role_sid`, `default_conversation_creator_role_sid` or `default_conversation_role_sid` from your configuration will cause the corresponding value to be retained after a Terraform apply. If you want to change any of the value you will need to either create a new `twilio_conversations_role` resource and set the argument to the generated `sid`. Alternatively, you can set the role sid to one of the roles that were created when the service was created

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_service_configuration" "configuration" {
  service_sid = twilio_conversations_service.service.sid
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service. Changing this forces a new resource

### Optional

- `default_chat_service_role_sid` (String) The SID of the default role assigned to users when they join the service
- `default_conversation_creator_role_sid` (String) The SID of the default role assigned to the creator of a conversation
- `default_conversation_role_sid` (String) The SID of the default role assigned to users when they are added to a conversation
- `reachability_enabled` (Boolean) Whether the reachability indicator is enabled for the service. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the service configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
- `update` (String)
