data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
  name = "SendMessageToAgent"

  workflow = var.workflow_sid
  channel  = var.channel_sid
  attributes = jsonencode({
    "name" : "{{trigger.message.ChannelAttributes.from}}",
    "channelType" : "{{trigger.message.ChannelAttributes.channel_type}}",
    "channelSid" : "{{trigger.message.ChannelSid}}"
  })

  offset {
    x = 200
    y = 240
  }
}

data "twilio_studio_flow_widget_trigger" "trigger" {
  name = "Trigger"

  transitions {
    incoming_message = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
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
    json = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.json
  }

  states {
    json = data.twilio_studio_flow_widget_trigger.trigger.json
  }
}
