---
page_title: "Twilio Studio Flow Definition"
subcategory: "Studio"
---

# twilio_studio_flow_definition Data Source

This data source can be used to generate a Studio Flow definition JSON which can be supplied as an argument to the `twilio_studio_flow` resource

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

### Studio Flow definition with Trigger, Send to Flex and Send to Autopilot widgets

```hcl
data "twilio_studio_flow_widget_send_to_autopilot" "send_to_autopilot" {
  name = "SendToAutopilot"

  offset {
    x = 200
    y = 240
  }

  transitions {
    failure = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
    timeout = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
  }

  autopilot_assistant_sid = "UAaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  from                    = "{{flow.channel.address}}"
  body                    = "{{trigger.message.Body}}"
  timeout                 = 14400
}

data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
  name = "SendMessageToAgent"

  workflow_sid = "WWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  channel_sid  = "TCaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })

  offset {
    x = 270
    y = 540
  }
}

data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  transitions {
    incoming_message = data.twilio_studio_flow_widget_send_to_autopilot.send_to_autopilot.name
  }

  offset {
    x = 200
    y = 0
  }
}

data "twilio_studio_flow_definition" "definition" {
  description   = "Bot flow for creating a Flex webchat task"
  initial_state = data.twilio_studio_flow_widget_trigger.trigger.name

  flags {
    allow_concurrent_calls = true
  }

  states {
    json = data.twilio_studio_flow_widget_trigger.trigger.json
  }

  states {
    json = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.json
  }

  states {
    json = data.twilio_studio_flow_widget_send_to_autopilot.send_to_autopilot.json
  }
}
```

### Studio Flow definition with Studio Flow Resource

```hcl
data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  offset {
    x = 200
    y = 0
  }
}

data "twilio_studio_flow_definition" "definition" {
  description   = "Flow with trigger widget"
  initial_state = data.twilio_studio_flow_widget_trigger.trigger.name

  flags {
    allow_concurrent_calls = true
  }

  states {
    json = data.twilio_studio_flow_widget_trigger.trigger.json
  }
}

resource "twilio_studio_flow" "flow" {
  friendly_name = "With widgets"
  status        = "draft"
  definition    = data.twilio_studio_flow_definition.definition.json
  validate      = true
}
```

## Argument Reference

The following arguments are supported:

- `description` - (Mandatory) A description of the flow
- `flags` - (Optional) A `flags` block as documented below
- `initial_state` - (Mandatory) The first state to transition to when executing the flow
- `states` - (Mandatory) A list of `state` blocks as documented below

---

A `flags` block supports the following:

- `allow_concurrent_calls` - (Mandatory) Whether the flow should allow concurrent calls

---

A `state` block supports the following:

- `json` - (Mandatory) A JSON string of the state definition

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the flow definition
- `json` - The JSON for the flow definition
