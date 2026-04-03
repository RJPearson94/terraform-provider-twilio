---
page_title: "twilio_studio_flow_widget_gather_input_on_call Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `finish_on_key` (String) The DTMF key that signals the end of digit input (e.g. `#`)
- `gather_language` (String) The language for automatic speech recognition (e.g. `en-US`)
- `hints` (List of String) A list of words or phrases to improve speech recognition accuracy
- `language` (String) The language for text-to-speech when using `say`. Conflicts with `play`
- `loop` (Number) How many times to repeat the prompt before timing out. Use 0 for infinite loop
- `number_of_digits` (Number) The exact number of DTMF digits to collect before automatically submitting
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `play` (String) URL of the audio file to play as a prompt. Exactly one of `say` or `play` must be set
- `profanity_filter` (String) Whether to filter profanity from speech recognition results. Valid values: `true`, `false`
- `say` (String) Text to speak to the caller as a prompt using text-to-speech. Exactly one of `say` or `play` must be set
- `speech_model` (String) The speech recognition model to use. Valid values: `default`, `numbers_and_commands`, `phone_call`
- `speech_timeout` (String) The duration of silence (in seconds) before speech recognition ends, or `auto` for automatic detection
- `stop_gather` (Boolean) Whether to allow a keypress to stop the current audio and submit the gathered input
- `timeout` (Number) The number of seconds to wait for a DTMF keypress before timing out (1–30)
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `voice` (String) The text-to-speech voice to use when using `say` (e.g. `alice`). Conflicts with `play`

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

- `keypress` (String) The name of the next widget when the caller presses a key (DTMF input)
- `speech` (String) The name of the next widget when speech input is detected
- `timeout` (String) The name of the next widget when the gather times out without input
