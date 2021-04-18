resource "twilio_iam_api_key" "api_key" {
  account_sid   = var.account_sid
  friendly_name = "Test API Key"
}
