resource "twilio_sip_trunking_trunk" "trunk" {}

resource "twilio_sip_trunking_phone_number" "phone_number" {
  trunk_sid = twilio_sip_trunking_trunk.trunk.sid
  sid       = var.phone_number
}