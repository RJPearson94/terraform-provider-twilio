---
page_title: "twilio_studio_flow_widget_trigger Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_trigger Data Source

Use this data source to generate the JSON for the Studio Flow trigger widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/trigger-start) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

### Basic

```hcl
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"
}
```

### With Offset and Transitions

```hcl
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  transitions {
    incoming_call    = "IncomingCallTransition"
    incoming_message = "IncomingMessageTransition"
    incoming_parent  = "IncomingParentTransition"
    incoming_request = "IncomingRequestTransition"
  }

  offset {
    x = 10
    y = 20
  }
}
```

### With Studio Flow Definition Data Source

```hcl
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"
}

data "twilio_studio_flow_definition" "definition" {
  description   = "Example Studio Flow with Trigger widget"
  initial_state = data.twilio_studio_flow_widget_trigger.trigger.name

  flags {
    allow_concurrent_calls = true
  }

  states {
    json = data.twilio_studio_flow_widget_trigger.trigger.json
  }
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
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

- `incoming_call` (String) The name of the next widget when the flow is triggered by an incoming call
- `incoming_message` (String) The name of the next widget when the flow is triggered by an incoming message
- `incoming_parent` (String) The name of the next widget when the flow is triggered by a parent flow
- `incoming_request` (String) The name of the next widget when the flow is triggered by an incoming REST API request
