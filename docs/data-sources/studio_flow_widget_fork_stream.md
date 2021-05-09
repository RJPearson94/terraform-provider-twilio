---
page_title: "Twilio Studio Flow Widget - Fork stream"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the fork stream widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented belows
- `stream_action` - (Mandatory) The action you want to perform on a stream. Valid values include: `start` or `stop`
- `stream_connector` - (Optional) The name of the SIPREC stream
- `stream_name` - (Optional) The friendly name of the stream
- `stream_parameters` - (Optional) A list of `stream_parameter` blocks as documented below
- `stream_track` - (Optional) The stream tracks which will be sent. Valid values include: `both_tracks`, `inbound_track` or `outbound_track`
- `stream_transport_type` - (Optional) The transport protocol to use. Valid values include: `siprec` or `websocket`
- `stream_url` - (Optional) The websocket URL to stream audio to. This value can be either a liquid template or a URL starting with `wss://`

~> Due to data type and validation restrictions liquid templates are not supported for the `stream_action`, `stream_track` and `stream_transport_type` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `stream_parameter` block supports the following:

- `key` - (Mandatory) The parameter name/ key sent to the remote service
- `value` - (Mandatory) The value of the parameter sent to the remote service

---

A `transitions` block supports the following:

- `next` - (Optional) The widget to transition to when the stream has been started or stopped

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the fork stream widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the fork stream widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the fork stream widget
- `json` - The JSON state definition for the fork stream widget
