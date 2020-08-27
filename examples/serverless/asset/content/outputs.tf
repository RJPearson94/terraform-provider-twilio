output "service" {
  description = "The Generated Serverless Service"
  value       = twilio_serverless_service.service
}

output "asset" {
  description = "The Generated Serverless Asset"
  value       = twilio_serverless_asset.asset
}
