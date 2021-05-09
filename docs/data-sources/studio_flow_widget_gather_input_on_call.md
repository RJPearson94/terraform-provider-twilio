---
page_title: "Twilio Studio Flow Widget - Gather input on call"
subcategory: "Studio"
---

# twilio_studio_flow_widget_gather_input_on_call Data Source

Use this data source to generate the JSON for the Studio Flow gather input on call widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/gather-input-call) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Play

```hcl
data "twilio_studio_flow_widget_gather_input_on_call" "gather_input_on_call" {
  name = "GatherInputOnCall"

  play = "http://localhost.com"
}
```

## Say

```hcl
data "twilio_studio_flow_widget_gather_input_on_call" "gather_input_on_call" {
  name = "GatherInputOnCall"

  say = "Hello World"
}
```

## With all say config

```hcl
data "twilio_studio_flow_widget_gather_input_on_call" "gather_input_on_call" {
  name = "GatherInputOnCall"

  transitions {
    keypress = "KeypressTransition"
    speech   = "SpeechTransition"
    timeout  = "TimeoutTransition"
  }

  finish_on_key   = "1"
  gather_language = "en-US"
  hints = [
    "test",
    "test2"
  ]
  language         = "en-US"
  loop             = 1
  number_of_digits = 3
  profanity_filter = "true"
  say              = "Hello World"
  speech_model     = "phone_call"
  speech_timeout   = "auto"
  stop_gather      = true
  timeout          = 5
  voice            = "alice"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the gather input on call widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `finish_on_key` - (Optional) The keypress which will cause the data to be submitted
- `gather_language` - (Optional) The language which the NLU engine will try to interpret
- `hints` - (Optional) A list of hints
- `language` - (Optional) The language to use when speaking the message. This argument conflicts with `play`
- `loop` - (Optional) The number of times to say/ play content
- `number_of_digits` - (Optional) The number of digits to wait for before submitting the data
- `play` - (Optional) The URL of media content to play. This value can be either a liquid template or a URL
- `profanity_filter` - (Optional) Whether profanity should be redacted from the speech results. This value can be either a liquid template or the string `true` or `false`
- `say` - (Optional) The text to say
- `speech_model` - (Optional) The NLU model that should be used to interpret the speech. This value can be either a liquid template or the string `default`, `numbers_and_commands` or `phone_call`
- `speech_timeout` - (Optional) The amount of time in seconds of silence to wait for before timing out. This value can be the string `auto` or a positive integer as a string
- `stop_gather` - (Optional) Whether Twilio should listen for a keypress to stop capturing digits from the caller
- `timeout` - (Optional) The amount of time in seconds to wait for the caller to press a key
- `voice` - (Optional) The voice to use when speaking the message. This argument conflicts with `play`

~> Exactly one of the following arguments: `play` or `say` must be specified
~> Due to data type and validation restrictions liquid templates are not supported for the `loop`, `number_of_digits`, `stop_gather` and `timeout` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `keypress` - (Optional) The widget to transition to when the data is collected via key presses
- `speech` - (Optional) The widget to transition to when the data is collected via speech
- `timeout` - (Optional) The widget to transition to when the timeout limit has been reached

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the gather input on call widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the gather input on call widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the gather input on call widget
- `json` - The JSON state definition for the gather input on call widget
