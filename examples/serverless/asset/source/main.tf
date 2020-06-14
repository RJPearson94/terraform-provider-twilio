resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_serverless_service" "service" {
  unique_name   = "rjpearson94-${random_string.random.result}"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}

resource "twilio_serverless_asset_version" "asset_version" {
  service_sid  = twilio_serverless_service.service.sid
  asset_sid    = twilio_serverless_asset.asset.sid
  source       = "module.png"
  source_hash  = filemd5("${module.path}/module.png")
  content_type = "image/png"
  path         = "/test-2"
  visibility   = "private"
}

