resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_role" "role" {
  service_sid   = twilio_conversations_service.service.sid
  friendly_name = "twilio-test-role"
  type          = "conversation"
  permissions = [
    "sendMediaMessage",
    "sendMessage"
  ]
}