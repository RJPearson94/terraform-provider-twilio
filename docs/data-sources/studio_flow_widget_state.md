---
page_title: "twilio_studio_flow_widget_state Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_state Data Source

This widget is the basic structure of a flow definition state object. This widget can be used in place of another pre-built widget or to build a widget that is not supported in the provider. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

### Basic

```hcl
data "twilio_studio_flow_widget_state" "state" {
  name = "State"
  type = "say-play"

  transitions {
    event = "audioComplete"
    next  = "State"
  }

  properties = {
    "digits" : "123"
  }
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `properties` (Map of String) A map of widget-specific properties. The expected keys depend on the `type` value
- `transitions` (Block List, Min: 1) The list of transitions for this generic state, each triggered by a named event (see [below for nested schema](#nestedblock--transitions))
- `type` (String) The widget type identifier (e.g. `send-message`, `make-http-request`). Use this generic state data source for widget types not covered by a dedicated data source

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Required:

- `event` (String) The event name that triggers this transition (e.g. `incomingMessage`, `audioComplete`)

Optional:

- `conditions` (Block List) Optional conditions that must be met for this transition to fire (see [below for nested schema](#nestedblock--transitions--conditions))
- `next` (String) The name of the next widget to transition to when this event fires

<a id="nestedblock--transitions--conditions"></a>
### Nested Schema for `transitions.conditions`

Required:

- `arguments` (List of String) The arguments to pass to the condition operator
- `friendly_name` (String) A human-readable label for this condition
- `type` (String) The comparison operator type (e.g. `equal_to`, `regex`)
- `value` (String) The value to compare against
