resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "twilio-test-${random_string.random.result}"
}
