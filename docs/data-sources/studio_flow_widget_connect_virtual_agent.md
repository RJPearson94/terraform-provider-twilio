---
page_title: "Twilio Studio Flow Widget - Connect virtual agent"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the connect virtual agent widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `connector` - (Mandatory) The unique name of the Dialogflow ES connector
- `language` - (Optional) The language used by the Dialogflow ES agent
- `sentiment_analysis` - (Optional) Whether sentiment analysis should be performed. This value can be either a liquid template or the string `true` or `false`
- `status_callback_url` - (Optional) URL which receives the status callbacks

---

A `transitions` block supports the following:

- `hangup` - (Optional) The widget to transition to when the caller hangs up
- `return` - (Optional) The widget to transition to when the call is complete

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the connect virtual agent widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the connect virtual agent widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the connect virtual agent widget
- `json` - The JSON state definition for the connect virtual agent widget
