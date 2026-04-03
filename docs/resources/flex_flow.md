---
page_title: "twilio_flex_flow Resource - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_flow Resource

Manages a flex-flow

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

## Example Usage

```hcl
resource "twilio_flex_flow" "flow" {
  friendly_name    = "twilio-test"
  chat_service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_type     = "web"

  integration_type = "external"
  integration {
    url = "https://test.com/external"
  }
}
```

## Schema

### Required

- `channel_type` (String) The channel type of the Flex flow. Valid values are `web`, `sms`, `facebook`, `whatsapp`, `line`, or `custom`
- `chat_service_sid` (String) The SID of the chat service associated with the Flex flow
- `friendly_name` (String) A human-readable label for the Flex flow
- `integration` (Block List, Min: 1, Max: 1) The integration settings for the Flex flow (see [below for nested schema](#nestedblock--integration))

### Optional

- `contact_identity` (String) The channel contact identity for the Flex flow
- `enabled` (Boolean) Whether the Flex flow is enabled. Defaults to `false`
- `integration_type` (String) The type of integration. Valid values are `studio`, `external`, or `task`
- `janitor_enabled` (Boolean) Whether the janitor is enabled to clean up expired channels. Defaults to `false`
- `long_lived` (Boolean) Whether the Flex flow channel is long-lived. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Flex flow
- `date_created` (String) The date and time the Flex flow was created, in RFC 3339 format
- `date_updated` (String) The date and time the Flex flow was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this Flex flow by Twilio
- `url` (String) The absolute URL of the Flex flow resource

<a id="nestedblock--integration"></a>
### Nested Schema for `integration`

Optional:

- `channel` (String) The channel for the integration
- `creation_on_message` (Boolean) Whether to create a task when a message is received
- `flow_sid` (String) The SID of the Studio flow for the integration
- `priority` (Number) The priority of the task in TaskRouter
- `retry_count` (Number) The number of times to retry the integration. Valid values are between `0` and `3`
- `timeout` (Number) The timeout in seconds for the integration
- `url` (String) The URL for an external integration. Must use HTTP or HTTPS
- `workflow_sid` (String) The SID of the TaskRouter workflow for the integration
- `workspace_sid` (String) The SID of the TaskRouter workspace for the integration


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A flow can be imported using the `/FlexFlows/{sid}` format, e.g.

```shell
terraform import twilio_flex_flow.flow /FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
