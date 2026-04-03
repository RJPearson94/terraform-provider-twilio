---
page_title: "twilio_studio_flow_widget_connect_call_to Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow_widget_connect_call_to Data Source

Use this data source to generate the JSON for the Studio Flow connect the call to widget. This data source can be used in combination with the `twilio_studio_flow_definition` to generate a Studio Flow definition. See the [docs](https://www.twilio.com/docs/studio/widget-library/connect-call) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

## Connect call to a client

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "client"
  to   = "test"
}
```

## Connect call to conference

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "conference"
  to   = "CFaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
```

## Connect call to a phone number

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "number"
  to   = "+441234567890"
}
```

## Connect call to multiple phone numbers

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "number-multi"
  to   = "+441234567890,+441234567891"
}
```

## Connect call to SIM

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"
  noun = "sim"
  to   = "DEaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
```

## Connect call to a SIP endpoint

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name         = "ConnectCallTo"
  noun         = "sip"
  sip_endpoint = "sip:test@test.com"
}
```

## With all config

```hcl
data "twilio_studio_flow_widget_connect_call_to" "connect_call_to" {
  name = "ConnectCallTo"

  transitions {
    call_completed = "CallCompletedTransition"
    hangup         = "HangupTransitions"
  }

  caller_id    = "{{contact.channel.address}}"
  record       = true
  noun         = "sip"
  timeout      = 30
  sip_username = "test"
  sip_password = "test2"
  sip_endpoint = "sip:test@test.com"

  offset {
    x = 10
    y = 20
  }
}
```

## Schema

### Required

- `name` (String) The unique name of this widget within the flow, used to reference it in transitions
- `noun` (String) The TwiML noun that determines the call destination type. Valid values: `client`, `conference`, `number`, `number-multi`, `sim`, `sip`

### Optional

- `caller_id` (String) The caller ID to display on the outbound call. Defaults to `{{contact.channel.address}}`
- `offset` (Block List, Max: 1) The position of this widget in the Studio visual editor (see [below for nested schema](#nestedblock--offset))
- `record` (Boolean) Whether to record the outbound call
- `sip_endpoint` (String) The SIP URI to dial when `noun` is `sip` (e.g. `sip:user@domain.com`)
- `sip_password` (String, Sensitive) The password for SIP authentication. Sensitive — will not be shown in logs or plans
- `sip_username` (String) The username for SIP authentication
- `timeout` (Number) The number of seconds to wait for the call to be answered before timing out
- `to` (String) The destination to call. Format depends on `noun` (e.g. a phone number in E.164 format, a client name, or a conference name)
- `transitions` (Block List, Max: 1) The next widget(s) to transition to after this widget (see [below for nested schema](#nestedblock--transitions))

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

- `call_completed` (String) The name of the next widget when the outbound call completes
- `hangup` (String) The name of the next widget when the caller hangs up
