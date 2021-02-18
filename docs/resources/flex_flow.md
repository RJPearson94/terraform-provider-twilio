---
page_title: "Twilio Flex Flow"
subcategory: "Flex"
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

## Argument Reference

The following arguments are supported:

- `channel_type` - (Mandatory) The type of channel which is integrated with the flow. Valid values are `web`, `sms`, `facebook`, `whatsapp`, `line` or `custom`
- `chat_service_sid` - (Mandatory) The chat service SID to associate with the flow
- `friendly_name` - (Mandatory) The friendly name of the flow
- `contact_identity` - (Optional) The contact identity for the channel
- `enabled` - (Optional) Whether the flow is active
- `integration_type` - (Optional) The type of integration with the flow. Valid values are `studio`, `external` or `task`
- `janitor_enabled` - (Optional) Clean up chat channels and proxy sessions when the task is completed
- `long_lived` - (Optional) Whether to reuse the same channel for any future interactions with the customer
- `integration` - (Optional) A `integration` block as documented below

---

An `integration` block supports the following:

- `channel` - (Optional) The channel to send new tasks too
- `creation_on_message` - (Optional) Whether to create a task when the first message arrives
- `flow_sid` - (Optional) The SID of the flow
- `priority` - (Optional) The priority assigned to any new task that is received
- `retry_count` - (Optional) The number of times a webhook request should be retried if the initial request fails
- `timeout` - (Optional) The timeout set for any new task that is received
- `url` - (Optional) The webhook URL
- `workflow_sid` - (Optional) The SID of the workflow to send tasks to
- `workspace_sid` - (Optional) The SID of the workspace to send tasks to

---

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the flow (Same as the `sid`)
- `sid` - The SID of the flow (Same as the `id`)
- `account_sid` - The account SID associated with the flow
- `channel_type` - The type of channel which is integrated with the flow
- `chat_service_sid` - The chat service SID associated with the flow
- `friendly_name` - The friendly name of the flow
- `contact_identity` - The contact identity for the channel
- `enabled` - Whether the flow is active
- `integration_type` - The type of integration with the flow
- `janitor_enabled` - Clean up chat channels and proxy sessions when the task is completed
- `long_lived` - Whether to reuse the same channel for any future interactions with the customer
- `integration` - A `integration` block as documented below
- `date_created` - The date in RFC3339 format that the flow was created
- `date_updated` - The date in RFC3339 format that the flow was updated
- `url` - The URL of the flow

---

An `integration` block supports the following:

- `channel` - The channel to send new tasks too
- `creation_on_message` - Whether to create a task when the first message arrives
- `flow_sid` - The SID of the flow
- `priority` - The priority assigned to any new task that is received
- `retry_count` - The number of times a webhook request should be retried if the initial request fails
- `timeout` - The timeout set for any new task that is received
- `url` - The webhook URL
- `workflow_sid` - The SID of the workflow to send tasks to
- `workspace_sid` - The SID of the workspace to send tasks to

---

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the flow
- `update` - (Defaults to 10 minutes) Used when updating the flow
- `read` - (Defaults to 5 minutes) Used when retrieving the flow
- `delete` - (Defaults to 10 minutes) Used when deleting the flow

## Import

A flow can be imported using the `/FlexFlows/{sid}` format, e.g.

```shell
terraform import twilio_flex_flow.flow /FlexFlows/FOXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
