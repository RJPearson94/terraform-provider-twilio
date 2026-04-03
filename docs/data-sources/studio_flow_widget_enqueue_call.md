---
page_title: "twilio_studio_flow_widget_enqueue_call Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `priority` (Number) The priority of the TaskRouter task created for this call. Conflicts with `queue_name`
- `queue_name` (String) The name of the Twilio queue to enqueue the call into. At least one of `workflow_sid` or `queue_name` must be set
- `task_attributes` (String) A JSON string of attributes to attach to the TaskRouter task. Conflicts with `queue_name`
- `timeout` (Number) The number of seconds to wait for a worker to accept the task before timing out. Conflicts with `queue_name`
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `wait_url` (String) The HTTP/HTTPS URL of the TwiML document to execute while the caller waits in the queue
- `wait_url_method` (String) The HTTP method to use when fetching the `wait_url` document. Valid values: `GET`, `POST`
- `workflow_sid` (String) The SID of the TaskRouter workflow to route the call through. At least one of `workflow_sid` or `queue_name` must be set

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

- `call_complete` (String) The name of the next widget when the enqueued call completes
- `call_failure` (String) The name of the next widget when the enqueued call fails
- `failed_to_enqueue` (String) The name of the next widget when the call cannot be added to the queue or workflow
