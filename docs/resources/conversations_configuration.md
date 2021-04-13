---
page_title: "Twilio Conversations Configuration"
subcategory: "Conversations"
---

# twilio_conversations_configuration Resource

Manages configuration for the conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> This resource modifies the Twilio conversation configuration for the account. No new resources will be provisioned. Instead, the configuration will be updated upon creation and the configuration will remain after the destruction of the resource.

!> Removing the `default_service_sid`, `default_closed_timer`, `default_inactive_timer` or `default_messaging_service_sid` from your configuration will cause the corresponding value to be retained after a Terraform apply. If you want to change any of the value you will need to update your configuration to set an appropriate value

## Example Usage

```hcl
resource "twilio_conversations_configuration" "configuration" {}
```

## Argument Reference

The following arguments are supported:

- `default_service_sid` - (Optional) The default conversation service to associate newly created conversations with
- `default_closed_timer` - (Optional) The default ISO8601 duration before a conversation will be marked as closed
- `default_inactive_timer` - (Optional) The default ISO8601 duration before a conversation will be marked as inactive
- `default_messaging_service_sid` - (Optional) The default messaging service to associate newly created conversations with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `account_sid`)
- `account_sid` - The account SID associated with the configuration (Same as the `id`)
- `default_service_sid` - The default conversation service to associate newly created conversations with
- `default_closed_timer` - The default ISO8601 duration before a conversation will be marked as closed
- `default_inactive_timer` - The default ISO8601 duration before a conversation will be marked as inactive
- `default_messaging_service_sid` - The default messaging service to associate newly created conversations with
- `url` - The URL of the configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `update` - (Defaults to 10 minutes) Used when updating the service configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the service configuration
