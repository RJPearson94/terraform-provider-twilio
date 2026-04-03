---
page_title: "twilio_conversations_configuration Data Source - twilio"
subcategory: "Conversations"
description: |-
  
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

## Schema

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this configuration
- `default_closed_timer` (String) The default ISO 8601 duration after which a conversation will be automatically closed
- `default_inactive_timer` (String) The default ISO 8601 duration after which a conversation will be marked as inactive
- `default_messaging_service_sid` (String) The SID of the default messaging service
- `default_service_sid` (String) The SID of the default conversations service
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
