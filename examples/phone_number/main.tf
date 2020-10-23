resource "twilio_phone_number" "phone_number" {
  account_sid   = var.account_sid
  phone_number  = var.phone_number
}
