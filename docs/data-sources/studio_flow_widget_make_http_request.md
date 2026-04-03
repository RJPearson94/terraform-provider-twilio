---
page_title: "twilio_studio_flow_widget_make_http_request Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_make_http_request Data Source

Use this data source to generate the JSON for the Studio Flow make HTTP request widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/http-request) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_make_http_request" "make_http_request" {
  name = "MakeHttpRequest"

  method       = "GET"
  content_type = "application/x-www-form-urlencoded"
  url          = "https://test.com"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_make_http_request" "make_http_request" {
  name = "MakeHttpRequest"

  transitions {
    failed  = "FailedTransition"
    success = "SuccessTransition"
  }

  method       = "POST"
  content_type = "application/json"
  url          = "https://test.com"
  body = jsonencode({
    "say" : "Hello World"
  })

  parameters {
    key   = "key"
    value = "value"
  }

  parameters {
    key   = "key2"
    value = "value2"
  }

  offset {
    x = 10
    y = 20
  }
}
```

## Schema

### Required

- `content_type` (String) The Content-Type header for the request. Valid values: `application/x-www-form-urlencoded`, `application/json`
- `method` (String) The HTTP method for the request. Valid values: `GET`, `POST`
- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `url` (String) The HTTP/HTTPS URL to send the request to. Supports Liquid template expressions

### Optional

- `body` (String) The request body content for POST requests
- `charset` (String) The character encoding for the request body. Defaults to `utf-8`
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `parameters` (Block List) Key/value parameters to include in the request as query parameters (GET) or form data (POST) (see [below for nested schema](#nestedblock--parameters))
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

- `failed` (String) The name of the next widget when the HTTP request fails
- `success` (String) The name of the next widget when the HTTP request succeeds
