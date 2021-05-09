---
page_title: "Twilio Studio Flow Widget - Send to Flex"
subcategory: "Studio"
---

# twilio_studio_flow_widget_send_to_flex Data Source

Use this data source to generate the JSON for the Studio Flow send to flex widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/send-flex) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
  name         = "SendToFlex"
  channel_sid  = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
  name = "SendToFlex"

  transitions {
    call_complete     = "CallCompleteTransition"
    call_failure      = "CallFailureTransition"
    failed_to_enqueue = "FailedToEnqueue"
  }

  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })
  channel_sid     = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  priority        = "10"
  timeout         = "3600"
  wait_url        = "https://localhost.com/hold"
  wait_url_method = "POST"
  workflow_sid    = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the send to flex widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `attributes` - (Optional) A JSON string of the task attributes
- `channel_sid` - (Mandatory) The SID of the Flex/ TaskRouter channel to associate with the task
- `priority` - (Optional) A string to represent the priority of the task
- `timeout` - (Optional) A string to represent the time in seconds which the task can live for
- `wait_url` - (Optional) The URL for custom wait/ hold music
- `wait_url_method` - (Optional) The method to be used to request the wait/ hold music
- `workflow_sid` - (Mandatory) The SID of the TaskRouter workflow to assign the user to

~> Due to data type and validation restrictions liquid templates are not supported for the `attributes` (liquid templates can be used to set the attribute values), `channel_sid`, `wait_url`, `wait_url_method` and `workflow_sid` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `call_complete` - (Optional) The widget to transition to when the call is complete/ task is created
- `call_failure` - (Optional) The widget to transition to when the call failed/ task fails to create
- `failed_to_enqueue` - (Optional) The widget to transition to when the task failed to create or the call failed to enqueue

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the send to flex widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the send to flex widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the send to flex widget
- `json` - The JSON state definition for the send to flex widget
