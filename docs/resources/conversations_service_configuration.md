---
page_title: "Twilio Conversations Service Configuration"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the configuration with. Changing this forces a new resource to be created
- `default_chat_service_role_sid` - (Optional) The default role to assign users when they are added to the service
- `default_conversation_creator_role_sid` - (Optional) The default role to assign creator users when they join a new conversation
- `default_conversation_role_sid` - (Optional) The default role to assign users when they join a new conversation
- `reachability_enabled` - (Optional) Whether Programmable Chat's reachability indicator is enabled or not. The default value is `false`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `service_sid`)
- `service_sid` - The service SID associated with the configuration (Same as the `id`)
- `default_chat_service_role_sid` - The default role assigned to users when they are added to the service
- `default_conversation_creator_role_sid` - The default role assigned to creator users when they join a new conversation
- `default_conversation_role_sid` - The default role assigned to users when they join a new conversation
- `reachability_enabled` - Whether Programmable Chat's reachability indicator is enabled or not
- `url` - The URL of the service configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `update` - (Defaults to 10 minutes) Used when updating the service configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the service configuration
