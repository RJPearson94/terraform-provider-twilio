resource "twilio_chat_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-channel"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "twilio-test"
}

resource "twilio_chat_channel_member" "member" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  identity    = twilio_chat_user.user.identity
  attributes  = <<EOF
{
	"test": "test"
}
EOF
}
