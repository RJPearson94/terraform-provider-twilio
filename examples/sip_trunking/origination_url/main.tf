resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_origination_url" "origination_url" {
  trunk_sid = twilio_sip_trunking_trunk.trunk.sid
  friendly_name = "twilio-test"
  enabled       = true
  priority      = 1
  sip_url       = "sip:test@test.com"
  weight        = 1
}