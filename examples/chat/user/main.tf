resource "twilio_chat_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "twilio-test"
}
