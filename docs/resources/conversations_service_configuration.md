---
page_title: "Twilio Conversations Service Configuration"
subcategory: "Conversations"
---

# twilio_conversations_service_configuration Resource

Manages configuration for a conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversations service configuration. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_service_configuration" "configuration" {
  service_sid = twilio_conversations_service.service.sid
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the configuration with. Changing this forces a new resource to be created
- `default_chat_service_role_sid` - (Optional) The default role to assign users when they are added to the service
- `default_conversation_creator_role_sid` - (Optional) The default role to assign creator users when they join a new conversation
- `default_conversation_role_sid` - (Optional) The default role to assign users when they join a new conversation
- `reachability_enabled` - (Optional) Whether Programmable Chat's reachability indicator is enabled or not

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the Service SID)
- `service_sid` - The service SID associated with the configuration (Same as ID)
- `default_chat_service_role_sid` - The default role assigned to users when they are added to the service
- `default_conversation_creator_role_sid` - The default role assigned to creator users when they join a new conversation
- `default_conversation_role_sid` - The default role assigned to users when they join a new conversation
- `reachability_enabled` - Whether Programmable Chat's reachability indicator is enabled or not
- `url` - The URL of the service configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `update` - (Defaults to 10 minutes) Used when updating the service configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the service configuration
