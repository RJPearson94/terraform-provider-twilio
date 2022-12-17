resource "twilio_verify_service" "service" {
  friendly_name  = "${var.prefix}-verify-service"
  lookup_enabled = false
}