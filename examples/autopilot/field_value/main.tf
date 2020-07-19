resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
}

resource "twilio_autopilot_field_value" "field_value" {
  assistant_sid  = twilio_autopilot_field_type.field_type.assistant_sid
  field_type_sid = twilio_autopilot_field_type.field_type.sid
  language       = "en-US"
  value          = "test"
}