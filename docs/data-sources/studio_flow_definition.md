---
page_title: "twilio_studio_flow_definition Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_definition Data Source

This data source can be used to generate a Studio Flow definition JSON which can be supplied as an argument to the `twilio_studio_flow` resource

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

### Studio Flow definition with Trigger and Send to Flex widgets

```hcl
data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
  name = "SendMessageToAgent"

  workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid  = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })

  offset {
    x = 270
    y = 540
  }
}

data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  transitions {
    incoming_message = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
  }

  offset {
    x = 200
    y = 0
  }
}

data "twilio_studio_flow_definition" "definition" {
  description   = "Bot flow for creating a Flex webchat task"
  initial_state = data.twilio_studio_flow_widget_trigger.trigger.name

  flags {
    allow_concurrent_calls = true
  }

  states {
    json = data.twilio_studio_flow_widget_trigger.trigger.json
  }

  states {
    json = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.json
  }
}
```

### Studio Flow definition with Studio Flow Resource

```hcl
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  offset {
    x = 200
    y = 0
  }
}

data "twilio_studio_flow_definition" "definition" {
  description   = "Flow with trigger widget"
  initial_state = data.twilio_studio_flow_widget_trigger.trigger.name

  flags {
    allow_concurrent_calls = true
  }

  states {
    json = data.twilio_studio_flow_widget_trigger.trigger.json
  }
}

resource "twilio_studio_flow" "flow" {
  friendly_name = "With widgets"
  status        = "draft"
  definition    = data.twilio_studio_flow_definition.definition.json
  validate      = true
}
```

## Schema

### Required

- `description` (String) A human-readable description of the flow's purpose
- `initial_state` (String) The name of the widget that handles the flow's entry point (typically the Trigger widget)
- `states` (Block List, Min: 1) The list of widget states that make up the flow. Use the widget data sources (e.g. `twilio_studio_flow_widget_trigger`) to generate each state's JSON (see [below for nested schema](#nestedblock--states))

### Optional

- `flags` (Block List, Max: 1) Optional feature flags controlling flow execution behaviour (see [below for nested schema](#nestedblock--flags))

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) The assembled JSON flow definition, suitable for use in the `definition` argument of `twilio_studio_flow`

<a id="nestedblock--states"></a>
### Nested Schema for `states`

Required:

- `json` (String) The JSON representation of a single widget state, produced by a widget data source


<a id="nestedblock--flags"></a>
### Nested Schema for `flags`

Required:

- `allow_concurrent_calls` (Boolean) Whether the flow allows multiple simultaneous executions for the same contact
