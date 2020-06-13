resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-${random_string.random.result}"
  friendly_name = "twilio-test"
}
