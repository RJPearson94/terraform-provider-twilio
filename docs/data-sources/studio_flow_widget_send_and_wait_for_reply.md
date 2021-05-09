---
page_title: "Twilio Studio Flow Widget - Send and wait for reply"
subcategory: "Studio"
---

# twilio_studio_flow_widget_send_and_wait_for_reply Data Source

Use this data source to generate the JSON for the Studio Flow send and wait for reply widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/send-wait-reply) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_send_and_wait_for_reply" "send_and_wait_for_reply" {
  name = "SendAndWaitForReply"

  transitions {
    incoming_message = "IncomingMessageTransition"
  }

  body = "Hello World"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_send_and_wait_for_reply" "send_and_wait_for_reply" {
  name = "SendAndWaitForReply"

  transitions {
    delivery_failure = "DeliveryFailureTransition"
    incoming_message = "IncomingMessageTransition"
    timeout          = "TimeoutTransition"
  }

  attributes = jsonencode({
    "channelSid" : "{{trigger.message.ChannelSid}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "name" : "{{trigger.message.ChannelAttributes.from}}"
  })
  body        = "Hello World"
  channel_sid = "{{trigger.message.ChannelSid}}"
  from        = "{{flow.channel.address}}"
  media_url   = "https://localhost.com"
  service_sid = "{{trigger.message.InstanceSid}}"
  timeout     = "300"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the send and wait for reply widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `attributes` - (Optional) A JSON string of Programmable Chat attributes to send with the message
- `body` - (Mandatory) The message content
- `channel_sid` - (Optional) The SID of the Programmable Chat channel the message should be sent. This value can be either a liquid template or a Programmable Chat Channel SID
- `from` - (Optional) The sender of the message. Default is `{{flow.channel.address}}`
- `media_url` - (Optional) A URL to media which is sent with the message. This value can be either a liquid template or a URL
- `service_sid` - (Optional) The SID of the Programmable Chat service the chat is hosted. This value can be either a liquid template or a Programmable Chat Service SID
- `timeout` - (Optional) A string to represent the time in seconds to wait for a reply. Default is `3600`

~> Due to data type and validation restrictions liquid templates are not supported for the `attributes` argument. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `delivery_failure` - (Optional) The widget to transition to when the message fails to send
- `incoming_message` - (Optional) The widget to transition to when the number receives an incoming message
- `timeout` - (Optional) The widget to transition to when the timeout limit is reached

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the send and wait for reply widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the send and wait for reply widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the send and wait for reply widget
- `json` - The JSON state definition for the send and wait for reply widget
