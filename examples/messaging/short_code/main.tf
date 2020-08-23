resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test-${random_string.random.result}"
}

resource "twilio_messaging_short_code" "short_code" {
  service_sid = twilio_messaging_service.service.sid
  sid         = var.short_code_sid
}
