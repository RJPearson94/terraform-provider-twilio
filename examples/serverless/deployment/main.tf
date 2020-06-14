resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_serverless_service" "service" {
  unique_name   = "rjpearson94-${random_string.random.result}"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}

resource "twilio_serverless_function_version" "function_version" {
  service_sid  = twilio_serverless_service.service.sid
  function_sid = twilio_serverless_function.function.sid
  content = <<EOF
exports.handler = function (context, event, callback) {
  callback(null, "Hello World");
};
EOF
  content_type = "application/javascript"
  content_file_name    = "helloWorld.js"
  path         = "/test-function"
  visibility   = "public"
}

resource "twilio_serverless_asset" "asset" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}

resource "twilio_serverless_asset_version" "asset_version" {
  service_sid  = twilio_serverless_service.service.sid
  asset_sid    = twilio_serverless_asset.asset.sid
  source       = "module.png"
  source_hash  = filemd5("${path.module}/module.png")
  content_type = "image/png"
  path         = "/test-asset"
  visibility   = "public"
}

resource "twilio_serverless_build" "build" {
  service_sid   = twilio_serverless_service.service.sid
  function_version_sids = [twilio_serverless_function_version.function_version.sid]
  asset_version_sids = [twilio_serverless_asset_version.asset_version.sid]
  dependencies = {
    "twilio": "3.6.3"
  }

  polling {
    enabled = true
  }
}

resource "twilio_serverless_environment" "environment" {
  service_sid   = twilio_serverless_service.service.sid
  unique_name   = "test"
}

resource "twilio_serverless_deployment" "deployment" {
  service_sid = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  build_sid = twilio_serverless_build.build.sid
}