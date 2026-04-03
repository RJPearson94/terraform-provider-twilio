---
page_title: "twilio_studio_flow_widget_connect_virtual_agent Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_connect_virtual_agent Data Source

Use this data source to generate the JSON for the Studio Flow connect virtual agent widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/connect-virtual-agent) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_connect_virtual_agent" "connect_virtual_agent" {
  name      = "ConnectVirtualAgent"
  connector = "test-connector"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_connect_virtual_agent" "connect_virtual_agent" {
  name = "ConnectVirtualAgent"

  transitions {
    hangup = "HangupTransition"
    return = "ReturnTransition"
  }

  connector           = "test-connector"
  sentiment_analysis  = "true"
  language            = "en-US"
  status_callback_url = "https://test.com"

  offset {
    x = 10
    y = 20
  }
}
```

## Schema

### Required

- `connector` (String) The unique name of the virtual agent connector to use
- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `language` (String) The BCP-47 language tag for the virtual agent session (e.g. `en-US`)
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `sentiment_analysis` (String) Whether to enable sentiment analysis during the virtual agent session. Valid values: `true`, `false`
- `status_callback_url` (String) The HTTP/HTTPS URL to receive status callback events for the virtual agent session
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

- `hangup` (String) The name of the next widget when the caller hangs up during the virtual agent session
- `return` (String) The name of the next widget when the virtual agent session completes and returns control
