resource "random_string" "random" {
  length  = 16
  special = false
}

resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test-${random_string.random.result}"
  friendly_name = "test"
}

resource "twilio_serverless_function" "function" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
  source        = "helloWorld.js"
  source_hash   = filemd5("${path.module}/helloWorld.js")
  content_type  = "application/javascript"
  path          = "/test-function"
  visibility    = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid

  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }

  dependencies = {
    "twilio" : "3.6.3"
  }
}
