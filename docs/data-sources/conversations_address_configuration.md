---
page_title: "Twilio Conversations Address Configuration"
subcategory: "Conversations"
---

# twilio_conversations_address_configuration Data Source

Use this data source to access information about address configuration. See the [API docs](https://www.twilio.com/docs/conversations/api/configuration-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_address_configuration" "address_configuration" {
  sid = "IGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "address_configuration" {
  value = data.twilio_conversations_address_configuration.address_configuration
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the address configuration

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the address configuration (Same as the `sid`)
- `sid` - The SID of the address configuration (Same as the `id`)
- `account_sid` - The account SID associated with the address configuration
- `auto_creation` - A `auto_creation` block as documented below
- `address` - The phone number or whatsapp number that has been configured
- `type` - The address type
- `date_created` - The date in RFC3339 format that the address configuration was created
- `date_updated` - The date in RFC3339 format that the address configuration was updated
- `url` - The URL of the address configuration

---

A `auto_creation` block supports the following:

- `service_sid` - The conversation service associated with the address configuration
- `integration_type` - The integration type used
- `enabled` - Whether conversation should auto-create when messages are received at the configured address
- `flow_sid` - The SID for the studio flow which will be called
- `retry_count` - The number of attempts to retry a failed webhook call
- `webhook_url` - The URL that is called when an event is triggered
- `webhook_filters` - A list of events that triggers a call to the webhook URL
- `webhook_method` - The HTTP method used to call the webhook URL

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving address configuration
