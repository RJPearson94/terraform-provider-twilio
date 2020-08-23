resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_proxy_service" "service" {
  unique_name = "twilio-test-${random_string.random.result}"
}

resource "twilio_proxy_phone_number" "phone_number" {
  service_sid = twilio_proxy_service.service.sid
  sid         = var.phone_number_sid
  is_reserved = true
}
