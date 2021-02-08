resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_service_configuration" "configuration" {
  service_sid = twilio_conversations_service.service.sid
  reachability_enabled = false
}

resource "twilio_conversations_service_notification" "service_notification" {
  service_sid = twilio_conversations_service.service.sid

  log_enabled = true

  new_message {
    enabled = true
    badge_count_enabled = true
    template = "$${CONVERSATION}:$${PARTICIPANT}: $${MESSAGE}"
  }

  added_to_conversation {
    enabled = true
    template = "You have been added to the conversation $${CONVERSATION} by $${PARTICIPANT}"
  }

  removed_from_conversation {
    enabled = true
    template = "$${PARTICIPANT} has removed you from the conversation $${CONVERSATION}"
  }
}