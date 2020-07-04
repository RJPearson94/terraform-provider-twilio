resource "twilio_chat_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-channel"
}

resource "twilio_chat_channel_studio_webhook" "studio_webhook" {
  service_sid = twilio_chat_service.service.sid
  channel_sid = twilio_chat_channel.channel.sid
  flow_sid    = "FWaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}
