resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test-${random_string.random.result}"
}

resource "twilio_messaging_alpha_sender" "alpha_sender" {
  service_sid  = twilio_messaging_service.service.sid
  alpha_sender = var.alpha_sender
}
