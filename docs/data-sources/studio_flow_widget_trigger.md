---
page_title: "Twilio Studio Flow Widget - Trigger"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the trigger widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below

---

A `transitions` block supports the following:

- `incoming_call` - (Optional) The widget to transition to when an incoming call is received
- `incoming_message` - (Optional) The widget to transition to when an incoming message is received
- `incoming_request` - (Optional) The widget to transition to when an incoming request is received

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the trigger widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the trigger widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the trigger widget
- `json` - The JSON state definition for the trigger widget
