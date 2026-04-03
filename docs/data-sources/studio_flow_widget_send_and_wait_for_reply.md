---
page_title: "twilio_studio_flow_widget_send_and_wait_for_reply Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `body` (String) The text body of the message to send. Supports Liquid template expressions
- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `transitions` (Block List, Min: 1, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))

### Optional

- `attributes` (String) A JSON string of custom attributes to attach to the message
- `channel_sid` (String) The SID of the Programmable Chat channel to send the message to
- `from` (String) The sender address for the outgoing message. Defaults to `{{flow.channel.address}}`
- `media_url` (String) The HTTP/HTTPS URL of a media file to include with the message (e.g. an image or PDF)
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `service_sid` (String) The SID of the Programmable Chat service to use for sending the message
- `timeout` (String) The number of seconds to wait for a reply before timing out. Defaults to `3600` (1 hour)

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Required:

- `incoming_message` (String) The name of the next widget when a reply message is received

Optional:

- `delivery_failure` (String) The name of the next widget when message delivery fails
- `timeout` (String) The name of the next widget when the wait for reply times out


<a id="nestedblock--offset"></a>
### Nested Schema for `offset`

Optional:

- `x` (Number) The x-axis position. Defaults to 0
- `y` (Number) The y-axis position. Defaults to 0
