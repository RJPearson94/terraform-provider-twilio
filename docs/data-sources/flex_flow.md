---
page_title: "Twilio Flex Flow"
subcategory: "Flex"
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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the flex-flow

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the flow
