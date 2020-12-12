output "trunk" {
  description = "The SIP Trunk"
  value       = twilio_sip_trunking_trunk.trunk
}

output "phone_number" {
  description = "The SIP Trunk Phone Number"
  value       = twilio_sip_trunking_phone_number.phone_number
}
