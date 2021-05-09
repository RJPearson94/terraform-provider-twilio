---
page_title: "Twilio Studio Flow Widget - Run function"
subcategory: "Studio"
---

# twilio_studio_flow_widget_run_function Data Source

Use this data source to generate the JSON for the Studio Flow run function widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/run-function) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_run_function" "run_function" {
  name = "RunFunction"
  url  = "https://test-function.twil.io/test-function"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_run_function" "run_function" {
  name = "RunFunction"

  transitions {
    fail    = "FailTransition"
    success = "SuccessTransition"
  }

  function_sid    = "ZHaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  environment_sid = "ZEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  service_sid     = "ZSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  parameters {
    key   = "key"
    value = "value"
  }
  parameters {
    key   = "key2"
    value = "value2"
  }
  url = "https://test-function.twil.io/test-function"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the run function widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `environment_sid` - (Optional) The SID of a Serverless environment
- `function_sid` - (Optional) The SID of a Serverless function
- `parameters` - (Optional) A list of `parameter` blocks as documented below
- `service_sid` - (Optional) The SID of a Serverless service or default, if you are using the legacy functions and assets
- `url` - (Mandatory) The function URL to call/ invoke

~> Due to data type and validation restrictions liquid templates are not supported for the `environment_sid`, `function_sid`, `service_sid` and `url` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `parameter` block supports the following:

- `key` - (Mandatory) The parameter name/ key to pass to the function
- `value` - (Mandatory) The value of the parameter to pass to the function

---

A `transitions` block supports the following:

- `fail` - (Optional) The widget to transition to when the function fails to invoke or the function returns a 4xx or 5xx status code
- `success` - (Optional) The widget to transition to when the function is successfully invoked

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the run function widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the run function widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the run function widget
- `json` - The JSON state definition for the run function widget
