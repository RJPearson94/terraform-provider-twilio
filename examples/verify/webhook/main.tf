resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_verify_webhook" "webhook" {
  service_sid   = twilio_verify_service.service.sid
  friendly_name = "Test Verify Webhook"
  event_types   = ["*"]
  webhook_url   = "https://localhost.com/webhook"
}