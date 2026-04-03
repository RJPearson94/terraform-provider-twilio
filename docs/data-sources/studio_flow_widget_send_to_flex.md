---
page_title: "twilio_studio_flow_widget_send_to_flex Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `channel_sid` (String) The SID of the TaskRouter task channel (e.g. voice, chat) to use for routing in Flex
- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `workflow_sid` (String) The SID of the TaskRouter workflow to route the task through in Flex

### Optional

- `attributes` (String) A JSON string of custom attributes to attach to the TaskRouter task
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `priority` (String) The priority of the TaskRouter task, as a string integer
- `timeout` (String) The number of seconds before the TaskRouter task times out, as a string integer
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `wait_url` (String) The HTTP/HTTPS URL of the TwiML document to execute while the caller waits for a Flex agent
- `wait_url_method` (String) The HTTP method to use when fetching the `wait_url` document. Valid values: `GET`, `POST`

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

- `call_complete` (String) The name of the next widget when the Flex interaction completes
- `call_failure` (String) The name of the next widget when the Flex interaction fails
- `failed_to_enqueue` (String) The name of the next widget when the task cannot be enqueued to Flex
