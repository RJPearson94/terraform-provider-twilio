---
page_title: "twilio_conversations_address_configuration Data Source - twilio"
subcategory: "Conversations"
description: |-
  
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

## Schema

### Required

- `sid` (String) The SID of the address configuration to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this address configuration
- `address` (String) The address (e.g. phone number) for the configuration
- `auto_creation` (List of Object) The auto-creation settings for the address configuration (see [below for nested schema](#nestedatt--auto_creation))
- `date_created` (String) The date and time the address configuration was created, in RFC 3339 format
- `date_updated` (String) The date and time the address configuration was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the address configuration
- `id` (String) The ID of this resource.
- `type` (String) The type of address
- `url` (String) The absolute URL of the address configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--auto_creation"></a>
### Nested Schema for `auto_creation`

Read-Only:

- `enabled` (Boolean)
- `flow_sid` (String)
- `integration_type` (String)
- `retry_count` (String)
- `service_sid` (String)
- `webhook_filters` (List of String)
- `webhook_method` (String)
- `webhook_url` (String)
