locals {
  include_autopilot = var.autopilot_assistant_sid != null

  states = compact([
    data.twilio_studio_flow_widget_trigger.trigger.json,
    data.twilio_studio_flow_widget_send_to_flex.send_to_flex.json,
    join("", data.twilio_studio_flow_widget_send_to_autopilot.send_to_autopilot.*.json),
  ])

  message_transition = coalesce(
    join("", data.twilio_studio_flow_widget_send_to_autopilot.send_to_autopilot.*.name),
    data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
  )
}

data "twilio_studio_flow_widget_send_to_autopilot" "send_to_autopilot" {
  count = local.include_autopilot ? 1 : 0
  name  = "SendToAutopilot"

  offset {
    x = 200
    y = 240
  }

  transitions {
    failure = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
    timeout = data.twilio_studio_flow_widget_send_to_flex.send_to_flex.name
  }

  autopilot_assistant_sid = var.autopilot_assistant_sid
}

data "twilio_studio_flow_widget_send_to_flex" "send_to_flex" {
  name = "SendMessageToAgent"

  workflow_sid = var.workflow_sid
  channel_sid  = var.channel_sid
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
    incoming_message = local.message_transition
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

  dynamic "states" {
    for_each = local.states
    content {
      json = states.value
    }
  }
}
