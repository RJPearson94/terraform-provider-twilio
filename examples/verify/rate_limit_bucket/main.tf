resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = twilio_verify_service.service.sid
  unique_name = "Test Service Rate Limit"
}

resource "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  service_sid    = twilio_verify_service_rate_limit.rate_limit.service_sid
  rate_limit_sid = twilio_verify_service_rate_limit.rate_limit.sid
  max            = 10
  interval       = 2
}