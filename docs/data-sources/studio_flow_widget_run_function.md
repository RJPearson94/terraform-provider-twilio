---
page_title: "twilio_studio_flow_widget_run_function Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `url` (String) The HTTPS URL of the Serverless function to execute

### Optional

- `environment_sid` (String) The SID of the Serverless environment to execute the function in
- `function_sid` (String) The SID of the Serverless function to execute
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `parameters` (Block List) Key/value parameters to pass to the function as arguments (see [below for nested schema](#nestedblock--parameters))
- `service_sid` (String) The SID of the Serverless service containing the function, or `default` for the default service
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--offset"></a>
### Nested Schema for `offset`

Optional:

- `x` (Number) The x-axis position. Defaults to 0
- `y` (Number) The y-axis position. Defaults to 0


<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `key` (String) The parameter name
- `value` (String) The parameter value


<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Optional:

- `fail` (String) The name of the next widget when the function execution fails
- `success` (String) The name of the next widget when the function executes successfully
