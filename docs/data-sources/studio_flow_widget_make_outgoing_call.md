---
page_title: "twilio_studio_flow_widget_make_outgoing_call Data Source - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions

### Optional

- `detect_answering_machine` (Boolean) Whether to enable answering machine detection on the outgoing call
- `from` (String) The caller ID phone number for the outgoing call. Defaults to `{{flow.channel.address}}`
- `machine_detection` (String) The answering machine detection mode. Valid values: `Enable`, `DetectMessageEnd`
- `machine_detection_silence_timeout` (String) The duration of silence in milliseconds after a greeting that indicates a machine (2000–10000)
- `machine_detection_speech_end_threshold` (String) The duration of no speech in milliseconds that marks the end of a greeting (500–5000)
- `machine_detection_speech_threshold` (String) The minimum duration of speech in milliseconds to classify as a human (1000–6000)
- `machine_detection_timeout` (String) The total time in seconds to wait for machine detection to complete (3–120)
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `record` (Boolean) Whether to record the outgoing call
- `recording_channels` (String) The number of recording channels. Valid values: `mono`, `dual`
- `recording_status_callback_url` (String) The HTTP/HTTPS URL to receive recording status callback events
- `send_digits` (String) DTMF digits to send after the call is connected (e.g. `ww1234#` where `w` is a 0.5s pause)
- `sip_auth_password` (String, Sensitive) The password for SIP authentication. Sensitive — will not be shown in logs or plans
- `sip_auth_username` (String) The username for SIP authentication
- `timeout` (Number) The number of seconds to wait for the call to be answered before timing out
- `to` (String) The phone number or SIP URI to call. Defaults to `{{contact.channel.address}}`
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))
- `trim` (String) Whether to trim silence from the beginning and end of recordings. Valid values: `trim-silence`, `do-not-trim`

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

- `answered` (String) The name of the next widget when the outgoing call is answered
- `busy` (String) The name of the next widget when the called party is busy
- `failed` (String) The name of the next widget when the outgoing call fails
- `no_answer` (String) The name of the next widget when the outgoing call is not answered
