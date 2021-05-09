---
page_title: "Twilio Studio Flow Widget - Make HTTP request"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the make HTTP request widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `body` - (Optional) The request body
- `charset` - (Optional) The character encoding of the request body. The default is `utf-8`
- `content_type` - (Mandatory) The content type of the request body. Valid values include: `application/x-www-form-urlencoded` or `application/json`
- `method` - (Mandatory) The HTTP method to be used when calling the URL. Valid values include: `GET` or `POST`
- `parameters` - (Optional) A list of `parameter` blocks as documented below
- `url` - (Mandatory) The URL which will be called. This value can be either a liquid template or a URL

~> Due to data type and validation restrictions liquid templates are not supported for the `content_type` and `method` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `parameter` block supports the following:

- `key` - (Mandatory) The parameter name/ key to pass to the HTTP request
- `value` - (Mandatory) The value of the parameter to pass to the HTTP request

---

A `transitions` block supports the following:

- `failed` - (Optional) The widget to transition to when the HTTP request fails
- `success` - (Optional) The widget to transition to when the HTTP request returns a 20X status code

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the make HTTP request widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the make HTTP request widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the make HTTP request widget
- `json` - The JSON state definition for the make HTTP request widget
