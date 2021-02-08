resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "test-user"
}