output "trunk" {
  description = "The SIP Trunk"
  value       = twilio_sip_trunking_trunk.trunk
}

output "origination_url" {
  description = "The SIP Trunk Origination URL"
  value       = twilio_sip_trunking_origination_url.origination_url
}
