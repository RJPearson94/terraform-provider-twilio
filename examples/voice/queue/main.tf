resource "twilio_voice_queue" "queue" {
  account_sid   = var.account_sid
  friendly_name = "Test Queue"
}
