resource "twilio_studio_flow" "flow" {
  friendly_name = "Test Flex Webchat Studio Flow"
  status        = "draft"
  definition = templatefile("${path.module}/flow.tpl.json", {
    workflow_sid = var.workflow_sid
    channel_sid  = var.channel_sid
  })
  validate = true
}
