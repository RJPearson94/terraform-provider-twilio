resource "twilio_sip_credential_list" "credential_list" {
  account_sid   = var.account_sid
  friendly_name = "Terraform managed - SIP credential list"
}

resource "twilio_sip_credential" "credential" {
  account_sid         = var.account_sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
  username            = var.sip_username
  password            = var.sip_password
}

resource "twilio_sip_domain" "domain" {
  account_sid      = var.account_sid
  domain_name      = var.domain_name
  friendly_name    = "Terraform managed - SIP domain"
  secure           = true
  sip_registration = false

  voice {
    url                = var.voice_url
    method             = "POST"
    status_callback_url    = var.voice_status_callback_url
    status_callback_method = "POST"
  }
}

resource "twilio_sip_domain_credential_list_mapping" "mapping" {
  account_sid         = var.account_sid
  domain_sid          = twilio_sip_domain.domain.sid
  credential_list_sid = twilio_sip_credential_list.credential_list.sid
}
