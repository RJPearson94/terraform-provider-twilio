resource "twilio_studio_flow" "flow" {
  friendly_name = "With widgets"
  status        = "draft"
  definition    = data.twilio_studio_flow_definition.definition.json
  validate      = true
}
