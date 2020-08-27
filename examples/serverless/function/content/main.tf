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

  content           = file("${path.module}/helloWorld.js")
  content_type      = "application/javascript"
  content_file_name = "helloWorld4.js"
  path              = "/test-function"
  visibility        = "private"
}
