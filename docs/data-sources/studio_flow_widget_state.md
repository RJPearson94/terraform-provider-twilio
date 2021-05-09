---
page_title: "Twilio Studio Flow Widget - State"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the widget
- `type` - (Mandatory) The type of the widget
- `properties` - (Mandatory) A map of properties for the widget
- `transitions` - (Optional) A list of `transition` blocks as documented below

---

A `transition` block supports the following:

- `event` - (Mandatory) The name of the event which will trigger a transition
- `next` - (Optional) The next state to transition to when the transition is activated
- `conditions` - (Optional) A list of `condition` blocks as documented below

---

A `condition` block supports the following:

- `arguments` - (Mandatory) A list of arguments to evaluate
- `friendly_name` - (Mandatory) The name of the condition
- `type` - (Mandatory) The type/ operator to use when comparing the arguments and value
- `value` - (Mandatory) The value or values to compare against

## Attributes Reference

The following attributes are exported:

- `id` - The name of the widget
- `json` - The JSON state definition for the widget
