---
page_title: "Twilio Studio Flow Widget - Record call"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the record call widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `record_call` - (Optional) Whether the call should be recorded. The default value is `false`
- `trim` - (Optional) A string to indicate whether silence should be removed from the end of the recording. This value can be either a liquid template or the string `trim-silence` or `do-not-trim`
- `recording_status_callback_method` - (Optional) The HTTP method to use when calling the recording status callback URL. This value can be either a liquid template or either `GET` or `POST`
- `recording_status_callback_url` - (Optional) The URL which receives the recording completion callback. This value can be either a liquid template or a URL
- `recording_channels` - (Optional) The channel or channels which are to be recorded. This value can be either a liquid template or either `dual` or `mono`
- `recording_status_callback_events` - (Optional) A list of events that causes the recording status callback URL to be called. The valid elements in the list include: `absent`, `completed` and `in-progress`

~> Due to data type and validation restrictions liquid templates are not supported for the `record_call` and `recording_channels` (items) arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `failed` - (Optional) The widget to transition to when the recording fails
- `success` - (Optional) The widget to transition to when the call is recorded

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the record call widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the record call widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the record call widget
- `json` - The JSON state definition for the record call widget
