---
page_title: "Twilio Studio Flow Widget - Set Variables"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the set variables widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `variables` - (Optional) A list of `variable` blocks as documented below

---

A `variable` block supports the following:

- `key` - (Mandatory) The variable name/ key which will be used to access the variable value in the Studio Flow
- `value` - (Mandatory) The value of the variable

---

A `transitions` block supports the following:

- `next` - (Optional) The widget to transition to when the variable or variables have been set

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the set variables widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the set variables widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the set variables widget
- `json` - The JSON state definition for the set variables widget
