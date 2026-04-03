---
page_title: "twilio_studio_flow_widget_record_voicemail Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_record_voicemail Data Source

Use this data source to generate the JSON for the Studio Flow record voicemail widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/record-voicemail) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_record_voicemail" "record_voicemail" {
  name = "RecordVoicemail"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_record_voicemail" "record_voicemail" {
  name = "RecordVoicemail"

  transitions {
    hangup             = "HangupTransition"
    no_audio           = "NoAudioTransition"
    recording_complete = "RecordingCompleteTransition"
  }

  max_length                    = 1000
  play_beep                     = "true"
  recording_status_callback_url = "http://localhost.com/recording"
  timeout                       = 10
  finish_on_key                 = "1"
  transcribe                    = true
  transcription_callback_url    = "http://localhost.com/transcript"
  trim                          = "trim-silence"

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

- `finish_on_key` (String) The DTMF key that ends the recording (e.g. `#`)
- `max_length` (Number) The maximum recording length in seconds (1–14400)
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `play_beep` (String) Whether to play a beep before starting the recording. Valid values: `true`, `false`
- `recording_status_callback_url` (String) The HTTP/HTTPS URL to receive recording status callback events
- `timeout` (Number) The number of seconds of silence before the recording automatically stops
- `transcribe` (Boolean) Whether to transcribe the voicemail recording
- `transcription_callback_url` (String) The HTTP/HTTPS URL to receive the transcription result
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

- `hangup` (String) The name of the next widget when the caller hangs up during recording
- `no_audio` (String) The name of the next widget when no audio is detected during recording
- `recording_complete` (String) The name of the next widget when the voicemail recording completes successfully
