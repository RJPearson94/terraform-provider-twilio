---
page_title: "Twilio Studio Flow Widget - Make outgoing call"
subcategory: "Studio"
---

# twilio_studio_flow_widget_make_outgoing_call Data Source

Use this data source to generate the JSON for the Studio Flow make outgoing call widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/make-outgoing-call) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Basic

```hcl
data "twilio_studio_flow_widget_make_outgoing_call" "make_outgoing_call" {
  name = "MakeOutgoingCall"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_make_outgoing_call" "make_outgoing_call" {
  name = "MakeOutgoingCall"

  transitions {
    answered  = "AnsweredTransition"
    busy      = "BusyTransition"
    failed    = "FailedTransition"
    no_answer = "NoAnswerTransition"
  }

  detect_answering_machine               = true
  from                                   = "{{flow.channel.address}}"
  to                                     = "{{contact.channel.address}}"
  machine_detection                      = "Enable"
  machine_detection_speech_end_threshold = "500"
  machine_detection_speech_threshold     = "1000"
  machine_detection_silence_timeout      = "2000"
  machine_detection_timeout              = "10"
  record                                 = true
  send_digits                            = "1234"
  sip_auth_password                      = "test2"
  sip_auth_username                      = "test"
  timeout                                = 5
  trim                                   = "trim-silence"

  offset {
    x = 10
    y = 20
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the make outgoing call widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `detect_answering_machine` - (Optional) Whether to detect if an answering machine has answered the call
- `from` - (Optional) The caller ID phone number. Default is `{{flow.channel.address}}`
- `machine_detection` - (Optional) The answering machine detection mode. This value can be either a liquid template or the string `Enable` or `DetectMessageEnd`. This is only applicable when `detect_answering_machine` is set to `true`
- `machine_detection_silence_timeout` - (Optional) The time in milliseconds of silence to wait for before timing out. This value can be either a liquid template or a number string in the range `2000` to `10000` (inclusive). This is only applicable when `detect_answering_machine` is set to `true`
- `machine_detection_speech_end_threshold` - (Optional) The time in milliseconds of no speech to wait for before completing the call. This value can be either a liquid template or a number string in the range `500` to `5000` (inclusive). This is only applicable when `detect_answering_machine` is set to `true`
- `machine_detection_speech_threshold` - (Optional) The time in milliseconds of speech that will determine if a human or a machine answered the call. This value can be either a liquid template or a number string in the range `1000` to `6000` (inclusive). This is only applicable when `detect_answering_machine` is set to `true`
- `machine_detection_timeout` - (Optional) The time in seconds to perform answer machine detection. This value can be either a liquid template or a number string in the range `3` to `120` (inclusive). This is only applicable when `detect_answering_machine` is set to `true`
- `record` - (Optional) Whether the call should be recorded
- `recording_channels` - (Optional) The channels which should be recorded. Valid values include: `mono` or `dual`
- `recording_status_callback_url` - (Optional) The URL which receives the recording completion callback. This value can be either a liquid template or a URL
- `send_digits` - (Optional) A set of keys to dial once the call is connected
- `sip_auth_password` - (Optional) A password to authenticate the caller with when making a SIP call
- `sip_auth_username` - (Optional) The username of the caller to use when making a SIP call
- `timeout` - (Optional) The time in seconds to wait will the call is ringing before timing out
- `to` - (Optional) The recipient of the call. Default is `{{contact.channel.address}}`
- `trim` - (Optional) A string to indicate whether silence should be removed from the end of the recording. This value can be either a liquid template or the string `trim-silence` or `do-not-trim`

~> Due to data type and validation restrictions liquid templates are not supported for the `detect_answering_machine`, `record` and `timeout` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `answered` - (Optional) The widget to transition to when the call was answered
- `busy` - (Optional) The widget to transition to when the call is busy
- `failed` - (Optional) The widget to transition to when the call fails
- `no_answer` - (Optional) The widget to transition to when the call is not answered

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the make outgoing call widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the make outgoing call widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the make outgoing call widget
- `json` - The JSON state definition for the make outgoing call widget
