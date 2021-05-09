---
page_title: "Twilio Studio Flow Widget - Enqueue call"
subcategory: "Studio"
---

# twilio_studio_flow_widget_enqueue_call Data Source

Use this data source to generate the JSON for the Studio Flow enqueue call widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/enqueue-call) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## With TaskRouter workflow

```hcl
data "twilio_studio_flow_widget_enqueue_call" "enqueue_call" {
  name         = "EnqueueCall"
  workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
```

## With Queue name

```hcl
data "twilio_studio_flow_widget_enqueue_call" "enqueue_call" {
  name       = "EnqueueCall"
  queue_name = "Test"
}
```

## With all TaskRouter workflow config

```hcl
data "twilio_studio_flow_widget_enqueue_call" "enqueue_call" {
  name = "EnqueueCall"

  transitions {
    call_complete     = "EnqueueCall"
    call_failure      = "EnqueueCall"
    failed_to_enqueue = "EnqueueCall"
  }

  priority = 1
  task_attributes = jsonencode({
    "test" : "test"
  })
  timeout         = 10
  wait_url        = "http://localhost.com"
  wait_url_method = "POST"
  workflow_sid    = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the enqueue call widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `priority` - (Optional) The priority of the task in the queue. This argument conflicts with `queue_name`
- `queue_name` - (Optional) The name of the queue to place the call on
- `task_attributes` - (Optional) A JSON string of attributes to be passed with the task. This argument conflicts with `queue_name`
- `timeout` - (Optional) The time in seconds which the task can remain on the queue before timing out. This argument conflicts with `queue_name`
- `wait_url` - (Optional) The URL for custom hold music
- `wait_url_method` - (Optional) The HTTP method to be used when calling the URL. Valid values include: `GET` or `POST`
- `workflow_sid` - (Optional) The SID of the TasksRouter workflow which will handle the task

~> Either the `queue_name` or `workflow_sid` argument must be set
~> Due to data type and validation restrictions liquid templates are not supported for the `priority`, `task_attributes`, `wait_url`, `wait_url_method` and `workflow_sid` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `call_complete` - (Optional) The widget to transition to when the call the enqueue action URL is requested
- `call_failure` - (Optional) The widget to transition to when a system error occurs
- `failed_to_enqueue` - (Optional) The widget to transition to when the call fails to enqueue

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the enqueue call widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the enqueue call widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the enqueue call widget
- `json` - The JSON state definition for the enqueue call widget
