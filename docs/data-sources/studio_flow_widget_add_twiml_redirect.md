---
page_title: "Twilio Studio Flow Widget - Add TwiML redirect"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the add TwiML redirect widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `method` - (Optional) The HTTP method to be used when calling the URL. This value can be either a liquid template or the string `GET` or `POST`
- `timeout` - (Optional) The time in seconds to wait for the redirect to complete and return control to the Studio flow. This value can be either a liquid template or a number string in the range `0` to `14400` (inclusive)
- `url` - (Mandatory) The URL where the call/ message will be redirected to. This value can be either a liquid template or a URL

---

A `transitions` block supports the following:

- `fail` - (Optional) The widget to transition to when the redirect fails
- `return` - (Optional) The widget to transition to when the redirect completes and the flow execution should continue
- `timeout` - (Optional) The widget to transition to when the redirect timeout is reached

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the add TwiML redirect widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the add TwiML redirect widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the add TwiML redirect widget
- `json` - The JSON state definition for the add TwiML redirect widget
