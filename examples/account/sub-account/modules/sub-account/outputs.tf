output "account_details" {
  description = "Account resource details"
  value       = twilio_account_sub_account.sub_account
  sensitive   = true
}