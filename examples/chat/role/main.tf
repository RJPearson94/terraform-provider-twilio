resource "twilio_chat_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-role"
  type          = "channel"
  permissions = [
    "sendMessage",
    "leaveChannel"
  ]
}
