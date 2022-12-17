resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "${var.prefix}-sub-account"
}