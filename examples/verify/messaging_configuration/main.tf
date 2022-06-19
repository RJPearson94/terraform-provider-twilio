resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "Test Messaging Service"
}

resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = twilio_verify_service.service.sid
  messaging_service_sid = twilio_messaging_service.service.sid
  country_code          = "GB"
}