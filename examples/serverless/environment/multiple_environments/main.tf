resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test-${random_string.random.result}"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "staging" {
  service_sid   = twilio_serverless_service.service.sid
  unique_name   = "staging"
  domain_suffix = "staging"
}

resource "twilio_serverless_environment" "prod" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "prod"
}
