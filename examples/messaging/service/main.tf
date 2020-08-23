resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test-${random_string.random.result}"
}
