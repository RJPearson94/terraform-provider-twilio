---
page_title: "twilio_studio_flow_widget_send_message Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `body` (String) The text body of the message to send. Supports Liquid template expressions
- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `attributes` (String) A JSON string of custom attributes to attach to the message
- `channel_sid` (String) The SID of the Programmable Chat channel to send the message to
- `from` (String) The sender address for the message. Defaults to `{{flow.channel.address}}`
- `media_url` (String) The HTTP/HTTPS URL of a media file to include with the message (e.g. an image or PDF)
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `service_sid` (String) The SID of the Programmable Chat service to use for sending the message
- `to` (String) The recipient address for the message. Defaults to `{{contact.channel.address}}`
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--offset"></a>
### Nested Schema for `offset`

Optional:

- `x` (Number) The x-axis position. Defaults to 0
- `y` (Number) The y-axis position. Defaults to 0


<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Optional:

- `failed` (String) The name of the next widget when the message fails to send
- `sent` (String) The name of the next widget when the message is sent successfully
