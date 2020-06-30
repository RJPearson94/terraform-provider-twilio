resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_proxy_service" "service" {
  unique_name = "twilio-test-${random_string.random.result}"
}
