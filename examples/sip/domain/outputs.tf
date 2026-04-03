output "credential_list" {
  description = "The SIP credential list"
  value       = twilio_sip_credential_list.credential_list
}

output "credential" {
  description = "The SIP credential"
  value       = twilio_sip_credential.credential
  sensitive   = true
}

output "domain" {
  description = "The SIP domain"
  value       = twilio_sip_domain.domain
}

output "credential_list_mapping" {
  description = "The SIP domain credential list mapping"
  value       = twilio_sip_domain_credential_list_mapping.mapping
}
