resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test-${random_string.random.result}"
  friendly_name = "test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
  source        = "module.png"
  source_hash   = filemd5("${path.module}/module.png")
  content_type  = "image/png"
  path          = "/test"
  visibility    = "private"
}
