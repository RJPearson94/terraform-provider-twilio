---
page_title: "twilio_flex_flow Data Source - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_flow Data Source

Use this data source to access information about an existing Twilio Flex Flow.

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

## Example Usage

```hcl
data "twilio_flex_flow" "flow" {
  sid = "FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "flow" {
  value = data.twilio_flex_flow.flow
}
```

## Schema

### Required

- `sid` (String) The SID of the Flex flow to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Flex flow
- `channel_type` (String) The channel type of the Flex flow. Valid values are `web`, `sms`, `facebook`, `whatsapp`, `line`, or `custom`
- `chat_service_sid` (String) The SID of the chat service associated with the Flex flow
- `contact_identity` (String) The channel contact identity for the Flex flow
- `date_created` (String) The date and time the Flex flow was created, in RFC 3339 format
- `date_updated` (String) The date and time the Flex flow was last updated, in RFC 3339 format
- `enabled` (Boolean) Whether the Flex flow is enabled
- `friendly_name` (String) A human-readable label for the Flex flow
- `id` (String) The ID of this resource.
- `integration` (List of Object) The integration settings for the Flex flow (see [below for nested schema](#nestedatt--integration))
- `integration_type` (String) The type of integration. Valid values are `studio`, `external`, or `task`
- `janitor_enabled` (Boolean) Whether the janitor is enabled to clean up expired channels
- `long_lived` (Boolean) Whether the Flex flow channel is long-lived
- `url` (String) The absolute URL of the Flex flow resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--integration"></a>
### Nested Schema for `integration`

Read-Only:

- `channel` (String)
- `creation_on_message` (Boolean)
- `flow_sid` (String)
- `priority` (Number)
- `retry_count` (Number)
- `timeout` (Number)
- `url` (String)
- `workflow_sid` (String)
- `workspace_sid` (String)
