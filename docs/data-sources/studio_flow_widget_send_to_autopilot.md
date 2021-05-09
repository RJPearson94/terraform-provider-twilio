---
page_title: "Twilio Studio Flow Widget - Send to Autopilot"
subcategory: "Studio"
---

# twilio_studio_flow_widget_send_to_autopilot Data Source

Use this data source to generate the JSON for the Studio Flow send to Autopilot widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/autopilot) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_send_to_autopilot" "send_to_autopilot" {
  name                    = "SendToAutopilot"
  autopilot_assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_send_to_autopilot" "send_to_autopilot" {
  name = "SendToAutopilot"

  transitions {
    failure       = "FailureTransition"
    session_ended = "SessionEndedTransition"
    timeout       = "TimeoutTransition"
  }

  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })
  autopilot_assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  body                    = "Hello World"
  channel_sid             = "{{trigger.message.ChannelSid}}"
  from                    = "test"
  memory_parameters {
    key   = "key"
    value = "value"
  }
  memory_parameters {
    key   = "key2"
    value = "value2"
  }
  service_sid = "{{trigger.message.InstanceSid}}"
  target_task = "Task"
  timeout     = 100

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the send to autopilot widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `attributes` - (Optional) A JSON string of Programmable Chat attributes to send to the assistant
- `autopilot_assistant_sid` - (Optional) The SID of the Autopilot Assistant to integrate with
- `body` - (Optional) The body/ message to send to the assistant. Default value is `{{trigger.Message.Body}}`
- `channel_sid` - (Optional) The SID of the Programmable Chat channel the bot should reply to. This value can be either a liquid template or a Programmable Chat Channel SID
- `from` - (Optional) The name of the bot that is visible in the chat. The default value is `{{flow.channel.address}}`
- `memory_parameters` - (Optional) A list of `memory_parameter` blocks as documented below
- `service_sid` - (Optional) The SID of the Programmable Chat service the chat is hosted. This value can be either a liquid template or a Programmable Chat Service SID
- `target_task` (Optional) Override the default task
- `timeout` (Optional) The time in seconds to wait before timing out. The default value is `14400`

~> Due to data type and validation restrictions liquid templates are not supported for the `autopilot_assistant_sid`, `attributes` (liquid templates can be used to set the attribute values) and `timeout` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `memory_parameter` block supports the following:

- `key` - (Mandatory) The parameter name/ key to store in the memory
- `value` - (Mandatory) The value of the parameter to store in the memory

---

A `transitions` block supports the following:

- `failure` - (Optional) The widget to transition to when the bot returns an error
- `session_ended` - (Optional) The widget to transition to when the session has ended
- `timeout` - (Optional) The widget to transition to when the integration has timed out

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the send to autopilot widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the send to autopilot widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the send to autopilot widget
- `json` - The JSON state definition for the send to autopilot widget
