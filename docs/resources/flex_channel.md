---
page_title: "Twilio Flex Channel"
subcategory: "Flex"
---

# twilio_flex_channel Resource

Manages a Twilio Flex channel

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

## Example Usage

```hcl
resource "twilio_flex_flow" "flow" {
  friendly_name    = "twilio-test"
  chat_service_sid = var.chat_service_sid
  channel_type     = "web"
  integration_type = "external"
  integration {
    url = "https://test.com/external"
  }
}

resource "twilio_flex_channel" "channel" {
  chat_friendly_name      = "twilio-test"
  chat_user_friendly_name = "twilio-test"
  flex_flow_sid           = twilio_flex_flow.flow.sid
  identity                = "test"
}
```

## Argument Reference

The following arguments are supported:

- `chat_friendly_name` - (Mandatory) The friendly name for the chat channel
- `chat_user_friendly_name` - (Mandatory) The friendly name for the chat participant/ user
- `flex_flow_sid` - (Mandatory) The SID of the flow
- `identity` - (Mandatory) The identifier associated with the chat user
- `chat_unique_name` - (Optional) The unique name for the chat channel
- `long_lived` - (Optional) Whether to create a long lived channel
- `pre_engagement_data` - (Optional) The pre-engagement data collected
- `target` - (Optional) The contact identifier of the target
- `task_attributes` - (Optional) The attributes to be assigned to the task
- `task_sid` - (Optional) The SID of the task

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel (Same as the SID)
- `sid` - The SID of the channel (Same as the ID)
- `account_sid` - The account SID associated with the channel
- `chat_friendly_name` - The friendly name for the chat channel
- `chat_user_friendly_name` - The friendly name for the chat participant/ user
- `flex_flow_sid` - The SID of the flow
- `identity` - The identifier associated with the chat user
- `chat_unique_name` - The unique name for the chat channel
- `long_lived` - Whether to create a long lived channel
- `pre_engagement_data` - The pre-engagement data collected
- `target` - The contact identifier of the target
- `task_attributes` - The attributes to be assigned to the task
- `task_sid` - The SID of the task
- `user_sid` - The SID of the user
- `date_created` - The date in RFC3339 format that the channel was created
- `date_updated` - The date in RFC3339 format that the channel was updated
- `url` - The url of the channel

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the channel
- `read` - (Defaults to 5 minutes) Used when retrieving the channel
- `delete` - (Defaults to 10 minutes) Used when deleting the channel

## Import

A channel can be imported using the `/Channels/{sid}` format, e.g.

```shell
terraform import twilio_flex_channel.channel /Channels/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> The following arguments "chat_friendly_name", "chat_unique_name", "chat_user_friendly_name", "long_lived" and "identity" cannot be imported, as the API doesn't return this data
