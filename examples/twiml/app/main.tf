resource "twilio_twiml_app" "app" {
  account_sid   = data.twilio_account_details.account_details.sid
  friendly_name = "Test TwiML app"
}
