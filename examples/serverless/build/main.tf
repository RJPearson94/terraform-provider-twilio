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
  source       = "helloWorld.js"
  source_hash  = filemd5("${path.module}/helloWorld.js")
  content_type = "application/javascript"
  path         = "/test-function"
  visibility   = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid           = twilio_serverless_service.service.sid
  function_version_sids = [twilio_serverless_function_version.function_version.sid]
  dependencies = {
    "twilio" : "3.6.3"
  }
}
