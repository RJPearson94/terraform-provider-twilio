resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test-webhook"
  webhook_url   = "http://localhost/webhook"
  events = [
    "onDialogueEnd"
  ]
}