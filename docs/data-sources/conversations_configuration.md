---
page_title: "Twilio Conversations Configuration"
subcategory: "Conversations"
---

# twilio_conversations_configuration Data Source

Use this data source to access information about conversations configuration. See the [API docs](https://www.twilio.com/docs/conversations/api/configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_configuration" "configuration" {}

output "configuration" {
  value = data.twilio_conversations_configuration.configuration
}
```

## Argument Reference

The following arguments are supported:

N/A - This data source has no arguments

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the account SID)
- `account_sid` - The account SID associated with the configuration (Same as the ID)
- `default_service_sid` - The default conversation service to associate newly created conversations with
- `default_closed_timer` - The default ISO8601 duration before a conversation will be marked as closed
- `default_inactive_timer` - The default ISO8601 duration before a conversation will be marked as inactive
- `default_messaging_service_sid` - The default messaging service to associate newly created conversations with
- `url` - The URL of the configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the configuration
