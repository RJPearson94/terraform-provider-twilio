---
page_title: "twilio_studio_flow_widget_fork_stream Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_fork_stream Data Source

Use this data source to generate the JSON for the Studio Flow fork stream widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/fork-stream) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Start stream

```hcl
data "twilio_studio_flow_widget_fork_stream" "fork_stream" {
  name = "ForkStream"

  stream_action         = "start"
  stream_name           = "test"
  stream_track          = "inbound_track"
  stream_transport_type = "websocket"
  stream_url            = "wss://test.com"
}
```

## Stop stream

```hcl
data "twilio_studio_flow_widget_fork_stream" "fork_stream" {
  name = "ForkStream"

  stream_transport_type = "websocket"
  stream_action         = "stop"
}
```

## With all start stream config

```hcl
data "twilio_studio_flow_widget_fork_stream" "fork_stream" {
  name = "ForkStream"

  transitions {
    next = "NextTransition"
  }

  stream_action    = "start"
  stream_connector = "connector"
  stream_name      = "test"
  stream_parameters {
    key   = "key"
    value = "value"
  }
  stream_parameters {
    key   = "key2"
    value = "value2"
  }
  stream_track          = "inbound_track"
  stream_transport_type = "siprec"

  offset {
    x = 10
    y = 20
  }
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `stream_action` (String) Whether to start or stop the audio stream. Valid values: `start`, `stop`

### Optional

- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `stream_connector` (String) The name of the SIPREC connector when using `siprec` transport
- `stream_name` (String) A friendly name for the stream, used to reference it when stopping
- `stream_parameters` (Block List) Additional key/value parameters to send to the remote stream service (see [below for nested schema](#nestedblock--stream_parameters))
- `stream_track` (String) Which audio track(s) to stream. Valid values: `both_tracks`, `inbound_track`, `outbound_track`
- `stream_transport_type` (String) The transport protocol for the stream. Valid values: `siprec`, `websocket`
- `stream_url` (String) The WebSocket URL (`wss://`) to stream audio to when using `websocket` transport
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))

### Read-Only

- `id` (String) The ID of this resource.
- `json` (String) A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source

<a id="nestedblock--offset"></a>
### Nested Schema for `offset`

Optional:

- `x` (Number) The x-axis position. Defaults to 0
- `y` (Number) The y-axis position. Defaults to 0


<a id="nestedblock--stream_parameters"></a>
### Nested Schema for `stream_parameters`

Required:

- `key` (String) The parameter name
- `value` (String) The parameter value


<a id="nestedblock--transitions"></a>
### Nested Schema for `transitions`

Optional:

- `next` (String) The name of the next widget after the stream is started or stopped
