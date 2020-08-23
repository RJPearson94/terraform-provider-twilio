resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_proxy_service" "service" {
  unique_name = "twilio-test-${random_string.random.result}"
}

resource "twilio_proxy_short_code" "short_code" {
  service_sid = twilio_proxy_service.service.sid
  sid         = var.short_code_sid
}
