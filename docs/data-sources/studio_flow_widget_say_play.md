---
page_title: "Twilio Studio Flow Widget - Say/ Play"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the say/ play widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `digits` - (Optional) The digits to send as DTMF tones
- `language` - (Optional) The language to use when speaking the message. This argument conflicts with `digits` and `play`
- `loop` - (Optional) The number of times to say/ play content
- `play` - (Optional) The URL of media content to play. This value can be either a liquid template or a URL
- `say` - (Optional) The text to say
- `voice` - (Optional) The voice to use when speaking the message. This argument conflicts with `digits` and `play`

~> Exactly one of the following arguments: `digits`, `play` or `say` must be specified
~> Due to data type and validation restrictions liquid templates are not supported for the `loop` argument. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `audio_complete` - (Optional) The widget to transition to when the message has been read or the audio content has been played

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the say/ play widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the say/ play widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the say/ play widget
- `json` - The JSON state definition for the say/ play widget
