resource "twilio_studio_flow" "flow" {
  friendly_name = "Test Studio Flow"
  status        = "published"
  definition    = file("${path.module}/flow.json")
}
