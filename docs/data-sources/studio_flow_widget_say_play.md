---
page_title: "twilio_studio_flow_widget_say_play Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_say_play Data Source

Use this data source to generate the JSON for the Studio Flow say/ play widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/sayplay) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Say

```hcl
data "twilio_studio_flow_widget_say_play" "say_play" {
  name = "SayPlay"
  say  = "Hello World"
}
```

## Play

```hcl
data "twilio_studio_flow_widget_say_play" "say_play" {
  name = "SayPlay"
  play = "http://localhost.com"
}
```

## Digits

```hcl
data "twilio_studio_flow_widget_say_play" "say_play" {
  name   = "SayPlay"
  digits = "123"
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `digits` (String) DTMF digits to play to the caller. Exactly one of `digits`, `play`, or `say` must be set
- `language` (String) The language for text-to-speech (e.g. `en-US`). Conflicts with `digits` and `play`
- `loop` (Number) How many times to repeat the audio or speech. Use 0 for infinite loop
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `play` (String) URL of the audio file to play. Exactly one of `digits`, `play`, or `say` must be set
- `say` (String) Text to speak to the caller using text-to-speech. Exactly one of `digits`, `play`, or `say` must be set
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `voice` (String) The text-to-speech voice to use (e.g. `alice`). Conflicts with `digits` and `play`

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

- `audio_complete` (String) The name of the next widget when audio playback or text-to-speech finishes
