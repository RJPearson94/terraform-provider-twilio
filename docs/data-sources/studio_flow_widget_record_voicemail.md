---
page_title: "Twilio Studio Flow Widget - Record voicemail"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the record voicemail widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `finish_on_key` - (Optional) The keypress which terminates the recording
- `max_length` - (Optional) The maximum length (in seconds) of the recording. The value must be between `1` and `14400`
- `play_beep` - (Optional) A string to represent whether a beep should be played or suppressed before recording starts. This value can be either a liquid template or the string `true` or `false`
- `recording_status_callback_url` - (Optional) The URL which receives the recording completion callback. This value can be either a liquid template or a URL
- `timeout` - (Optional) The time in seconds of silence to wait before terminating the recording
- `transcribe` - (Optional) Whether the call should be transcribed
- `trim` - (Optional) A string to indicate whether silence should be removed from the end of the recording. This value can be either a liquid template or the string `trim-silence` or `do-not-trim`
- `transcription_callback_url` - (Optional) The URL which receives the transcription callback. This value can be either a liquid template or a URL

~> Due to data type and validation restrictions liquid templates are not supported for the `transcribe`, `max_length` and `timeout` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `hangup` - (Optional) The widget to transition to when the caller ends the call
- `no_audio` - (Optional) The widget to transition to when no audio is received
- `recording_complete` - (Optional) The widget to transition to when the recording is complete

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the record voicemail widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the record voicemail widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the record voicemail widget
- `json` - The JSON state definition for the record voicemail widget
