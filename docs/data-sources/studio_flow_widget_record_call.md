---
page_title: "twilio_studio_flow_widget_record_call Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_record_call Data Source

Use this data source to generate the JSON for the Studio Flow record call widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/call-recording) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_record_call" "record_call" {
  name        = "RecordCall"
  record_call = false
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_record_call" "record_call" {
  name = "RecordCall"

  transitions {
    failed  = "FailedTransition"
    success = "SuccessTransition"
  }

  record_call = true
  recording_status_callback_events = [
    "in-progress",
    "completed"
  ]
  recording_channels               = "mono"
  recording_status_callback_method = "GET"
  recording_status_callback_url    = "http://localhost.com"
  trim                             = "do-not-trim"

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
- `record_call` (Boolean) Whether to enable call recording. Defaults to false
- `recording_channels` (String) The number of recording channels. Valid values: `dual`, `mono`
- `recording_status_callback_events` (List of String) The recording events that trigger a callback. Valid values: `absent`, `completed`, `in-progress`
- `recording_status_callback_method` (String) The HTTP method for the recording status callback. Valid values: `GET`, `POST`
- `recording_status_callback_url` (String) The HTTP/HTTPS URL to receive recording status callback events
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `trim` (String) Whether to trim silence from the recording. Valid values: `trim-silence`, `do-not-trim`

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

- `failed` (String) The name of the next widget when the recording operation fails
- `success` (String) The name of the next widget when the recording operation succeeds
