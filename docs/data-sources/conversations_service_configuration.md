---
page_title: "Twilio Conversations Service Configuration"
subcategory: "Conversations"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the configuration is associated with

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

- `read` - (Defaults to 5 minutes) Used when retrieving the service configuration
