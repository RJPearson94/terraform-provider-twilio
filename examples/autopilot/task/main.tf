resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_task" "task" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
  actions       = <<EOF
{
	"actions": [
		{
			"say": {
				"speech": "Hello World"
			}
		}
	]
}
EOF
}
