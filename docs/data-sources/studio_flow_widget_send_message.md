---
page_title: "Twilio Studio Flow Widget - Send message"
subcategory: "Studio"
---

# twilio_studio_flow_widget_send_message Data Source

Use this data source to generate the JSON for the Studio Flow send message widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/send-message) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_send_message" "send_message" {
  name = "SendMessage"
  body = "Hello World"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_send_message" "send_message" {
  name = "SendMessage"

  transitions {
    failed = "FailedTransition"
    sent   = "SentTransition"
  }

  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })
  body        = "Hello World"
  channel_sid = "{{trigger.message.ChannelSid}}"
  from        = "{{flow.channel.address}}"
  media_url   = "https://test.com"
  service_sid = "{{trigger.message.InstanceSid}}"
  to          = "{{contact.channel.address}}"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the send message widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `attributes` - (Optional) A JSON string of Programmable Chat attributes to send with the message
- `body` - (Mandatory) The message content
- `channel_sid` - (Optional) The SID of the Programmable Chat channel the message should be sent. This value can be either a liquid template or a Programmable Chat Channel SID
- `from` - (Optional) The sender of the message. Default is `{{flow.channel.address}}`
- `media_url` - (Optional) A URL to media which is sent with the message. This value can be either a liquid template or a URL
- `service_sid` - (Optional) The SID of the Programmable Chat service the chat is hosted. This value can be either a liquid template or a Programmable Chat Service SID
- `to` - (Optional) The recipient of the message. Default is `{{contact.channel.address}}`

~> Due to data type and validation restrictions liquid templates are not supported for the `attributes` argument. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `failed` - (Optional) The widget to transition to when the message fails to send
- `sent` - (Optional) The widget to transition to when the message has been sent

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the send message widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the send message widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the send message widget
- `json` - The JSON state definition for the send message widget
