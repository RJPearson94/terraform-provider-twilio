variable "account_sid" {
  description = "The SID of the Twilio account"
  type        = string
}

variable "domain_name" {
  description = "The unique SIP domain name (must end in .sip.twilio.com)"
  type        = string
}

variable "voice_url" {
  description = "The URL Twilio calls when the domain receives an inbound SIP call"
  type        = string
}

variable "voice_status_callback_url" {
  description = "The URL Twilio calls to pass status parameters to your application"
  type        = string
}

variable "sip_username" {
  description = "The username for the SIP credential"
  type        = string
}

variable "sip_password" {
  description = "The password for the SIP credential"
  type        = string
  sensitive   = true
}
