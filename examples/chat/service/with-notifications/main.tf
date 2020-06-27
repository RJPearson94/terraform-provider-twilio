resource "twilio_chat_service" "service" {
  friendly_name = "twilio-test-with-notifications"

  notifications {
    log_enabled = true

    new_message {
      enabled  = true
      template = "$${CHANNEL};$${USER}: $${MESSAGE}"
    }

    added_to_channel {
      enabled  = true
      template = "You have been added to channel $${CHANNEL} by $${USER}"
    }

    removed_from_channel {
      enabled  = true
      template = "$${USER} has removed you from the channel $${CHANNEL}"
    }

    invited_to_channel {
      enabled  = true
      template = "$${USER} has invited you to join the channel $${CHANNEL}"
    }
  }
}
