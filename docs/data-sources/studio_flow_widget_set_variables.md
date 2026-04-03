---
page_title: "twilio_studio_flow_widget_set_variables Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_set_variables Data Source

Use this data source to generate the JSON for the Studio Flow set variables widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/set-variables) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"
}
```

## Multiple variables

```hcl
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"

  variables {
    key   = "test"
    value = "testValue"
  }

  variables {
    key   = "test2"
    value = "testValue2"
  }
}
```

## With Transitions and Offset

```hcl
data "twilio_studio_flow_widget_set_variables" "set_variables" {
  name = "SetVariables"

  transitions {
    next = "NextTransition"
  }

  variables {
    key   = "test"
    value = "testValue"
  }

  offset {
    x = 10
    y = 20
  }
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `variables` (Block List) The list of flow variables to set or update (see [below for nested schema](#nestedblock--variables))

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

- `next` (String) The name of the next widget after the variables are set


<a id="nestedblock--variables"></a>
### Nested Schema for `variables`

Required:

- `key` (String) The variable name
- `value` (String) The variable value. Supports Liquid template expressions
