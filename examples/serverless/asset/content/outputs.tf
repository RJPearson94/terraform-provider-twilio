output "service" {
  description = "The Generated Serverless Service"
  value       = twilio_serverless_service.service
}

output "asset" {
  description = "The Generated Serverless Asset"
  value       = twilio_serverless_asset.asset
}

output "asset_version" {
  description = "The Generated Serverless Asset Version"
  value       = twilio_serverless_asset_version.asset_version
}