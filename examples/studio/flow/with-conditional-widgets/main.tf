resource "twilio_studio_flow" "flow" {
  friendly_name = "Conditional Widgets"
  status        = "draft"
  definition    = data.twilio_studio_flow_definition.definition.json
  validate      = true
}
