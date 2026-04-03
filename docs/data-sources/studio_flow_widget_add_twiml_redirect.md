---
page_title: "twilio_studio_flow_widget_add_twiml_redirect Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_add_twiml_redirect Data Source

Use this data source to generate the JSON for the Studio Flow add TwiML redirect widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/twiml-redirect) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_add_twiml_redirect" "add_twiml_redirect" {
  name = "AddTwiMLRedirect"
  url  = "https://test.com/twiml"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_add_twiml_redirect" "add_twiml_redirect" {
  name = "AddTwiMLRedirect"

  transitions {
    fail    = "FailTransition"
    return  = "ReturnTransition"
    timeout = "TimeoutTransition"
  }

  method  = "POST"
  timeout = "100"
  url     = "https://test.com/twiml"

  offset {
    x = 10
    y = 20
  }
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `url` (String) The absolute URL of the TwiML document to redirect to. Must be an HTTP/HTTPS URL or a Liquid template expression

### Optional

- `method` (String) The HTTP method to use when fetching the TwiML document. Valid values: `GET`, `POST`
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `timeout` (String) The timeout in seconds for the TwiML redirect request (0–14400)
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

- `fail` (String) The name of the next widget when the TwiML redirect request fails
- `return` (String) The name of the next widget when the TwiML redirect completes and returns
- `timeout` (String) The name of the next widget when the TwiML redirect request times out
