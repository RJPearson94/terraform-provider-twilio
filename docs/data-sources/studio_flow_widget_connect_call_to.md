---
page_title: "Twilio Studio Flow Widget - Connect call to"
subcategory: "Studio"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the connect call to widget
- `transitions` - (Optional) A `transitions` block as documented below
- `offset` - (Optional) A `offset` block as documented below
- `noun` - (Optional) The noun to use to indicate how the call should be connected. Valid values include: `client`, `conference`, `number`, `number-multi`, `sim` or `sip`
- `caller_id` - (Optional) The phone number which will be used as the caller ID for the call. Default value is `{{contact.channel.address}}`
- `record` - (Optional) Whether the call should be recorded
- `sip_endpoint` - (Optional) The SIP endpoint to connect the call to. This should be set when setting the `noun` to `sip`
- `sip_password` - (Optional) A password to authenticate the caller with when connecting the call
- `sip_username` - (Optional) The username of the caller to use when connecting the call
- `timeout` - (Optional) The amount of time in seconds to wait before timing out
- `to` - (Optional) The target (phone number, SIM, conference, etc.) to connect the call to

~> Due to data type and validation restrictions liquid templates are not supported for the `noun`, `record` and `timeout` arguments. Please see the widget documentation to determine whether other arguments support liquid templates

---

A `transitions` block supports the following:

- `call_completed` - (Optional) The widget to transition to when the call is complete
- `hangup` - (Optional) The widget to transition to when the caller hangs up

---

An `offset` block supports the following:

- `x` - (Optional) The x coordinate to display the connect call to widget in the Studio console. The default value is `0`
- `y` - (Optional) The y coordinate to display the connect call to widget in the Studio console. The default value is `0`

## Attributes Reference

The following attributes are exported:

- `id` - The name of the connect call to widget
- `json` - The JSON state definition for the connect call to widget
